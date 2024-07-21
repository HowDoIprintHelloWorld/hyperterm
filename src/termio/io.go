package termio

import (
	"fmt"
	"strings"
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

	err error
}

func (p *Prompt) Get_line() (string, error) {
	p.make_prompt_deco()
	p.print_prompt_deco()
	line, err := read_line()
	p.line = line
	p.err = err
	return line, err
}

func (p *Prompt) make_prompt_deco() {
	p.prompt_deco = ""
	for _, part := range p.prompt_deco_template {
		if strings.HasPrefix(part, "$") && strings.HasSuffix(part, "$") {
			part = part[1 : len(part)-1]
			variable_function, ok := prompt_deco_variable_functions[part]
			if ok {
				part = variable_function()
			}
		}
		p.prompt_deco += part
	}
}

func (p *Prompt) parse_prompt_deco_string_template() {

}

func (p *Prompt) print_prompt_deco() {
	fmt.Print(p.prompt_deco)
}

func Make_prompt() Prompt {
	prompt_deco_template := []string{"[", "$t:H$", ":", "$t:M$", ":", "$t:S$", "]"}
	prompt := Prompt{
		"",

		"",
		"[$t:H$:$t:M$:$t:S$]",
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
