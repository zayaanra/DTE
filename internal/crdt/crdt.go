package crdt

import "github.com/zayaanra/RED/api"

// A "pair" represents the combination of an index and pid for a single edit in a document.
// The char represents the character at this position.
// The pos represents the position at which a character is edited (inserted or where it is deleted).
// The pid is the process ID of the process that placed an edit at that pos.
type Pair struct {
	char byte
	pos  int
	pid  []byte
}

// A document consists of pairs. It can have as many pairs as necessary. Pairs may be deleted if the user performs a delete operation.
// The counter denotes the position at which the next character will be inserted at. It only ever increments.
type CRDT struct {
	counter int
	pairs   []*Pair
}

func NewCRDT() *CRDT {
	crdt := CRDT{0, make([]*Pair, 0)}
	return &crdt
}

// Increments the document's counter.
func (crdt *CRDT) Increment() {
	crdt.counter += 1
}

// Updates the CRDT based on the operation. Upon completion, it returns the document as a string for the GUI to receive as an update.
func (crdt *CRDT) UpdateCRDT(rmsg *api.REDMessage) string {
	e := rmsg.Edit
	switch e.Type {
	case api.EditType_INSERT:
		return crdt.Insert(byte(e.Char), int(e.Pos), nil)
	}
	return ""
}

// Insert the given character at the specified position.
func (crdt *CRDT) Insert(char byte, pos int, pid []byte) string {
	pair := &Pair{char, pos, pid}
	crdt.pairs = append(crdt.pairs, pair)
	// TODO - We're assuming there is no conflict for now. So let's send the updated text anyways.
	return crdt.Stringify()
}

// Converts the document into a string. Useful for debugging.
func (crdt *CRDT) Stringify() string {
	bytes := make([]byte, len(crdt.pairs))
	for _, pair := range crdt.pairs {
		bytes = append(bytes, pair.char)
	}
	return string(bytes)
}
