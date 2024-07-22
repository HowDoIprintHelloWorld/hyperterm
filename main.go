package main

import (
	"fmt"
	parsing "hyperterm/src/parsing"
	termio "hyperterm/src/termio"
)

func main() {
	prompt := termio.Make_prompt("%red[$t:H$:$t:M$:$t:S$]  ")
	parser := parsing.NewParser()
	for {
		line, _ := prompt.Get_line()
		commands, _ := parser.Parse(line)
		for _, command := range commands {
			fmt.Print(command, " | ")
		}
		fmt.Println()
	}
}
