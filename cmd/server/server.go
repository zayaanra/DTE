package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gotk3/gotk3/gtk"
	"github.com/zayaanra/RED/api"
	"github.com/zayaanra/RED/internal/red"
)

func run(comms chan<- api.REDServer, addr string) {
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
	close(comms)
	rs.Terminate()
}

func boot(addr string) api.REDServer {
	comms := make(chan api.REDServer)

	// Run the server
	log.Println("Running server...")
	go run(comms, addr)

	rs := <-comms
	return rs
}

func refresh(doc *gtk.Entry, rs api.REDServer) {
	updates := rs.Fetch()
	for text := range updates {
		doc.SetText(text)
	}
}

func findDifference(old, new string) int {
	var i int
	for i = 0; i < len(old) && i < len(new); i++ {
	}
	return i
}

// TODO - Need to send EDIT message to peers
func documentUpdated(buffer *gtk.TextBuffer) {
	startIter, endIter := buffer.GetBounds()
	text, _ := buffer.GetText(startIter, endIter, false)
	fmt.Println("Text changed:", text)
}

func main() {
	gtk.Init(nil)

	builder, err := gtk.BuilderNewFromFile("cmd/server/gui.glade")
	if err != nil {
		log.Printf("failed to load glade file - %v\n", err)
	}

	obj, err := builder.GetObject("main")
	if err != nil {
		log.Printf("failed to get main window - %v\n", err)
	}

	mainWindow, ok := obj.(*gtk.Window)
	if !ok {
		log.Fatal("Error casting to window")
	}
	mainWindow.Connect("destroy", func() {
		gtk.MainQuit()
	})

	obj, _ = builder.GetObject("startBtn")
	startBtn, _ := obj.(*gtk.Button)

	obj, _ = builder.GetObject("hostEntry")
	hostEntry, _ := obj.(*gtk.Entry)
	obj, _ = builder.GetObject("portEntry")
	portEntry, _ := obj.(*gtk.Entry)

	var rs api.REDServer

	startBtn.Connect("clicked", func() {
		host, _ := hostEntry.GetText()
		port, _ := portEntry.GetText()
		addr := host + ":" + port
		log.Printf("Attempting to boot server under %s", addr)

		rs = boot(addr)
		if rs == nil {
			os.Exit(1)
		}

		obj, _ := builder.GetObject("docWindow")
		docWindow, _ := obj.(*gtk.Window)

		obj, _ = builder.GetObject("textView")
		textView, _ := obj.(*gtk.TextView)
		buffer, _ := gtk.TextBufferNew(nil)
		textView.SetBuffer(buffer)
		buffer.Connect("changed", func() {
			documentUpdated(buffer)
		})
		docWindow.ShowAll()
	})

	obj, _ = builder.GetObject("inviteBtn")
	inviteBtn, _ := obj.(*gtk.ToolButton)
	inviteBtn.Connect("clicked", func() {
		obj, _ = builder.GetObject("inviteWindow")
		inviteWindow := obj.(*gtk.Window)
		inviteWindow.ShowAll()

		obj, _ = builder.GetObject("sendInvBtn")
		sendInvBtn := obj.(*gtk.Button)

		sendInvBtn.Connect("clicked", func() {
			// Invite the peer with the given peer address
			obj, _ = builder.GetObject("hostEntry2")
			hostEntry, _ := obj.(*gtk.Entry)
			obj, _ = builder.GetObject("portEntry2")
			portEntry, _ := obj.(*gtk.Entry)
			host, _ := hostEntry.GetText()
			port, _ := portEntry.GetText()
			addr := host + ":" + port
			rs.Invite(addr)
			inviteWindow.Close()
		})
	})

	mainWindow.ShowAll()

	gtk.Main()

}

/*
package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/zayaanra/RED/api"
	"github.com/zayaanra/RED/internal/red"
)

func run(comms chan<- api.REDServer) {
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

func findDifference(old, new string) int {
	var i int
	for i = 0; i < len(old) && i < len(new); i++ {
	}
	return i
}

func main() {
	rs := boot()
	if rs == nil {
		os.Exit(1)
	}

	a := app.New()
	w := a.NewWindow("RED Editor")
	w.Resize(fyne.NewSize(500, 500))

	document := widget.NewMultiLineEntry()
	documentContainer := container.NewScroll(document)
	documentContainer.Resize(fyne.NewSize(1000, 1000))

	oldText := ""

	document.OnChanged = func(s string) {
		// TODO - local GUI not updating because calling doc.setText causes a document.onChanged resulting in an infinite Notify() broadcast
		newText := strings.TrimSpace(s)
		cursorPos := findDifference(oldText, newText)
		var editType int
		if len(newText) > len(oldText) {
			editType = int(api.EditType_INSERT)
		} else {
			editType = int(api.EditType_DELETE)
		}
		var char byte
		if cursorPos == 0 {
			char = 0x0
		} else {
			char = byte(newText[cursorPos])
		}
		rs.Notify(char, cursorPos, editType)
	}

	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(theme.DocumentCreateIcon(), func() {
			log.Println("New document")
		}),
		widget.NewToolbarSeparator(),
		widget.NewToolbarAction(theme.ContentCutIcon(), func() {}),
		widget.NewToolbarAction(theme.ContentCopyIcon(), func() {}),
		widget.NewToolbarAction(theme.ContentPasteIcon(), func() {}),
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.MailSendIcon(), func() {
			configWindow := a.NewWindow("Invite Configuration")
			size := fyne.NewSize(300, 150)
			configWindow.Resize(size)

			hostLabel := widget.NewLabel("Host")
			hostEntry := widget.NewEntry()
			portLabel := widget.NewLabel("Port")
			portEntry := widget.NewEntry()
			inviteBtn := widget.NewButton("Invite", func() {
				addr := hostEntry.Text + ":" + portEntry.Text
				rs.Invite(addr)
				configWindow.Close()
			})

			container := container.NewVBox(
				hostLabel, hostEntry,
				portLabel, portEntry,
				inviteBtn,
			)
			configWindow.SetContent(container)
			configWindow.Show()
		}),
	)

	content := fyne.NewContainerWithLayout(
		layout.NewBorderLayout(toolbar, nil, nil, nil),
		toolbar,
		documentContainer,
	)
	w.SetContent(content)

	go refresh(document, rs)

	w.ShowAndRun()
}
*/
