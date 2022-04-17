package token

type Type int

const (
	Title Type = iota
	Subtitle
	Text
	Email
	Town
	Web
	Github
	Mobile
	Empty
	JobTitle
)

type Token struct {
	Type  Type
	Value string
}
