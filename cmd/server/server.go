package main

import (
	"log"
	"os"

	"github.com/therecipe/qt/widgets"
)

// This command starts a server (e.g. a peer) in our network.
// Any peer is capable of sending/receiving a message.
func main() {
	// To test locally, the user should input two arguments to the command line (executable and the port, e.g. "localhost:3000")
	// Port 8080 CANNOT be used because it is already used to serve HTTP requests. If you get an error under that port, that's why.
	app := widgets.NewQApplication(len(os.Args), os.Args)

	window := widgets.NewQMainWindow(nil, 0)
	window.SetWindowTitle("Server Configuration")
	window.SetMinimumSize2(400, 300)

	layout := widgets.NewQVBoxLayout()
	widget := widgets.NewQWidget(nil, 0)
	widget.SetLayout(layout)
	window.SetCentralWidget(widget)

	label := widgets.NewQLabel2("Server Response:", nil, 0)
	layout.AddWidget(label, 0, 0)

	button := widgets.NewQPushButton2("Send Request", nil)
	layout.AddWidget(button, 0, 0)

	serverButton := widgets.NewQPushButton2("Start Server", nil)
	layout.AddWidget(serverButton, 0, 0)

	serverButton.ConnectClicked(func(bool) {
		log.Println("Server start clicked")
		serverButton.SetEnabled(false)
	})

	// Start the GUI
	window.Show()
	app.Exec()

	// TODO - everything below here is the actual server code, it will be used once user clicks server start button
	// if len(os.Args) != 2 {
	// 	fmt.Fprintf(os.Stderr, "usage: ./server port")
	// } else {
	// 	// Parse the port and start a REDServer under that port
	// 	addr := os.Args[1]
	// 	exit := make(chan bool)
	// 	rs, err := red.NewREDServer(addr, exit)
	// 	if err != nil {
	// 		fmt.Fprintf(os.Stderr, "REDServer failed to start up")
	// 		return
	// 	}
	// 	// // Read from stdin until we hit EOF. Then, we can safely terminate our REDServer.
	// 	// os.Stdin.Read(make([]byte, 1))
	// 	<-exit
	// 	rs.Terminate()
	// }
}
