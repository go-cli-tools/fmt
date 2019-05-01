package format

import (
	"fmt"
	"github.com/go-cli-tools/format/colour"
	"github.com/go-cli-tools/format/style"
	"io"
)

func WithStyle(style style.Style) *Format {
	return new().WithStyle(style)
}

func WithColour(colour colour.Colour) *Format {
	return new().WithColour(colour)
}

func WithColour256(colour int) *Format {
	return new().WithColour256(colour)
}

func WithColourRgb(r, g, b int) *Format {
	return new().WithColourRgb(r, g, b)
}

func WithBackground(colour colour.Colour) *Format {
	return new().WithBackground(colour)
}

func WithBackground256(colour int) *Format {
	return new().WithBackground256(colour)
}

func WithBackgroundRgb(r, g, b int) *Format {
	return new().WithBackgroundRgb(r, g, b)
}

func Errorf(format string, a ...interface{}) error {
	return fmt.Errorf(format, a...)
}

func Fprint(w io.Writer, a ...interface{}) (n int, err error) {
	return fmt.Fprint(w, a...)
}

func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error) {
	return fmt.Fprintf(w, format, a...)
}

func Fprintln(w io.Writer, a ...interface{}) (n int, err error) {
	return fmt.Fprintln(w, a...)
}

func Fscan(r io.Reader, a ...interface{}) (n int, err error) {
	return fmt.Fscan(r, a...)
}

func Fscanf(r io.Reader, format string, a ...interface{}) (n int, err error) {
	return fmt.Fscanf(r, format, a...)
}

func Fscanln(r io.Reader, a ...interface{}) (n int, err error) {
	return fmt.Fscanln(r, a...)
}

func Print(a ...interface{}) (n int, err error) {
	return fmt.Print(a...)
}

func Printf(format string, a ...interface{}) (n int, err error) {
	return fmt.Printf(format, a...)
}

func Println(a ...interface{}) (n int, err error) {
	return fmt.Println(a...)
}

func Scan(a ...interface{}) (n int, err error) {
	return fmt.Scan(a...)
}

func Scanf(format string, a ...interface{}) (n int, err error) {
	return fmt.Scanf(format, a...)
}

func Scanln(a ...interface{}) (n int, err error) {
	return fmt.Scanln(a...)
}

func Sprint(a ...interface{}) string {
	return fmt.Sprint(a...)
}

func Sprintf(format string, a ...interface{}) string {
	return fmt.Sprintf(format, a...)
}

func Sprintln(a ...interface{}) string {
	return fmt.Sprintln(a...)
}

func Sscan(str string, a ...interface{}) (n int, err error) {
	return fmt.Sscan(str, a...)
}

func Sscanf(str string, format string, a ...interface{}) (n int, err error) {
	return fmt.Sscanf(str, format, a...)
}

func Sscanln(str string, a ...interface{}) (n int, err error) {
	return fmt.Sscanln(str, a...)
}

type Formatter fmt.Formatter
type GoStringer fmt.GoStringer
type ScanState fmt.ScanState
type Scanner fmt.Scanner
type State fmt.State
type Stringer fmt.Stringer
