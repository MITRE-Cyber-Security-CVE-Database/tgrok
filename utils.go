package main

import "github.com/fatih/color"

// PrintError prints an error message in red
func PrintError(msg string) {
	red := color.New(color.FgRed)
	red.Println("Error:", msg)
}