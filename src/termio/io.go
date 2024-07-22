package termio

import (
	"bufio"
	"os"
)

type Variable_function func() string

var prompt_deco_variable_functions map[string]Variable_function = map[string]Variable_function{
	"t:S": get_time_seconds,
	"t:M": get_time_minutes,
	"t:H": get_time_hours,
}

type Prompt struct {
	line string

	prompt_deco                 string
	prompt_deco_string_template string
	prompt_deco_template        []string
	prompt_deco_variables       map[string]string

	reader *bufio.Reader
	err    error
}

// Deprecated
/*
func (p *Prompt) Get_line() (string, error) {
	p.make_prompt_deco()
	p.print_prompt_deco()
	line, err := p.read_line()
	p.line = line
	p.err = err
	return line, err
}
*/

func Make_prompt(prompt_deco_string_template string) Prompt {
	prompt := Prompt{
		"",

		"",
		prompt_deco_string_template,
		[]string{},
		map[string]string{},

		bufio.NewReader(os.Stdin),
		nil,
	}
	prompt.parse_prompt_deco_string_template()
	return prompt
}
