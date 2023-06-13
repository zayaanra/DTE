package tests

import (
	"testing"

	"github.com/zayaanra/RED/backend/api"
	"github.com/zayaanra/RED/backend/internal/handler"
)

/* handler_recv_test.go contains all test cases that test if a handler can properly receive messages */

// This tests if we can receive a proper INVITE message
func TestRecvInvite(t *testing.T) {
	rh1_addr := "localhost:4000"
	rh2_addr := "localhost:4001"
	rh1, err1 := handler.NewHandler(rh1_addr)
	rh2, err2 := handler.NewHandler(rh2_addr)
	if err1 != nil || err2 != nil {
		t.Fatalf("Handler failed to initialize")
	}
	defer rh1.Terminate()
	defer rh2.Terminate()
	smsg := &api.REDMessage{Type: api.MessageType_INVITE, Sender: rh1_addr, Receipient: rh2_addr}
	rh1.Send(smsg, rh2_addr)
	rmsg := <-rh2.M
	if rmsg.Type != api.MessageType_INVITE {
		t.Errorf("Received message type is incorrect")
	}
	if rmsg.Sender != rh1_addr {
		t.Errorf("Received message sender is incorrect")
	}
	if rmsg.Receipient != rh2_addr {
		t.Errorf("Received message receipient is incorrect")
	}
}

// This tests if we can receive a proper EDIT message of type EDIT
func TestRecvEditINSERT(t *testing.T) {
	rh1_addr := "localhost:4002"
	rh2_addr := "localhost:4003"
	rh1, err1 := handler.NewHandler(rh1_addr)
	rh2, err2 := handler.NewHandler(rh2_addr)
	if err1 != nil || err2 != nil {
		t.Fatalf("Handler failed to initialize")
	}
	defer rh1.Terminate()
	defer rh2.Terminate()
	e := &api.Edit{Type: api.EditType_INSERT, Pos: 103, Char: 256}
	smsg := &api.REDMessage{Type: api.MessageType_EDIT, Sender: rh1_addr, Receipient: rh2_addr, Edit: e}
	rh1.Send(smsg, rh2_addr)
	rmsg := <-rh2.M
	if rmsg.Type != api.MessageType_EDIT {
		t.Errorf("Received message type is incorrect")
	}
	if rmsg.Sender != rh1_addr {
		t.Errorf("Received message sender is incorrect")
	}
	if rmsg.Receipient != rh2_addr {
		t.Errorf("Received message receipient is incorrect")
	}
	re := rmsg.Edit
	if re.Type != e.Type || re.Pos != e.Pos || re.Char != e.Char {
		t.Errorf("Received Edit is incorrect")
	}
}

// This test if we can receive a proper EDIT message of type DELETE
func TestRecvEditDELETE(t *testing.T) {
	rh1_addr := "localhost:4004"
	rh2_addr := "localhost:4005"
	rh1, err1 := handler.NewHandler(rh1_addr)
	rh2, err2 := handler.NewHandler(rh2_addr)
	if err1 != nil || err2 != nil {
		t.Fatalf("Handler failed to initialize")
	}
	defer rh1.Terminate()
	defer rh2.Terminate()
	e := &api.Edit{Type: api.EditType_DELETE, Pos: 0, Char: 0}
	smsg := &api.REDMessage{Type: api.MessageType_EDIT, Sender: rh1_addr, Receipient: rh2_addr, Edit: e}
	rh1.Send(smsg, rh2_addr)
	rmsg := <-rh2.M
	if rmsg.Type != api.MessageType_EDIT {
		t.Errorf("Received message type is incorrect")
	}
	if rmsg.Sender != rh1_addr {
		t.Errorf("Received message sender is incorrect")
	}
	if rmsg.Receipient != rh2_addr {
		t.Errorf("Received message receipient is incorrect")
	}
	re := rmsg.Edit
	if re.Type != e.Type || re.Pos != e.Pos || re.Char != e.Char {
		t.Errorf("Received Edit is incorrect")
	}
}
