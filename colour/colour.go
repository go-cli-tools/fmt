// +build !windows

package colour

type Colour uint8

const (
	BackgroundOffset  Colour = 10
	CustomColour      Colour = 38
	TwoFiftySixColour Colour = 5
	RgbColour         Colour = 2
	DefaultColour     Colour = 39
)

const (
	Black Colour = iota + 30
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White
)

const (
	BrightBlack Colour = iota + 90
	BrightRed
	BrightGreen
	BrightYellow
	BrightBlue
	BrightMagenta
	BrightCyan
	BrightWhite
)
