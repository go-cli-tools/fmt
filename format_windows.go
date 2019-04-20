// +build windows

package format

import (
	c "github.com/go-cli-tools/format/colour"
	"syscall"
)

func (f *Format) code() string {
	var bgColour uint16
	bgColour

	_, _ = c.SetConsoleTextAttr(uintptr(syscall.Stdout), uint16(f.fgColour|f.bgColour))

	return ""
}

func (f *Format) endCode() string {
	var dc c.Colour

	if c.DefScreenInfo.Attributes == uint16(c.DefaultColour) {
		dc = c.DefaultColour
	} else {
		dc = c.Colour(c.DefScreenInfo.Attributes)
	}
	_, _ = c.SetConsoleTextAttr(uintptr(syscall.Stdout), uint16(dc))

	return ""
}
