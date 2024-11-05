# Paintman
An implementation of lolcat in Go. 

![cat-space](https://github.com/user-attachments/assets/10fb3826-b548-44dd-9218-9e8b9920d2a6)

## Compilation
1. move in to the directory with main.go file
2. go build -o paintman main.go

## modules used
- bufio
- flag
- math
- os
- fmt

## Usage
Takes in text from stdin (supports piping) and colors it and outputs in stdout.

eg: echo hello, terminal emulator!! are you in color??? | ./paintman -rainbow

![screenshot](/paintman.png)
