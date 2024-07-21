package src

import (
	"fmt"
)

type Prompt struct {
	line string

	prompt_deco           string
	prompt_deco_template  []string
	prompt_deco_variables map[string]string

	err error
}

func (p Prompt) Get_line() (string, error) {
	p.print_prompt_deco()
	line, err := read_line()
	p.line = line
	p.err = err
	return line, err
}

func (p Prompt) make_prompt_deco() {

}

func (p Prompt) print_prompt_deco() {
	fmt.Print(p.prompt_deco)
}

func Make_prompt() Prompt {
	prompt_deco_template := []string{"[", "$t$", "]"}
	prompt := Prompt{
		"",

		"",
		prompt_deco_template,
		map[string]string{},

		nil,
	}
	return prompt
}

func read_line() (string, error) {
	var line string
	_, err := fmt.Scanln(&line)
	return line, err
}
