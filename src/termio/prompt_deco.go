package termio

import (
	"strings"
)

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

// func (p *Prompt) print_prompt_deco() {
// 	fmt.Print(p.prompt_deco)
// }
