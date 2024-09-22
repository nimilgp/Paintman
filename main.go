package main

import (
	"flag"
	"fmt"
	"math"
	"math/rand"
)

const RESET = "\033[0m"

func rgb(i int) (int, int, int) {
	var f = 0.1
	return int(math.Sin(f*float64(i)+0)*127 + 128),
		int(math.Sin(f*float64(i)+2*math.Pi/3)*127 + 128),
		int(math.Sin(f*float64(i)+4*math.Pi/3)*127 + 128)
}

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
	// var bgcolor string
	var rainbow bool
	var showColors bool
	flag.StringVar(&fgcolor, "color", "system default", "select color")
	// flag.StringVar(&bgcolor, "bg", "cyan", "select background color")
	flag.BoolVar(&showColors, "available", false, "show available color options")
	flag.BoolVar(&rainbow, "rainbow", false, "rainbow text mode")
	flag.Parse()

	if fgcolor != "system default" && rainbow {
		fmt.Println("color and rainbow flags are mutually exclusive")
		return
	}
	if showColors {
		fmt.Println("red, green, yellow, blue, magenta, cyan, white, black")
		return
	}

	fg, ok := colorValueMap[fgcolor]
	if ok {
		fmt.Printf("\033[%dm", fg)
	} else if fgcolor != "system default" {
		fmt.Println("color is invalid")
		return
	}
	// bg, ok := colorValueMap[bgcolor]
	// if !ok {
	// 	fmt.Println("background color is invalid")
	// 	return
	// }

	contentStr := "this is a line\nthis is the second line\nfinal line"
	if rainbow {
		randomseed := rand.Intn(100)
		for _, char := range contentStr {
			r, g, b := rgb(randomseed)
			randomseed++
			fmt.Printf("\033[38;2;%d;%d;%dm%c", r, g, b, char)
		}
	} else {
		// fmt.Printf("\033[%dm", bg+10)
		fmt.Print(contentStr)
		fmt.Printf("%s", RESET)
	}
}
