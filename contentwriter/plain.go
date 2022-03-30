package contentwriter

import (
	"net/http"
	"strings"

	"github.com/meox/meox.dev/color"
	"github.com/meox/meox.dev/token"
)

type PlainWriter struct {
	w http.ResponseWriter
}

func NewPlain(w http.ResponseWriter) *PlainWriter {
	return &PlainWriter{
		w: w,
	}
}

func (p *PlainWriter) ContentType() {
	p.w.Header().Set("Content-Type", "plain/text; charset=utf-8")
}

func (p *PlainWriter) Header() error {
	return nil
}

func (p *PlainWriter) Close() string {
	return "\n"
}

func (p *PlainWriter) Write(parsed token.Token) string {
	var rs strings.Builder

	switch parsed.Type {
	case token.Title:
		rs.WriteString(color.Colorize("# ", 67, 97, 238))
		rs.WriteString(color.Colorize(parsed.Value, 247, 37, 133))
	case token.Subtitle:
		rs.WriteString(color.Colorize("## ", 67, 97, 238))
		rs.WriteString(color.Colorize(parsed.Value, 239, 71, 111))
	case token.JobTitle:
		rs.WriteString(color.Colorize(parsed.Value, 147, 129, 255))
	case token.Text:
		rs.WriteString(parsed.Value)
	case token.Email:
		rs.Write([]byte{0xE2, 0x9C, 0x89})
		rs.WriteString(" " + parsed.Value)
	case token.Mobile:
		rs.Write([]byte{0xE2, 0x98, 0x8E})
		rs.WriteString(" " + parsed.Value)
	case token.Github:
		rs.Write([]byte{0xF0, 0x9F, 0x90, 0xB1})
		rs.WriteString(" " + parsed.Value)
	case token.Town:
		rs.Write([]byte{0xE2, 0x8C, 0x82})
		rs.WriteString(" " + parsed.Value)
	case token.Web:
		rs.Write([]byte{0xF0, 0x9F, 0x8C, 0x90})
		rs.WriteString(" " + parsed.Value)
	case token.ElementList:
		rs.Write([]byte{0xE1, 0x9B, 0xAB})
		rs.WriteString(" " + parsed.Value)
	case token.Empty:
		rs.WriteString("")
	}

	rs.WriteString("\n")
	return rs.String()
}
