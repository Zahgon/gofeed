package shared

import (
	"io"
)

// NewXMLSanitizerReader creates an io.Reader that
// wraps another io.Reader and removes illegal xml
// characters from the io stream.
func NewXMLSanitizerReader(xml io.Reader) io.Reader {
	_ = "STUB: not implemented"
	return *new(io.Reader)
}
