package main

import (
	fmt "github.com/go-cli-tools/format"
	"github.com/go-cli-tools/format/colour"
	"github.com/go-cli-tools/format/style"
)

func main() {
	_, _ = fmt.Println("STYLES")
	_, _ = fmt.Println("")
	_, _ = fmt.Printf("%4s", "Normal")
	_, _ = fmt.WithStyle(style.Bold).Printf("%-6s", "Bold")
	_, _ = fmt.WithStyle(style.Italics).Printf("%-6s", "Italics")
	_, _ = fmt.WithStyle(style.Light).Printf("%-6s", "Light")
	_, _ = fmt.WithStyle(style.Underlined).Printf("%-6s", "Underlined")
	_, _ = fmt.WithStyle(style.Blink).Printf("%-6s", "Blink")
	_, _ = fmt.WithStyle(style.Inverted).Printf("%-6s", "Inverted")
	_, _ = fmt.WithStyle(style.Hidden).Printf("%-6s", "Hidden")
	_, _ = fmt.WithStyle(style.StrikeThrough).Printf("%-6s", "StrikeThrough")
	_, _ = fmt.Println("")
	_, _ = fmt.Println("")

	_, _ = fmt.Printf("%-32s%-32s\n", "STANDARD COLOURS", "BRIGHT COLOURS")
	_, _ = fmt.Println("")
	_, _ = fmt.WithColour(colour.Black).Printf("%-6s", "Black")
	_, _ = fmt.WithColour(colour.Red).Printf("%-6s", "Red")
	_, _ = fmt.WithColour(colour.Green).Printf("%-6s", "Green")
	_, _ = fmt.WithColour(colour.Yellow).Printf("%-6s", "Yellow")
	_, _ = fmt.WithColour(colour.Blue).Printf("%-6s", "Blue")
	_, _ = fmt.WithColour(colour.Magenta).Printf("%-6s", "Magenta")
	_, _ = fmt.WithColour(colour.Cyan).Printf("%-6s", "Cyan")
	_, _ = fmt.WithColour(colour.White).Printf("%-6s", "White")
	_, _ = fmt.WithColour(colour.BrightBlack).Printf("%-6s", "Black")
	_, _ = fmt.WithColour(colour.BrightRed).Printf("%-6s", "Red")
	_, _ = fmt.WithColour(colour.BrightGreen).Printf("%-6s", "Green")
	_, _ = fmt.WithColour(colour.BrightYellow).Printf("%-6s", "Yellow")
	_, _ = fmt.WithColour(colour.BrightBlue).Printf("%-6s", "Blue")
	_, _ = fmt.WithColour(colour.BrightMagenta).Printf("%-6s", "Magenta")
	_, _ = fmt.WithColour(colour.BrightCyan).Printf("%-6s", "Cyan")
	_, _ = fmt.WithColour(colour.BrightWhite).Printf("%-6s", "White")
	_, _ = fmt.Println("")
	_, _ = fmt.Println("")

	_, _ = fmt.Printf("%-64s%-64s\n", "BACKGROUND COLOURS", "BRIGHT BACKGROUND COLOURS")
	_, _ = fmt.Println("")
	_, _ = fmt.WithBackground(colour.Black).Printf("%-6s", "Black")
	_, _ = fmt.WithBackground(colour.Red).Printf("%-6s", "Red")
	_, _ = fmt.WithBackground(colour.Green).Printf("%-6s", "Green")
	_, _ = fmt.WithBackground(colour.Yellow).Printf("%-6s", "Yellow")
	_, _ = fmt.WithBackground(colour.Blue).Printf("%-6s", "Blue")
	_, _ = fmt.WithBackground(colour.Magenta).Printf("%-6s", "Magenta")
	_, _ = fmt.WithBackground(colour.Cyan).Printf("%-6s", "Cyan")
	_, _ = fmt.WithBackground(colour.White).WithColour(colour.Black).Printf("%-6s", "White")
	_, _ = fmt.WithBackground(colour.BrightBlack).Printf("%-6s", "Black")
	_, _ = fmt.WithBackground(colour.BrightRed).Printf("%-6s", "Red")
	_, _ = fmt.WithBackground(colour.BrightGreen).Printf("%-6s", "Green")
	_, _ = fmt.WithBackground(colour.BrightYellow).Printf("%-6s", "Yellow")
	_, _ = fmt.WithBackground(colour.BrightBlue).Printf("%-6s", "Blue")
	_, _ = fmt.WithBackground(colour.BrightMagenta).Printf("%-6s", "Magenta")
	_, _ = fmt.WithBackground(colour.BrightCyan).Printf("%-6s", "Cyan")
	_, _ = fmt.WithBackground(colour.BrightWhite).WithColour(colour.Black).Printf("%-6s", "White")
	_, _ = fmt.Println("")
	_, _ = fmt.Println("")

	_, _ = fmt.Println("256 COLOURS MODE - FOREGROUND")
	for i := 0; i < 256; i++ {
		if i%16 == 0 {
			_, _ = fmt.Println("")
		}
		_, _ = fmt.WithColour256(i).Printf("  %-4d", i)
	}
	_, _ = fmt.Println("")
	_, _ = fmt.Println("")

	_, _ = fmt.Println("256 COLOURS MODE - BACKGROUND")
	for i := 0; i < 256; i++ {
		if i%16 == 0 {
			_, _ = fmt.Println("")
		}
		if i < 9 {
			_, _ = fmt.WithBackground256(i).Printf("  %-4d", i)
		} else if i < 99 {
			if i == 15 || (i >= 46 && i <= 51) || (i >= 82 && i <= 87) {
				_, _ = fmt.WithBackground256(i).WithColour(colour.Black).Printf("  %-4d", i)
			} else {
				_, _ = fmt.WithBackground256(i).Printf("  %-4d", i)
			}
		} else {
			if (i >= 118 && i <= 123) || (i >= 154 && i <= 159) || (i >= 190 && i <= 195) || (i >= 224 && i <= 231) || i >= 240 {
				_, _ = fmt.WithBackground256(i).WithColour(colour.Black).Printf("  %-4d", i)
			} else {
				_, _ = fmt.WithBackground256(i).Printf("  %-4d", i)
			}
		}
	}
	_, _ = fmt.Println("")
	_, _ = fmt.Println("")
}
