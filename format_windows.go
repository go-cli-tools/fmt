// +build windows

package format

import (
	c "github.com/go-cli-tools/format/colour"
	"syscall"
)

func (f *Format) code() string {

	_, _ = c.SetConsoleTextAttr(uintptr(syscall.Stdout), uint16(f.fgColour|c.BackgroundOffset*f.bgColour))

	return ""
}

func (f *Format) endCode() string {
	var dc c.Colour

	if c.DefScreenInfo.Attributes != 86 {
		dc = c.DefaultColour
	} else {
		dc = c.DefaultColourPS
	}
	_, _ = c.SetConsoleTextAttr(uintptr(syscall.Stdout), uint16(dc))

	return ""
}
