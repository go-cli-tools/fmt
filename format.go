package format

import (
	"fmt"
	"io"
)

type Format struct {
	style Style
}

func (f *Format) WithStyle(style Style) *Format {
	f.style = style
	return f
}

func (f *Format) Fprint(w io.Writer, a ...interface{}) (n int, err error) {
	line := fmt.Sprint(a...)
	return fmt.Fprint(w, fmt.Sprintf("\033[%dm%v\033[%dm", f.style, line, f.style+ResetOffset))
}

func (f *Format) Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error) {
	line := fmt.Sprintf(format, a...)
	return fmt.Fprint(w, fmt.Sprintf("\033[%dm%v\033[%dm", f.style, line, f.style+ResetOffset))
}

func (f *Format) Fprintln(w io.Writer, a ...interface{}) (n int, err error) {
	line := fmt.Sprintln(a...)
	return fmt.Fprint(w, fmt.Sprintf("\033[%dm%v\033[%dm", f.style, line, f.style+ResetOffset))
}

func (f *Format) Print(a ...interface{}) (n int, err error) {
	line := fmt.Sprint(a...)
	return fmt.Print(fmt.Sprintf("\033[%dm%v\033[%dm", f.style, line, f.style+ResetOffset))
}

func (f *Format) Printf(format string, a ...interface{}) (n int, err error) {
	line := fmt.Sprintf(format, a...)
	return fmt.Print(fmt.Sprintf("\033[%dm%v\033[%dm", f.style, line, f.style+ResetOffset))
}

func (f *Format) Println(a ...interface{}) (n int, err error) {
	line := fmt.Sprintln(a...)
	return fmt.Print(fmt.Sprintf("\033[%dm%v\033[%dm", f.style, line, f.style+ResetOffset))
}

func (f *Format) Sprint(a ...interface{}) string {
	line := fmt.Sprint(a...)
	return fmt.Sprint("\033[%dm%v\033[%dm", f.style, line, f.style+ResetOffset)
}

func (f *Format) Sprintf(format string, a ...interface{}) string {
	line := fmt.Sprintf(format, a...)
	return fmt.Sprint("\033[%dm%v\033[%dm", f.style, line, f.style+ResetOffset)
}

func (f *Format) Sprintln(a ...interface{}) string {
	line := fmt.Sprintln(a...)
	return fmt.Sprint("\033[%dm%v\033[%dm", f.style, line, f.style+ResetOffset)
}
