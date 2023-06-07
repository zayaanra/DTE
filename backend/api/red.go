package api

import (
	"errors"
	"net"
	"io"
	"log"
)


type REDServer interface {
	// Sends the given message to the given peer
	// It returns an error if sending the message failed and nil if it succeeded
	Send(peer string) error

	// Reads a single message sent to this server's channel
	Recv() 
	
	// Terminate the server, closing the listening socket
	// on this server.
	Terminate()
}

type RServer struct {
	// The address for this server
	addr string

	// The listener socket for this server.
	// It listens for incoming connections.
	ln net.Listener
}

// Create a new RED server associated with the given address.
// The newly created RED server begins send or receive messages immedaiately.
// This function returns an error if the server was not able to be created.
func NewREDServer(addr string) (REDServer, error) {
	// TODO - Should we use TCP or UDP?
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}
	rs := &RServer{addr, ln}
	go func() {
		for {
			conn, err := ln.Accept()
			// TODO - If error upon accepting a connection, how should we handle it
			if err != nil {
				return;
			}
			defer conn.Close()
			buffer := make([]byte, 24)
			bytes, _ := io.ReadFull(conn,buffer)
			log.Printf("Received a message: %v\n", bytes)
		}
	}()
	return rs, nil
}

func (rs *RServer) Send(addr string) error {
	return errors.New("TODO")
}

func (rs *RServer) Recv() {
}

func (rs *RServer) Terminate() {
	rs.ln.Close()
}
