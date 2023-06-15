package main

import (
	"fmt"
	"os"

	"github.com/zayaanra/RED/backend/internal/red"
)

// This command starts a server (e.g. a peer) in our network.
// Any peer is capable of sending/receiving a message.
func main() {
	// To test locally, the user should input two arguments to the command line (executable and the port, e.g. "localhost:3000")
	// Port 8080 CANNOT be used because it is already used to serve HTTP requests. If you get an error under that port, that's why.
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
		// // Read from stdin until we hit EOF. Then, we can safely terminate our REDServer.
		os.Stdin.Read(make([]byte, 1))
		rs.Terminate()
	}

}
