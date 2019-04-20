// +build windows

package colour

import (
	"syscall"
	"unsafe"
)

type Colour uint16

const (
	Black Colour = iota
	Blue
	Green
	Cyan
	Red
	Magenta
	Yellow
	White
	BrightBlack
	BrightBlue
	BrightGreen
	BrightCyan
	BrightRed
	BrightMagenta
	BrightYellow
	BrightWhite
)

const (
	BackgroundOffset  Colour = 16
	DefaultColour            = Black | White
	WinFgIntensity    Colour = 0x0008
	WinBgIntensity    Colour = 0x0080
	WinOpLeading      Colour = 0x0100
	WinOpTrailing     Colour = 0x0200
	WinOpHorizontal   Colour = 0x0400
	WinOpReverse      Colour = 0x4000
	WinOpUnderscore   Colour = 0x8000
	CustomColour      Colour = 38
	TwoFiftySixColour Colour = 5
	RgbColour         Colour = 2
)

// refer: golang.org/x/sys/windows
type (
	// coordinate cursor position coordinates
	coordinate struct {
		x int16
		y int16
	}

	rectangle struct {
		left   int16
		top    int16
		right  int16
		bottom int16
	}

	// Used with GetConsoleScreenBuffer to retrieve information about a console
	// screen buffer. See
	// https://docs.microsoft.com/en-us/windows/console/console-screen-buffer-info-str
	// for details.
	consoleScreenBufferInfo struct {
		Size              coordinate
		CursorPosition    coordinate
		Attributes        uint16 // is windows color setting
		Window            rectangle
		MaximumWindowSize coordinate
	}
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
	DefScreenInfo consoleScreenBufferInfo
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
	_ = getConsoleScreenBufferInfo(uintptr(syscall.Stdout), &DefScreenInfo)
}

// from package: golang.org/x/sys/windows
func getConsoleScreenBufferInfo(consoleOutput uintptr, info *consoleScreenBufferInfo) (err error) {
	r1, _, e1 := syscall.Syscall(procGetConsoleScreenBufferInfo.Addr(), 2, consoleOutput, uintptr(unsafe.Pointer(info)), 0)
	if r1 == 0 {
		if e1 != 0 {
			err = e1
		} else {
			err = syscall.EINVAL
		}
	}

	return
}

func SetConsoleTextAttr(consoleOutput uintptr, winAttr uint16) (n int, err error) {
	ret, _, err := procSetTextAttribute.Call(consoleOutput, uintptr(winAttr))

	return int(ret), err
}
