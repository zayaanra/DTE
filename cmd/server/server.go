package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/zayaanra/RED/api"
	"github.com/zayaanra/RED/internal/red"
)

func run(comms chan<- api.REDServer) {
	// TODO - everything below here is the actual server code, it will be used once the user clicks the server start button
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: ./server port")
	} else {
		// Parse the port and start a REDServer under that port
		addr := os.Args[1]
		updates := make(chan string)
		rs, err := red.NewREDServer(addr, updates)
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

func boot() api.REDServer {
	comms := make(chan api.REDServer)

	// Run the server
	log.Println("Running server...")
	go run(comms)

	rs := <-comms
	return rs
}

func refresh(doc *widget.Entry, rs api.REDServer) {
	updates := rs.Fetch()
	for text := range updates {
		doc.SetText(text)
	}
}

func main() {
	rs := boot()
	if rs == nil {
		os.Exit(1)
	}

	a := app.New()
	w := a.NewWindow("RED Editor")

	document := widget.NewMultiLineEntry()
	document.OnChanged = func(s string) {
		rs.Notify(strings.TrimSpace(s))
	}

	invite := widget.NewButton("Invite", func() {
		configWindow := a.NewWindow("Invite Configuration")
		size := fyne.NewSize(300, 150)
		configWindow.Resize(size)

		hostLabel := widget.NewLabel("Host")
		hostEntry := widget.NewEntry()
		portLabel := widget.NewLabel("Port")
		portEntry := widget.NewEntry()
		inviteBtn := widget.NewButton("Invite", func() {
			addr := hostEntry.Text + ":" + portEntry.Text
			rs.Invite(addr, document)
			configWindow.Close()
		})

		container := container.NewVBox(
			hostLabel, hostEntry,
			portLabel, portEntry,
			inviteBtn,
		)
		configWindow.SetContent(container)
		configWindow.Show()
	})

	content := container.NewVBox(invite, document)
	w.SetContent(content)

	go refresh(document, rs)

	w.ShowAndRun()
}
