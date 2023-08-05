package api

type REDServer interface {
	// Invites a peer to their editing session using the given address. It returns an error if it the invitation was not possible.
	Invite(addr string) error

	// Accepts an invitation from a peer. Upon accepting, the sender of the invite is added as a peer to the receipient's list of peers.
	Accept()

	// Notifies all peers within this server's editing session of an update to the text editor.
	// It sends an EDIT message containing the edits
	Notify(text string)

	// Returns the updates channel which is a channel that sends text updates to the GUI.
	Fetch() (updates chan string)

	// Terminate the server, closing the listening socket on this server.
	Terminate()
}
