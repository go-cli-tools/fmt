package fmt

type Style uint8

const (
	Bold Style = iota
	Light
	_
	Underlined
	Blink
	_
	Inverted
	Hidden
)

const (
	EndBold Style = iota + 20
	EndLight
	_
	EndUnderlined
	EndBlink
	_
	EndInverted
	EndHidden
)
