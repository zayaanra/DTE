package main

import (
	"fmt"
	"log"
	"os"

	"github.com/therecipe/qt/widgets"
	"github.com/zayaanra/RED/api"
	"github.com/zayaanra/RED/internal/red"
)

// Creates the REDServer and begins listening for incoming messages
func run(comms chan<- api.REDServer) {
	// TODO - everything below here is the actual server code, it will be used once user clicks server start button
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: ./server port")
	} else {
		// Parse the port and start a REDServer under that port
		addr := os.Args[1]
		rs, err := red.NewREDServer(addr)
		if err != nil {
			fmt.Fprintf(os.Stderr, "REDServer failed to start up")
			comms <- nil
			return
		}
		comms <- rs

		// Read from stdin until we hit EOF. Then, we can safely terminate our REDServer.
		os.Stdin.Read(make([]byte, 1))
		log.Println("Terminating server...")
		close(comms)
		rs.Terminate()
	}
}

// Boot REDServer (initialization)
func boot() api.REDServer {
	comms := make(chan api.REDServer)

	// Run the server
	log.Println("Running server...")
	go run(comms)

	rs := <-comms
	return rs
}

// This command starts a server (e.g. a peer) in our network.
// Any peer is capable of sending/receiving a message.
func main() {
	rs := boot()
	if rs == nil {
		os.Exit(1)
	}

	// To test locally, the user should input two arguments to the command line (executable and the port, e.g. "localhost:3000")
	app := widgets.NewQApplication(0, nil)

	window := widgets.NewQMainWindow(nil, 0)
	window.SetWindowTitle("RED Editor")
	window.SetMinimumSize2(400, 300)

	layout := widgets.NewQVBoxLayout()
	widget := widgets.NewQWidget(nil, 0)
	widget.SetLayout(layout)
	window.SetCentralWidget(widget)

	document := widgets.NewQPlainTextEdit2("", nil)
	document.ConnectTextChanged(func() {
		s := document.ToPlainText()
		log.Println(s)
	})

	invite := widgets.NewQPushButton2("Invite", nil)
	invite.ConnectClicked(func(bool) {
		configWindow := widgets.NewQDialog(nil, 0)
		configWindow.SetWindowTitle("Invite Configuration")
		configWindow.SetModal(true)
		configWindow.SetFixedSize2(300, 100)

		layout := widgets.NewQVBoxLayout()

		hostLabel := widgets.NewQLabel2("Host", nil, 0)
		hostEntry := widgets.NewQLineEdit(nil)
		layout.AddWidget(hostLabel, 0, 0)
		layout.AddWidget(hostEntry, 0, 0)

		portLabel := widgets.NewQLabel2("Port", nil, 0)
		portEntry := widgets.NewQLineEdit(nil)
		layout.AddWidget(portLabel, 0, 0)
		layout.AddWidget(portEntry, 0, 0)

		btn := widgets.NewQPushButton2("Invite", nil)
		btn.ConnectClicked(func(bool) {
			addr := hostEntry.Text() + ":" + portEntry.Text()
			rs.Invite(addr)
		})
		layout.AddWidget(btn, 0, 0)

		configWindow.SetLayout(layout)
		configWindow.Show()
	})

	layout.AddWidget(invite, 0, 0)
	layout.AddWidget(document, 0, 0)

	// Start the GUI
	window.Show()
	app.Exec()

}
