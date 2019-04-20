package format

import (
	"fmt"
	c "github.com/go-cli-tools/format/colour"
	s "github.com/go-cli-tools/format/style"
	"io"
)

type Format struct {
	style         s.Style
	fgColour      c.Colour
	bgColour      c.Colour
	fgColour256   int
	bgColour256   int
	fgR, fgG, fgB int
	bgR, bgG, bgB int
}

func (f *Format) ResetStyle() *Format {
	f.style = s.DefaultStyle
	return f
}

func (f *Format) ResetColour() *Format {
	f.fgColour = c.DefaultColour
	return f
}

func (f *Format) WithStyle(style s.Style) *Format {
	f.style = style
	return f
}

func (f *Format) WithColour(colour c.Colour) *Format {
	f.fgColour = colour
	return f
}

func (f *Format) WithColour256(colour int) *Format {
	f.fgColour256 = colour
	return f
}

func (f *Format) WithColourRgb(red, green, blue int) *Format {
	f.fgColour = c.RgbColour
	f.fgR, f.fgG, f.fgB = red, green, blue
	return f
}

func (f *Format) WithBackground(colour c.Colour) *Format {
	f.bgColour = colour
	return f
}

func (f *Format) WithBackground256(colour int) *Format {
	f.bgColour = c.TwoFiftySixColour
	f.bgColour256 = colour
	return f
}

func (f *Format) WithBackgroundRgb(red, green, blue int) *Format {
	f.bgColour = c.RgbColour
	f.bgR, f.bgG, f.bgB = red, green, blue
	return f
}

func (f *Format) Fprint(w io.Writer, a ...interface{}) (n int, err error) {
	line := fmt.Sprint(a...)
	return fmt.Fprint(w, fmt.Sprintf("\033[%sm%v\033[%sm", f.code(), line, f.endCode()))
}

func (f *Format) Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error) {
	line := fmt.Sprintf(format, a...)
	return fmt.Fprint(w, fmt.Sprintf("\033[%sm%v\033[%sm", f.code(), line, f.endCode()))
}

func (f *Format) Fprintln(w io.Writer, a ...interface{}) (n int, err error) {
	line := fmt.Sprint(a...)
	return fmt.Fprint(w, fmt.Sprintf("\033[%sm%v\033[%sm\n", f.code(), line, f.endCode()))
}

func (f *Format) Print(a ...interface{}) (n int, err error) {
	line := fmt.Sprint(a...)
	return fmt.Print(fmt.Sprintf("\033[%sm%v\033[%sm", f.code(), line, f.endCode()))
}

func (f *Format) Printf(format string, a ...interface{}) (n int, err error) {
	line := fmt.Sprintf(format, a...)
	//return fmt.Print(fmt.Sprintf("\033[%sm%v\033[%sm", f.code(), line, f.endCode()))
	f.code()
	aa, bb := fmt.Print(line)
	f.endCode()

	return aa, bb
}

func (f *Format) Println(a ...interface{}) (n int, err error) {
	line := fmt.Sprint(a...)
	return fmt.Printf(fmt.Sprintf("\033[%sm%v\033[%sm\n", f.code(), line, f.endCode()))
}

func (f *Format) Sprint(a ...interface{}) string {
	line := fmt.Sprint(a...)
	return fmt.Sprintf("\033[%sm%v\033[%sm", f.code(), line, f.endCode())
}

func (f *Format) Sprintf(format string, a ...interface{}) string {
	line := fmt.Sprintf(format, a...)
	return fmt.Sprintf("\033[%sm%v\033[%sm", f.code(), line, f.endCode())
}

func (f *Format) Sprintln(a ...interface{}) string {
	line := fmt.Sprint(a...)
	return fmt.Sprintf("\033[%sm%v\033[%sm\n", f.code(), line, f.endCode())
}
