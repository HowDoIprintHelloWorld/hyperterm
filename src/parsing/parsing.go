package main

import (
	"fmt"
	"strings"
)

const (
	errNotFound = iota
)

type Command struct {
	BaseCommand string
	Args        []string
}

type ParserError struct {
	Code    int
	Message string
}

func (pe *ParserError) Error() string {
	return fmt.Sprintf("Error %d: %s", pe.Code, pe.Message)
}

type Parser interface {
	Parse(input string) ([]*Command, error)
}

type parser struct {
}

func NewParser() Parser {
	return &parser{}
}

func (p *parser) Parse(input string) ([]*Command, error) {
	var commands = make([]*Command, 0)

	parts := strings.Split(input, "|")
	for _, part := range parts {
		cmd, err := p.parseCommand(part)
		if err != nil {
			return nil, err
		}
		commands = append(commands, cmd)
	}

	return commands, nil
}

func (p *parser) parseCommand(input string) (*Command, error) {
	var cmd = &Command{}
	var currentArg strings.Builder
	var insideQuote bool

	for i := 0; i < len(input); i++ {
		char := input[i]

		switch char {
		case '"':
			insideQuote = !insideQuote
		case ' ':
			if insideQuote {
				currentArg.WriteByte(char)
			} else {
				if currentArg.Len() > 0 {
					if cmd.BaseCommand == "" {
						cmd.BaseCommand = currentArg.String()
					} else {
						cmd.Args = append(cmd.Args, currentArg.String())
					}
					currentArg.Reset()
				}
			}
		default:
			currentArg.WriteByte(char)
		}
	}

	if currentArg.Len() > 0 {
		if cmd.BaseCommand == "" {
			cmd.BaseCommand = currentArg.String()
		} else {
			cmd.Args = append(cmd.Args, currentArg.String())
		}
	}

	if cmd.BaseCommand == "" {
		return cmd, &ParserError{
			Code:    errNotFound,
			Message: "No Base Command Found",
		}
	}

	return cmd, nil
}
