package style

type Style uint8

const ResetAll Style = 0
const DefaultStyle Style = 10

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
