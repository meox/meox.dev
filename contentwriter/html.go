package contentwriter

import (
	"net/http"
	"strings"

	"github.com/meox/meox.dev/token"
)

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
	_, err := h.w.Write([]byte("<!doctype html>\n\n<html lang=\"en\"><head>\n<meta charset=\"utf-8\"></head>"))
	return err
}

func (h *HtmlWriter) Close() string {
	return "</body>"
}

func (h *HtmlWriter) Write(t token.Token) string {
	var rs strings.Builder

	if !h.bodyWritten {
		rs.WriteString("<body style=\"background-color: black;color:white\">")
	}

	return rs.String()
}
