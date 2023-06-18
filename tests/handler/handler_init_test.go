package tests

import (
	"testing"

	"github.com/zayaanra/RED/internal/handler"
)

/* handler_init_test.go contains all test cases that test if handler initialization was successful */

// Basic server initialization test.
// It just checks if no error is returned.
func TestHandlerInit(t *testing.T) {
	addr := "localhost:2000"
	rs, err := handler.NewHandler(addr)
	if err != nil {
		t.Fatalf("Handler failed to initialize")
	}
	rs.Terminate()
}

// We should not be able to start a handler with an invalid address.
// This should return an error.
func TestHandlerInitBadAddr(t *testing.T) {
	addr := "badaddr"
	_, err := handler.NewHandler(addr)
	if err == nil {
		t.Fatalf("Handler initialized even w/ a bad address")
	}
}
