package contentwriter

import (
	_ "embed"
	"net/http"
	"strings"

	"github.com/meox/meox.dev/token"
)

//go:embed resources/style.txt
var style []byte

type HtmlWriter struct {
	w           http.ResponseWriter
	bodyWritten bool
}

func NewHtml(w http.ResponseWriter) *HtmlWriter {
	return &HtmlWriter{
		w: w,
	}
}

func (h *HtmlWriter) ContentType() {
	h.w.Header().Set("Content-Type", "text/html; charset=utf-8")
}

func (h *HtmlWriter) Header() error {
	var rs strings.Builder
	rs.WriteString("<!doctype html>\n\n")
	rs.WriteString("<html lang=\"en\">")
	rs.WriteString("<head>\n")
	rs.WriteString("<meta charset=\"utf-8\">")
	rs.Write(style)
	rs.WriteString("</head>")

	_, err := h.w.Write([]byte(rs.String()))
	return err
}

func (h *HtmlWriter) Close() string {
	return "<br/></body></html>"
}

func (h *HtmlWriter) Write(parsed token.Token) string {
	var rs strings.Builder

	if !h.bodyWritten {
		rs.WriteString("<body>")
		h.bodyWritten = true
	}

	switch parsed.Type {
	case token.Title:
		rs.WriteString("<h2>")
		rs.WriteString(parsed.Value)
		rs.WriteString("</h2>")
	case token.Subtitle:
		rs.WriteString("<h3>")
		rs.WriteString(parsed.Value)
		rs.WriteString("</h3>")
	case token.JobTitle:
		rs.WriteString("<p style='font-size:20pt;'>")
		rs.WriteString(parsed.Value)
		rs.WriteString("</p>")
	case token.Text:
		rs.WriteString(parsed.Value)
		rs.WriteString("<br>")
	case token.Email:
		rs.WriteString("&#128237;")
		rs.WriteString(" " + parsed.Value)
		rs.WriteString("<br>")
	case token.Mobile:
		rs.WriteString("&#9742;")
		rs.WriteString(" " + parsed.Value)
		rs.WriteString("<br>")
	case token.Github:
		rs.WriteString("&#128025;")
		rs.WriteString("<a href=\"" + parsed.Value + "\" target=\"_blank\"> " + parsed.Value + "</a>")
		rs.WriteString("<br>")
	case token.Town:
		rs.WriteString("&#127968;")
		rs.WriteString(" " + parsed.Value)
		rs.WriteString("<br>")
	case token.Web:
		rs.WriteString("&#127760;")
		rs.WriteString("<a href=\"" + parsed.Value + "\" target=\"_blank\"> " + parsed.Value + "</a>")
		rs.WriteString("<br>")
	case token.Empty:
		rs.WriteString("<br>")
	}

	return rs.String()
}
