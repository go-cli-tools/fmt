package fmt

type Colour uint8

// Foreground Colours

const (
	// Basic Colours

	FgBlack Colour = iota + 30
	FgRed
	FgGreen
	FgYellow
	FgBlue
	FgMagenta
	FgCyan
	FgGray
	FgDefault Colour = 39

	// Aliases

	FgPurple Colour = 35
	FgWhite  Colour = 39
)

const (
	// Extended Colours

	FgDarkGray Colour = iota + 90
	FgLightRed
	FgLightGreen
	FgLightYellow
	FgLightBlue
	FgLightMagenta
	FgLightCyan
	FgLightWhite
)

// Background Colours

const (
	// Basic Colours

	BgBlack Colour = iota + 40
	BgRed
	BgGreen
	BgYellow
	BgBlue
	BgMagenta
	BgCyan
	BgGray
	BgDefault Colour = 49

	// Aliases

	BgPurple Colour = 45
	BgWhite  Colour = 49
)

const (
	// Extended Colours

	BgDarkGray Colour = iota + 100
	BgLightRed
	BgLightGreen
	BgLightYellow
	BgLightBlue
	BgLightMagenta
	BgLightCyan
	BgLightWhite
)
