package red

import (
	"log"

	"github.com/zayaanra/RED/api"
	"github.com/zayaanra/RED/internal/handler"
)

type RServer struct {
	// The address for this server
	addr string

	// The handler for this server
	handler *handler.Handler
}

// Create a new RED server associated with the given address.
// The newly created RED server begins send or receive messages immedaiately.
// This function returns an error if the server was not able to be created.
func NewREDServer(addr string) (api.REDServer, error) {
	rh, err := handler.NewHandler(addr)
	if err != nil {
		return nil, err
	}

	go func(rh *handler.Handler) {
		for {
			select {
			case rmsg := <-rh.M:
				if rmsg == nil {
					log.Println("Killing server...")
					return
				}
				log.Printf("Received message: %v\n", rmsg)
			}
		}
	}(rh)
	rs := &RServer{addr, rh}
	return rs, nil
}

// Terminates the REDServer. It closes any resources that are currently being used.
func (rs *RServer) Terminate() {
	rs.handler.Terminate()
}
