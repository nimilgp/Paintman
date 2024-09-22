# Paintman
An implementation of lolcat in Go. 

## Compilation
1. move in to the directory with main.go file
2. go build -o paintman main.go

## modules used
- bufio
- flag
- math
- os
- fmt

## Ussage
Takes in text from stdin (supports piping) and colors it and outputs in stdout.
eg: echo hello, terminal emulator!! are you in color??? | ./paintman -rainbow
