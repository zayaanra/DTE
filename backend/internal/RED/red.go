package RED

import (
	"errors"
	"net"
	"io"
	"log"

	"github.com/zayaanra/RED/backend/api"
)

type Server struct {
	// The address for this server
	addr string

	// The listener socket for this server.
	// It listens for incoming connections.
	ln net.Listener
}

// Create a new RED server associated with the given address.
// The newly created RED server begins send or receive messages immedaiately.
// This function returns an error if the server was not able to be created.
func NewREDServer(addr string) (api.REDServer, error) {
	// TODO - Should we use TCP or UDP?
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}
	s := &Server{addr, ln}
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
	return s, nil
}

func (s *Server) Send(addr string) error {
	return errors.New("TODO")
}

func (s *Server) Recv() {
}

func (s *Server) Terminate() {
	s.ln.Close()
}
