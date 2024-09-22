package main

import (
	"flag"
	"fmt"
)

const RESET = "\033[0m"

func main() {
	colorValueMap := map[string]int{
		"black":   30,
		"red":     31,
		"green":   32,
		"yellow":  33,
		"blue":    34,
		"magenta": 35,
		"cyan":    36,
		"white":   97,
	}
	var fgcolor string
	var bgcolor string
	var showColors bool
	flag.StringVar(&fgcolor, "fg", "white", "select foreground color")
	flag.StringVar(&bgcolor, "bg", "cyan", "select background color")
	flag.BoolVar(&showColors, "colors", false, "show availible color options")
	flag.Parse()

	if showColors {
		fmt.Println("red, green, yellow, blue, magenta, cyan, white, black")
		return
	}

	fg, ok := colorValueMap[fgcolor]
	if !ok {
		fmt.Println("foreground color is invalid")
		return
	}
	bg, ok := colorValueMap[bgcolor]
	if !ok {
		fmt.Println("background color is invalid")
		return
	}

	contentStr := "this is a line\nthis is the second line\nfinal line"
	fmt.Printf("\033[%dm", fg)
	fmt.Printf("\033[%dm", bg+10)
	fmt.Print(contentStr)
	fmt.Printf("%s", RESET)
}
