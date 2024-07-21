package termio

import (
	"bufio"
	"fmt"
	"os"
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

	reader *bufio.Reader
	err    error
}

func (p *Prompt) Get_line() (string, error) {
	p.make_prompt_deco()
	p.print_prompt_deco()
	line, err := p.read_line()
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
	escaped := false
	variable := false
	p.prompt_deco_template = []string{}
	current_part := ""
	for _, char := range p.prompt_deco_string_template {

		if char == '\\' && !escaped {
			escaped = true
			continue
		} else if escaped && char == '$' {
			current_part += string(char)
		} else if char == '$' {
			if variable {
				p.prompt_deco_template = append(p.prompt_deco_template, current_part+"$")
				current_part = ""
			} else {
				p.prompt_deco_template = append(p.prompt_deco_template, current_part)
				current_part = "$"
			}
			variable = !variable
		} else {
			current_part += string(char)
		}
		escaped = false
	}
	p.prompt_deco_template = append(p.prompt_deco_template, current_part)
}

func (p *Prompt) print_prompt_deco() {
	fmt.Print(p.prompt_deco)
}

func Make_prompt() Prompt {
	prompt := Prompt{
		"",

		"",
		// "[$t:H$:$t:M$:$t:S$]",
		"[$t:H$:$t:M$:$t:S$]",
		[]string{},
		map[string]string{},

		bufio.NewReader(os.Stdin),
		nil,
	}
	prompt.parse_prompt_deco_string_template()
	return prompt
}

func (p *Prompt) read_line() (string, error) {
	line, err := p.reader.ReadString('\n')
	return line[:len(line)-1], err
}
