package tests

import (
	"testing"

	"github.com/zayaanra/RED/backend/internal/RED"
)

/* server_init_test.go contains all test cases that test if server initialization was successful */

// Basic server initialization test.
// It just checks if no error is returned.
func TestServerInit(t *testing.T) {
	addr := "localhost:2000"
	rs, err := RED.NewREDServer(addr)
	if err != nil {
		t.Fatalf("Server failed to initialize")
	}
	rs.Terminate()
}

// We should not be able to start a server with an invalid address.
// This should return an error.
func TestServerInitBadAddr(t *testing.T) {
	addr := "badaddr"
	_, err := RED.NewREDServer(addr)
	if err == nil {
		t.Fatalf("Server initialized even w/ a bad address")
	}
}
