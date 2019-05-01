// +build !windows

package format

import (
	"github.com/go-cli-tools/format/colour"
	"github.com/go-cli-tools/format/style"
)

func new() *Format {
	return &Format{
		style:    style.DefaultStyle,
		fgColour: colour.DefaultColour,
		bgColour: colour.DefaultColour + colour.BackgroundOffset,
	}
}

