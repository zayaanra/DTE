package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/zayaanra/RED/backend/internal/red"
)

// Serves the front-end HTML content
func serve(addr string) {
	port := strings.Split(addr, ":")[1]
	port = ":" + port
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to serve HTML request - %v", err)
	}
}

// This command starts a server (e.g. a peer) in our network.
// Any peer is capable of sending/receiving a message.
func main() {
	// To test locally, the user should input two arguments to the command line (executable and the port, e.g. "localhost:8080")
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: ./server port")
	} else {
		// Parse the port and start a REDServer under that port
		addr := os.Args[1]
		rs, err := red.NewREDServer(addr)
		if err != nil {
			fmt.Fprintf(os.Stderr, "REDServer failed to start up")
			return
		}
		http.Handle("/", http.FileServer(http.Dir("../../../frontend/")))
		go serve(addr)

		// Read from stdin until we hit EOF. Then, we can safely terminate our REDServer.
		os.Stdin.Read(make([]byte, 1))
		rs.Terminate()
	}

}
