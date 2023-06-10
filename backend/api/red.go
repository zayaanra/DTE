package api


type REDServer interface {
	// Terminate the server, closing the listening socket
	// on this server.
	Terminate()
}
