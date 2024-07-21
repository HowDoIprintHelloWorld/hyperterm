package main

import (
	"fmt"
	termio "hyperterm/src/termio"
)

func main() {
	prompt := termio.Make_prompt()
	for {
		line, _ := prompt.Get_line()
		fmt.Println(line)
	}
}
