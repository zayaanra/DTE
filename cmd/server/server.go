package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gotk3/gotk3/gtk"
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

func main() {
	rs := boot()
	if rs == nil {
		os.Exit(1)
	}

	gtk.Init(nil)

	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Fatal("Unable to create window:", err)
	}
	win.SetTitle("RED Editor")
	win.Connect("destroy", gtk.MainQuit)
	win.SetDefaultSize(800, 600)

	// Create a vertical box container
	vBox, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 5)
	if err != nil {
		log.Fatal("Unable to create box:", err)
	}

	// Create a toolbar
	toolbar, err := gtk.ToolbarNew()
	if err != nil {
		log.Fatal("Unable to create toolbar:", err)
	}
	// Add toolbar actions here (New document, Cut, Copy, Paste, Invite, etc.)
	vBox.PackStart(toolbar, false, false, 0)

	// Create a scrolled window
	scrolledWindow, err := gtk.ScrolledWindowNew(nil, nil)
	if err != nil {
		log.Fatal("Unable to create scrolled window:", err)
	}
	vBox.PackStart(scrolledWindow, true, true, 0)

	// Create a multiline text view
	textView, err := gtk.TextViewNew()
	if err != nil {
		log.Fatal("Unable to create text view:", err)
	}
	textBuffer, _ := textView.GetBuffer()
	scrolledWindow.Add(textView)

	// Create an "Invite" button
	inviteButton, err := gtk.ButtonNewWithLabel("Invite")
	if err != nil {
		log.Fatal("Unable to create button:", err)
	}
	vBox.PackStart(inviteButton, false, false, 0)

	// Handle the "Invite" button click event
	inviteButton.Connect("clicked", func() {
		// Implement your Invite button logic here
		fmt.Println("Invite button clicked!")
	})

	// Connect "changed" signal of the text buffer
	textBuffer.Connect("changed", func() {
		// Handle text buffer changes here
		// You can get the text from the buffer using textBuffer.GetText()
	})

	// Add the vertical box container to the window
	win.Add(vBox)

	win.ShowAll()
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
