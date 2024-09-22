package main

import "fmt"

const Reset = "\033[0m"
const Red = "\033[31m"
const Green = "\033[32m"
const Yellow = "\033[33m"
const Blue = "\033[34m"
const Magenta = "\033[35m"
const Cyan = "\033[36m"
const White = "\033[97m"

func main() {
	contentStr := "this is a line\nthis is the second line\nfinal line"
	r, g, b := 0, 0, 255
	for _, char := range contentStr {
		fmt.Printf("\033[38;2;%d;%d;%dm%c", r, g, b, char)
	}
}
