package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"math"
	"math/rand"
	"os"
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
	var color string
	var rainbow bool
	var showColors bool
	flag.StringVar(&color, "color", "system default", "select color")
	flag.BoolVar(&showColors, "available", false, "show available color options")
	flag.BoolVar(&rainbow, "rainbow", false, "rainbow text mode")
	flag.Parse()

	if color != "system default" && rainbow {
		fmt.Println("color and rainbow flags are mutually exclusive")
		return
	}
	if showColors {
		fmt.Println("red, green, yellow, blue, magenta, cyan, white, black")
		return
	}

	clr, ok := colorValueMap[color]
	if ok {
		fmt.Printf("\033[%dm", clr)
	} else if color != "system default" {
		fmt.Println("color is invalid")
		return
	}
	var contentStr []rune
	reader := bufio.NewReader(os.Stdin)
	for {
		input, _, err := reader.ReadRune()
		if err == io.EOF && err != nil{
			break
		}
		contentStr = append(contentStr, input)
	}

	if rainbow {
		randomseed := rand.Intn(100)
		for _, char := range contentStr {
			r, g, b := rgb(randomseed)
			randomseed++
			fmt.Printf("\033[38;2;%d;%d;%dm%c", r, g, b, char)
		}
	} else {
		for _, char := range contentStr {
			fmt.Printf("%c", char)
		}
	}
	fmt.Printf("%s\n", RESET)
}
