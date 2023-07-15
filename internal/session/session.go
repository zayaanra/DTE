package session

import (
	"github.com/therecipe/qt/widgets"
)

type Session struct {
	doc *widgets.QPlainTextEdit
}

func NewSession(doc *widgets.QPlainTextEdit) *Session {
	return &Session{doc: doc}
}

func (s *Session) Close() {
	s.doc = nil
}

func (s *Session) Update(text string) {
	s.doc.SetPlainText(text)
}
