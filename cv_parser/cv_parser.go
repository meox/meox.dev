package cv_parser

import (
	"bufio"
	"errors"
	"os"
	"strings"

	"github.com/meox/meox.dev/color"
)

type tokenType int

const (
	title tokenType = iota
	subtitle
	text
	email
	elementList
	town
	web
	github
	mobile
	empty
	jobTitle
)

type token struct {
	tokenType tokenType
	value     string
}

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

func (md *MDFile) Parse() (string, error) {
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
		switch parsed.tokenType {
		case title:
			rs.WriteString(color.Colorize("# ", 67, 97, 238))
			rs.WriteString(color.Colorize(parsed.value, 247, 37, 133))
		case subtitle:
			rs.WriteString(color.Colorize("## ", 67, 97, 238))
			rs.WriteString(color.Colorize(parsed.value, 239, 71, 111))
		case jobTitle:
			rs.WriteString(color.Colorize(parsed.value, 147, 129, 255))
		case text:
			rs.WriteString(parsed.value)
		case email:
			rs.Write([]byte{0xE2, 0x9C, 0x89})
			rs.WriteString(" " + parsed.value)
		case mobile:
			rs.Write([]byte{0xE2, 0x98, 0x8E})
			rs.WriteString(" " + parsed.value)
		case github:
			rs.Write([]byte{0xF0, 0x9F, 0x90, 0xB1})
			rs.WriteString(" " + parsed.value)
		case town:
			rs.Write([]byte{0xE2, 0x8C, 0x82})
			rs.WriteString(" " + parsed.value)
		case web:
			rs.Write([]byte{0xF0, 0x9F, 0x8C, 0x90})
			rs.WriteString(" " + parsed.value)
		case elementList:
			rs.Write([]byte{0xE1, 0x9B, 0xAB})
			rs.WriteString(" " + parsed.value)
		case empty:
			rs.WriteString("")
		}

		// parse the color

		// put it
		rs.WriteString("\n")
	}

	rs.WriteString("\n")
	return rs.String(), nil
}

func parseLine(s string) token {
	if s == "" {
		return token{tokenType: empty}
	}
	if strings.HasPrefix(s, "# ") {
		return token{tokenType: title, value: strings.TrimPrefix(s, "# ")}
	}
	if strings.HasPrefix(s, "## ") {
		return token{tokenType: subtitle, value: strings.TrimPrefix(s, "## ")}
	}
	if strings.HasPrefix(s, ":email: ") {
		return token{tokenType: email, value: strings.TrimPrefix(s, ":email: ")}
	}
	if strings.HasPrefix(s, ":github: ") {
		return token{tokenType: github, value: strings.TrimPrefix(s, ":github: ")}
	}
	if strings.HasPrefix(s, ":mobile: ") {
		return token{tokenType: mobile, value: strings.TrimPrefix(s, ":mobile: ")}
	}
	if strings.HasPrefix(s, ":town: ") {
		return token{tokenType: town, value: strings.TrimPrefix(s, ":town: ")}
	}
	if strings.HasPrefix(s, ":web: ") {
		return token{tokenType: web, value: strings.TrimPrefix(s, ":web: ")}
	}
	if strings.HasPrefix(s, ":job-title: ") {
		return token{tokenType: jobTitle, value: strings.TrimPrefix(s, ":job-title: ")}
	}
	if strings.HasPrefix(s, "- ") {
		return token{tokenType: elementList, value: strings.TrimPrefix(s, "- ")}
	}
	return token{tokenType: text, value: s}
}
