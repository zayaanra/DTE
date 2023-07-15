package api

type Session interface {
	// Closes this editing session. An editing session is closed if a server is closed.
	Close()

	// Updates it's own GUI to match the received EDIT changes.
	Update(text string)
}
