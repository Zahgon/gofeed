package shared

import (
	"io"
)

func NewReaderLabel(label string, input io.Reader) (io.Reader, error) {
	_ = "STUB: not implemented"
	return *new(io.Reader), nil
}

// Wrap the charset decoder reader with a XML sanitizer
//clean := NewXMLSanitizerReader(conv)
