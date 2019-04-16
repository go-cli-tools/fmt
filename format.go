package format

import (
	"fmt"
	"github.com/go-cli-tools/format/colour"
	"github.com/go-cli-tools/format/style"
	"io"
)

type Format struct {
	style    style.Style
	fgColour colour.Colour
	bgColour colour.Colour
}

func (f *Format) ResetStyle() *Format {
	f.style = style.DefaultStyle
	return f
}

func (f *Format) ResetColour() *Format {
	f.fgColour = colour.DefaultColour
	return f
}

func (f *Format) WithStyle(style style.Style) *Format {
	f.style = style
	return f
}

func (f *Format) WithColour(colour colour.Colour) *Format {
	f.fgColour = colour
	return f
}

func (f *Format) WithBackground(color colour.Colour) *Format {
	f.bgColour = color + colour.BackgroundOffset
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
	return fmt.Print(fmt.Sprintf("\033[%sm%v\033[%sm", f.code(), line, f.endCode()))
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

func (f *Format) code() string {
	return fmt.Sprintf("%d;%d;%d", f.style, f.fgColour, f.bgColour)
}

func (f *Format) endCode() string {
	return fmt.Sprintf("%d", style.ResetAll)
}
