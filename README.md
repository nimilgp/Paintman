# Paintman
An implementation of lolcat in Go. 

![cat-space](https://github.com/user-attachments/assets/10fb3826-b548-44dd-9218-9e8b9920d2a6)

## Installation

1. **Clone the Repository**

    ```bash
    git clone https://github.com/nimilgp/Paintman
    ```

2. **Navigate to the Directory**

    ```bash
    cd Paintman
    ```
3. **Compile the source**

   ```bash
   go build -o paintman main.go
   ```
4. **Example**
   ```bash
   echo hello, terminal emulator!! are you in color??? | ./paintman -rainbow
   ```

## Modules used
- bufio
- flag
- math
- os
- fmt

## Usage
Takes in text from stdin (supports piping) and colors it and outputs in stdout.

![screenshot](/paintman.png)
