package tests

import (
	"testing"

	"github.com/zayaanra/RED/backend/api"
	"github.com/zayaanra/RED/backend/internal/handler"
)

/* handler_send_test.go contains all test cases that deal with seeing if our handler can properly send messages. It does not check to see if the receiver can receive such messages. */

// This tests if we can send an INVITE message
func TestSendInvite(t *testing.T) {
	rh1_addr := "localhost:3000"
	rh2_addr := "localhost:3001"
	rh1, err1 := handler.NewHandler(rh1_addr)
	rh2, err2 := handler.NewHandler(rh2_addr)
	if err1 != nil || err2 != nil {
		t.Fatalf("Handler failed to initialize")
	}
	defer rh1.Terminate()
	defer rh2.Terminate()
	smsg := &api.REDMessage{Type: api.MessageType_INVITE, Sender: rh1_addr, Receipient: rh2_addr}
	err := rh1.Send(smsg, rh2_addr)
	if err != nil {
		t.Errorf("Sending an INVITE failed")
	}
}

// This tests if we can send an EDIT message.
// We should be able to send an EDIT of type INSERT.
func TestSendEditINSERT(t *testing.T) {
	rh1_addr := "localhost:3002"
	rh2_addr := "localhost:3003"
	rh1, err1 := handler.NewHandler(rh1_addr)
	rh2, err2 := handler.NewHandler(rh2_addr)
	if err1 != nil || err2 != nil {
		t.Fatalf("Handler failed to initialize")
	}
	defer rh1.Terminate()
	defer rh2.Terminate()
	e := &api.Edit{Type: api.EditType_INSERT, Pos: 1, Char: 2}
	smsg := &api.REDMessage{Type: api.MessageType_EDIT, Sender: rh1_addr, Receipient: rh2_addr, Edit: e}
	err := rh1.Send(smsg, rh2_addr)
	if err != nil {
		t.Errorf("Sending an EDIT of type INSERT failed")
	}
}

// We should be able send an EDIT message of type DELETE
func TestSendEditDELETE(t *testing.T) {
	rh1_addr := "localhost:3004"
	rh2_addr := "localhost:3005"
	rh1, err1 := handler.NewHandler(rh1_addr)
	rh2, err2 := handler.NewHandler(rh2_addr)
	if err1 != nil || err2 != nil {
		t.Fatalf("Handler failed to initialize")
	}
	defer rh1.Terminate()
	defer rh2.Terminate()
	e := &api.Edit{Type: api.EditType_DELETE, Pos: 12, Char: 32}
	smsg := &api.REDMessage{Type: api.MessageType_EDIT, Sender: rh1_addr, Receipient: rh2_addr, Edit: e}
	err := rh1.Send(smsg, rh2_addr)
	if err != nil {
		t.Errorf("Sending an EDIT of type DELETE failed")
	}	
}
