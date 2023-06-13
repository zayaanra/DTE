package handler

import (
	"encoding/binary"
	"io"
	"net"

	"github.com/golang/protobuf/proto"

	"github.com/zayaanra/RED/backend/api"
)

type Handler struct {
	// The listener socket for this Handler
	ln net.Listener

	// Received messages are put onto this channel
	M chan *api.REDMessage
}

// Create a new handler that gets attached to the given address.
// It returns an error upon failed initialization.
func NewHandler(addr string) (*Handler, error) {
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}
	msgChan := make(chan *api.REDMessage)
	h := &Handler{ln: ln, M: msgChan}

	go h.Handle()

	return h, nil
}

// Sends the given message to the given address
// It returns an error if sending the message failed.
func (h *Handler) Send(msg *api.REDMessage, addr string) error {
	// Dial the peer
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return err
	}
	defer conn.Close()

	// Marshal the message
	out, err := proto.Marshal(msg)
	if err != nil {
		return err
	}

	// Encode a header into the 'out' array with a little endian unsigned 16-bit integer
	buffer := make([]byte, len(out)+2)
	binary.LittleEndian.PutUint16(buffer, uint16(len(out)))
	copy(buffer[2:], out)

	// Send data to peer
	_, err = conn.Write(buffer)
	if err != nil {
		return err
	}
	return nil
}

// Receives a message and then places it on the msgChan.
func (h *Handler) Recv(conn net.Conn) {
	// When we receive a message, we have to parse the header so we know how many bytes to read.
	header := make([]byte, 2)
	_, err := io.ReadFull(conn, header)
	if err != nil {
		return
	}

	// The header is created as an unsigned 16-bit integer in little endian
	length := binary.LittleEndian.Uint16(header)
	buffer := make([]byte, length)
	_, err = io.ReadFull(conn, buffer)
	if err != nil {
		return
	}

	// Unmarshal the message
	rmsg := &api.REDMessage{}
	if err := proto.Unmarshal(buffer, rmsg); err == nil {
		h.M <- rmsg
	}
}

// Handles accepting incoming connections and receiving messages from those connections.
func (h *Handler) Handle() {
	for {
		conn, err := h.ln.Accept()
		if err != nil {
			return
		}
		defer conn.Close()
		h.Recv(conn)
	}
}

// Closes any resources that are currently in use by this Handler
func (h *Handler) Terminate() {
	h.ln.Close()
}
