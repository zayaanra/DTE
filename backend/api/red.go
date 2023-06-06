package api

type REDServer interface {
	// Sends the given message to the given peer
	// It returns an error if sending the message failed and nil if it succeeded
	Send(peer string) error

	// Reads a single message sent to this server's channel
	Recv() 
	
	// Terminate the server, closing the listening socket
	// on this server.
	Terminate()
}
