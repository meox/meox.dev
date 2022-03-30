package contentwriter

import "github.com/meox/meox.dev/token"

type ContentWriter interface {
	ContentType()
	Header() error
	Write(token token.Token) string
	Close() string
}
