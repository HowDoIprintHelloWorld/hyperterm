package termio

import (
	"log"

	nc "github.com/rthornton128/goncurses"
)

type Key_command int

const (
	No_command Key_command = iota
	Quit
)

type Screen struct {
	stdscr *nc.Window
	Prompt *Prompt
}

type Line struct {
	content         string
	cursor_position int
	max_len         int
}

func (s *Screen) End() {
	nc.End()
}

func (s *Screen) create_line() *Line {
	_, max_x := s.stdscr.MaxYX()
	return &Line{
		"",
		0,
		max_x,
	}
}

func New_screen() *Screen {
	stdscr, err := nc.Init()
	if err != nil {
		log.Fatal("[E] Error getting screen: ", err)
	}
	nc.Raw(true)
	// nc.CBreak(true)
	nc.Cursor(1)
	nc.Echo(false)
	nc.SetEscDelay(0)
	stdscr.Keypad(true)
	stdscr.ScrollOk(true)
	// NOTE: replace with deco_template from settings
	prompt := Make_prompt("%red[$t:H$:$t:M$:$t:S$]  ")

	screen := Screen{
		stdscr,
		&prompt,
	}
	return &screen
}

func (l *Line) insert_character(position int, character string) {
	l.content = l.content[:position] + character + l.content[position:]
}

func (l *Line) delete_character(position int) {
	l.content = l.content[:position-1] + l.content[position:]
}

func (s *Screen) print_prompt_content(y int, x int, initial_x int, line *Line) {
	s.stdscr.Move(y, initial_x)
	s.stdscr.ClearToEOL()
	s.stdscr.Print(line.content)

	// This is for testing only. It shows the current line content
	s.stdscr.Move(3, 2)
	s.stdscr.ClearToEOL()
	s.stdscr.Print(line.content)
	s.stdscr.Move(y, x+1)
}

func (s *Screen) Print_prompt() {
	s.Prompt.make_prompt_deco()
	prompt_deco := s.Prompt.prompt_deco
	y, _ := s.stdscr.CursorYX()
	s.stdscr.Move(y, 0)
	s.stdscr.Print(prompt_deco)
}

func (s *Screen) Get_line() (string, int) {
	line := s.create_line()
	// y, _ := s.stdscr.CursorYX()
	// s.stdscr.Move(y, 0)
	_, initial_x := s.stdscr.CursorYX()
	active := true
	var char nc.Key
	key_command := No_command
	nc.FlushInput()

	for active {
		s.stdscr.Refresh()
		char = s.stdscr.GetChar()
		y, x := s.stdscr.CursorYX()
		max_y, max_x := s.stdscr.MaxYX()
		switch char {
		case 10: // Enter
			active = false
			if y >= max_y-1 {
				s.stdscr.Scroll(1)
			}
			s.stdscr.Move(y+1, 0)
			s.stdscr.ClearToEOL()
		case 4: // CTRL+D
			active = false
			key_command = Quit
		case nc.KEY_BACKSPACE:
			if x-initial_x <= 0 {
				break
			}
			s.stdscr.Move(y, x-1)
			s.stdscr.DelChar()
			line.delete_character(x - initial_x)
			s.print_prompt_content(y, x-2, initial_x, line)
		case nc.KEY_LEFT:
			if x <= initial_x {
				break
			}
			s.stdscr.Move(y, x-1)
		case nc.KEY_RIGHT:
			if x-initial_x > len(line.content)-1 {
				break
			}
			s.stdscr.Move(y, x+1)
		default:
			if char < 33 || char > 126 {
				break
			} else if x+2 >= max_x {
				break
			}
			line.insert_character(x-initial_x, nc.KeyString(char))
			s.print_prompt_content(y, x, initial_x, line)
		}
	}
	return line.content, int(key_command)
}
