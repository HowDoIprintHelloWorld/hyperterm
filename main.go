package main

import (
	"fmt"
	termio "hyperterm/src/io"
)

func main() {
	prompt := termio.Make_prompt()
	for {
		line, _ := prompt.Get_line()
		fmt.Println(line)
	}
}
