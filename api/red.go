package api

import (
	"fyne.io/fyne/v2/widget"
	"github.com/therecipe/qt/widgets"
)

type REDServer interface {
	// Invites a peer to their editing session using the given address. It returns an error if it the invitation was not possible.
	// It also opens an editing session.
	Invite(addr string, doc *widget.Entry) error

	// Accepts an invitation from a peer. Upon accepting, the sender of the invite is added as a peer to the receipient's list of peers.
	Accept()

	// Hosts an editing session for this server. It contains the document provided by the GUI.
	Open(doc *widgets.QPlainTextEdit)

	// Notifies all peers within this server's editing session of an update to the text editor.
	// It sends an EDIT message containing the edits
	Notify(text string)

	// Fetches the most recent update needed for the GUI. It returns a bool denoting if it possible to fetch said update.
	// It's only not possible when the server has been terminated.
	Fetch() (updates chan string)

	// Terminate the server, closing the listening socket on this server.
	Terminate()
}
