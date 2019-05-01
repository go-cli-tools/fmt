// +build windows

package format

import (
	"github.com/go-cli-tools/format/colour"
	"github.com/go-cli-tools/format/style"
)

func new() *Format {
	if colour.DefScreenInfo.Attributes == 86 {
		return &Format{
			style:    style.DefaultStyle,
			fgColour: colour.White,
			bgColour: colour.Magenta,
		}
	} else {
		return &Format{
			style:    style.DefaultStyle,
			fgColour: colour.White,
			bgColour: colour.Black,
		}
	}
}
