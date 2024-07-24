package termio

import (
	"log"

	nc "github.com/rthornton128/goncurses"
)

type Key_command int

const (
	No_command Key_command = iota
	Quit
	Break
)

type Screen struct {
	stdscr *nc.Window
	Prompt *Prompt
}

type Line struct {
	content_string    string
	content_formatted []Line_segment
	cursor_position   int
	max_len           int
}

/*
Lines are set up as follows:
	line := []Line_segment{line_segment1, line_segment2, ..., line_segmentN}
A line consists of multiple line segments, each having one distinguishing feature from the next.
For example, line_segment2 may be displayed bold and red, whereas line_segment1 is colored white.
The function responsible for printing lines can iterate over every line_segment and set its
output-style according to the specifications of the line_segment.
*/

type Line_segment struct {
	text_content string
	bold         bool
	italic       bool

	color int
}

func (l *Line) content_to_string() string {
	content := ""
	for _, segment := range l.content_formatted {
		content += segment.text_content
	}
	return content
}

// IMPORTANT: not yet implemented
func (l *Line) format() {
	l.content_formatted = []Line_segment{
		Line_segment{
			l.content_string,
			false,
			false,
			nc.C_WHITE,
		},
	}
}

func (l *Line) insert_character(position int, character string) {
	content_string := l.content_string
	content_string = content_string[:position] + character + content_string[position:]
	l.content_string = content_string
	l.format()
}

func (l *Line) delete_character(position int) {
	content_string := l.content_string
	content_string = content_string[:position-1] + content_string[position:]
	l.content_string = content_string
	l.format()
}

func (s *Screen) End() {
	nc.End()
}

func (s *Screen) create_line() *Line {
	_, max_x := s.stdscr.MaxYX()
	return &Line{
		"",
		[]Line_segment{},
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
	nc.StartColor()
	nc.UseDefaultColors()
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

func (s *Screen) print_prompt_content(y int, x int, initial_x int, line *Line) {
	s.stdscr.Move(y, initial_x)
	s.stdscr.ClearToEOL()
	for _, segment := range line.content_formatted {
		s.stdscr.Print(segment.text_content)
	}

	// This is for testing only. It shows the current line content
	s.stdscr.Move(3, 2)
	s.stdscr.ClearToEOL()
	s.stdscr.Print(line.content_formatted)
	s.stdscr.Move(y, x+1)
}

func (s *Screen) Print_prompt() {
	s.Prompt.make_prompt_deco()
	prompt_deco := s.Prompt.prompt_deco
	y, _ := s.stdscr.CursorYX()
	s.stdscr.Move(y, 0)
	nc.InitPair(1, nc.C_CYAN, -1)
	s.stdscr.ColorOn(1)
	s.stdscr.Print(prompt_deco)
	s.stdscr.ColorOff(1)
}

func (s *Screen) new_line() {
	y, _ := s.stdscr.CursorYX()
	max_y, _ := s.stdscr.MaxYX()
	if y >= max_y-1 {
		s.stdscr.Scroll(1)
	}
	s.stdscr.Move(y+1, 0)
	s.stdscr.ClearToEOL()
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
		_, max_x := s.stdscr.MaxYX()
		line_length := len(line.content_string)
		switch char {
		case 10: // Enter
			active = false
			s.new_line()
		case 3: // CTRL+C
			active = false
			key_command = Break
			s.new_line()
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
			if x-initial_x > line_length-1 {
				break
			}
			s.stdscr.Move(y, x+1)
		default:
			if char < 32 || char > 126 {
				break
			} else if x+2 >= max_x {
				break
			}
			line.insert_character(x-initial_x, nc.KeyString(char))
			s.print_prompt_content(y, x, initial_x, line)
		}
	}
	return line.content_to_string(), int(key_command)
}
