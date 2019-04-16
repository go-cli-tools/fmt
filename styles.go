package format

type Style uint8

const ResetAll Style = 0
const ResetOffset Style = 20

const (
	Bold Style = iota + 1
	Light
	Italics
	Underlined
	Blink
	_
	Inverted
	Hidden
	StrikeThrough
)
