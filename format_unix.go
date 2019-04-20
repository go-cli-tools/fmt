// +build !windows

package format

import (
	"fmt"
	c "github.com/go-cli-tools/format/colour"
	s "github.com/go-cli-tools/format/style"
)

func (f *Format) code() string {
	var fgColourCode, bgColourCode string

	if f.fgColour == c.TwoFiftySixColour {
		fgColourCode = fmt.Sprintf("%d;%d;%d", c.CustomColour, f.fgColour, f.fgColour256)
	} else if f.fgColour == c.RgbColour {
		fgColourCode = fmt.Sprintf("%d;%d;%d;%d;%d", c.CustomColour, f.fgColour, f.fgR, f.fgG, f.fgB)
	} else {
		fgColourCode = fmt.Sprint(f.fgColour)
	}

	if f.bgColour == c.TwoFiftySixColour {
		bgColourCode = fmt.Sprintf("%d;%d;%d", c.CustomColour+c.BackgroundOffset, f.bgColour, f.bgColour256)
	} else if f.bgColour == c.RgbColour {
		bgColourCode = fmt.Sprintf("%d;%d;%d;%d;%d", c.CustomColour+c.BackgroundOffset, f.bgColour, f.bgR, f.bgG, f.bgB)
	} else {
		bgColourCode = fmt.Sprint(f.bgColour + c.BackgroundOffset)
	}

	return fmt.Sprintf("\033[%d;%s;%sm", f.style, fgColourCode, bgColourCode)
}

func (f *Format) endCode() string {
	return fmt.Sprintf("\033[%dm", s.ResetAll)
}
