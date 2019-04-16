package main

import (
	fmt "github.com/go-cli-tools/format"
)

//noinspection ALL
func main() {
	fmt.Println("Normal")
	fmt.Println("")
	fmt.WithStyle(fmt.Bold).Println("Bold")
	fmt.Println("")
	fmt.WithStyle(fmt.Italics).Println("Italics")
	fmt.Println("")
	fmt.WithStyle(fmt.Light).Println("Light")
	fmt.Println("")
	fmt.WithStyle(fmt.Underlined).Println("Underlined")
	fmt.Println("")
	fmt.WithStyle(fmt.Blink).Println("Blink")
	fmt.Println("")
	fmt.WithStyle(fmt.Inverted).Println("Inverted")
	fmt.Println("")
	fmt.WithStyle(fmt.Hidden).Println("Hidden")
	fmt.Println("")
	fmt.WithStyle(fmt.StrikeThrough).Println("StrikeThrough")
	fmt.Println("")
}
