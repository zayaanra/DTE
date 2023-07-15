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

	// List of peers that are connected to this REDServer's editing session
	peers []string
}

// Create a new RED server associated with the given address.
// The newly created RED server begins send or receive messages immedaiately.
// This function returns an error if the server was not able to be created.
func NewREDServer(addr string) (api.REDServer, error) {
	rh, err := handler.NewHandler(addr)
	if err != nil {
		return nil, err
	}

	peers := []string{}
	rs := &RServer{addr, rh, peers}
	go func(rh *handler.Handler) {
		for {
			select {
			case rmsg := <-rh.M:
				if rmsg == nil {
					log.Println("Killing server...")
					return
				}
				if rmsg.Type == api.MessageType_INVITE {
					log.Printf("%s accepted an INVITE from %s\n", rs.addr, rmsg.Sender)
					rs.peers = append(rs.peers, rmsg.Sender)
				}
			}
		}
	}(rh)
	return rs, nil
}

// Invites a peer to their editing session by sending an INVITE message.
func (rs *RServer) Invite(addr string) error {
	log.Printf("%s is sending an INVITE to %s\n", rs.addr, addr)
	smsg := &api.REDMessage{Type: api.MessageType_INVITE, Sender: rs.addr, Receipient: addr}
	err := rs.handler.Send(smsg, addr)
	rs.peers = append(rs.peers, addr)
	return err
}

// Accepts an invitation from a peer.
func (rs *RServer) Accept() {
}

// Notifies all peers in this editing session of an EDIT
func (rs *RServer) Notify() {
	// TODO - for now, we'll just send the entire text document and have the peer on the other update it's GUI
}

// Terminates the REDServer. It closes any resources that are currently being used.
func (rs *RServer) Terminate() {
	rs.handler.Terminate()
}
