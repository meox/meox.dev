package cv_parser

import (
	"bufio"
	"errors"
	"os"
	"strings"

	"github.com/meox/meox.dev/contentwriter"
	"github.com/meox/meox.dev/token"
)

type MDFile struct {
	fd *os.File
}

func Open(filename string) (*MDFile, error) {
	fd, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	return &MDFile{fd: fd}, nil
}

func (md *MDFile) Close() error {
	return md.fd.Close()
}

func (md *MDFile) Parse(c contentwriter.ContentWriter) (string, error) {
	if md.fd == nil {
		return "", errors.New("cv is not readable")
	}

	var rs strings.Builder

	body := bufio.NewScanner(md.fd)
	for body.Scan() {
		// retrieve the line
		line := body.Text()
		line = strings.TrimSpace(line)

		parsed := parseLine(line)
		s := c.Write(parsed)
		rs.WriteString(s)
	}

	rs.WriteString(c.Close())
	return rs.String(), nil
}

func parseLine(s string) token.Token {
	if s == "" {
		return token.Token{Type: token.Empty}
	}
	if strings.HasPrefix(s, "# ") {
		return token.Token{Type: token.Title, Value: strings.TrimPrefix(s, "# ")}
	}
	if strings.HasPrefix(s, "## ") {
		return token.Token{Type: token.Subtitle, Value: strings.TrimPrefix(s, "## ")}
	}
	if strings.HasPrefix(s, ":email: ") {
		return token.Token{Type: token.Email, Value: strings.TrimPrefix(s, ":email: ")}
	}
	if strings.HasPrefix(s, ":github: ") {
		return token.Token{Type: token.Github, Value: strings.TrimPrefix(s, ":github: ")}
	}
	if strings.HasPrefix(s, ":mobile: ") {
		return token.Token{Type: token.Mobile, Value: strings.TrimPrefix(s, ":mobile: ")}
	}
	if strings.HasPrefix(s, ":town: ") {
		return token.Token{Type: token.Town, Value: strings.TrimPrefix(s, ":town: ")}
	}
	if strings.HasPrefix(s, ":web: ") {
		return token.Token{Type: token.Web, Value: strings.TrimPrefix(s, ":web: ")}
	}
	if strings.HasPrefix(s, ":job-title: ") {
		return token.Token{Type: token.JobTitle, Value: strings.TrimPrefix(s, ":job-title: ")}
	}
	return token.Token{Type: token.Text, Value: s}
}
