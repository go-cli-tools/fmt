// +build windows

package colour

type Colour uint16

const (
	BackgroundOffset Colour = 16
	DefaultColour    Colour = Black*BackgroundOffset | White
	WinFgIntensity   Colour = 0x0008
	WinBgIntensity   Colour = 0x0080
	WinOpLeading     Colour = 0x0100
	WinOpTrailing    Colour = 0x0200
	WinOpHorizontal  Colour = 0x0400
	WinOpReverse     Colour = 0x4000
	WinOpUnderscore  Colour = 0x8000
)

const (
	Black Colour = iota
	Blue
	Green
	Aqua
	Red
	Magenta
	Yellow
	White
	BrightBlack
	BrightBlue
	BrightGreen
	BrightAqua
	BrightRed
	BrightMagenta
	BrightYellow
	BrightWhite
)

var (
	// for cmd.exe
	// echo %ESC%[1;33;40m Yellow on black %ESC%[0m
	escChar = ""
	// isMSys bool
	kernel32 *syscall.LazyDLL

	procGetConsoleMode *syscall.LazyProc

	procSetTextAttribute           *syscall.LazyProc
	procGetConsoleScreenBufferInfo *syscall.LazyProc

	// console screen buffer info
	// eg {size:{x:215 y:3000} cursorPosition:{x:0 y:893} attributes:7 window:{left:0 top:882 right:214 bottom:893} maximumWindowSize:{x:215 y:170}}
	defScreenInfo consoleScreenBufferInfo
)

func init() {
	// isMSys = utils.IsMSys()
	kernel32 = syscall.NewLazyDLL("kernel32.dll")

	// https://docs.microsoft.com/en-us/windows/console/setconsolemode
	procGetConsoleMode = kernel32.NewProc("GetConsoleMode")
	// procSetConsoleMode = kernel32.NewProc("SetConsoleMode")

	procSetTextAttribute = kernel32.NewProc("SetConsoleTextAttribute")
	// https://docs.microsoft.com/en-us/windows/console/getconsolescreenbufferinfo
	procGetConsoleScreenBufferInfo = kernel32.NewProc("GetConsoleScreenBufferInfo")

	// fetch console screen buffer info
	getConsoleScreenBufferInfo(uintptr(syscall.Stdout), &defScreenInfo)
}