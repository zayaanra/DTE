package red

import (
	
	"github.com/zayaanra/RED/backend/api"
	"github.com/zayaanra/RED/backend/internal/handler"
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
	rs := &RServer{addr, rh}
	return rs, nil
}


// Terminates the REDServer. It closes any resources that are currently being used.
func (rs *RServer) Terminate() {
	rs.handler.Terminate()
}
