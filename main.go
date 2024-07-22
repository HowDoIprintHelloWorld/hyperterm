package main

import (
	termio "hyperterm/src/termio"
)

func main() {
	screen := termio.New_screen()
	defer screen.End()

	for {
		screen.Print_prompt()
		_, key_command := screen.Get_line()
		if key_command == int(termio.Quit) {
			break
		}
	}
	// line, _ := screen.Prompt.Get_line()
	// fmt.Println(line)

}
