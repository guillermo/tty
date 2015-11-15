package tty

import (
	"fmt"
	"strconv"
	"strings"
)

func (t *Tty) writeSequence(fargs ...interface{}) {
	format := fargs[0].(string)
	format = "\x1b[" + format
	fmt.Fprintf(t, format, fargs[1:]...)
}

func (t *Tty) writeRawSequence(fargs ...interface{}) {
	format := fargs[0].(string)
	format = "\x1b[" + format
	fmt.Fprintf(t.Stdout, format, fargs[1:]...)
}

func (t *Tty) disableCursor() {
	t.writeRawSequence("?25l")
}

func (t *Tty) enableCursor() {
	t.writeRawSequence("?25h")
}

// SetBlink makes next text to be writen to blink
func (t *Tty) SetBlink() {
	t.writeSequence("5m")
	t.writeSequence("6m")
}

// SetNoBlink makes next text to be writen not to blink
func (t *Tty) SetNoBlink() {
	t.writeSequence("25m")
}

// SetBold makes the next text to be writen to be bold
func (t *Tty) SetBold() {
	t.writeSequence("1m")
}

// SetReverse makes the next text to be writen to be with reversed colors
func (t *Tty) SetReverse() {
	t.writeSequence("7m")
}

// SetUnderscore makes the next text to be writen to have underscore
func (t *Tty) SetUnderscore() {
	t.writeSequence("4m")
}

// EraseLine removes all the text in the line
func (t *Tty) EraseLine() {
	t.writeSequence("K")
}

// Clean cleans all the screen
func (t *Tty) Clean() {
	t.writeSequence("2J")
}

// SetNormal disables Bold Underlink Reverse and Blink for the next text to be writen
func (t *Tty) SetNormal() {
	t.writeSequence("0m")
}

// SetDefaultBgColor makes the next text to be writen to use the default Background color
func (t *Tty) SetDefaultBgColor() {
	t.writeSequence("49m")
}

// SetDefaultFgColor makes the next text to be writen to use the default Foreground color
func (t *Tty) SetDefaultFgColor() {
	t.writeSequence("39m")
}

// SetDefaultColor is the same as calling SetDefaultFgColor and SetDefaultBgColor
func (t *Tty) SetDefaultColor() {
	t.SetDefaultBgColor()
	t.SetDefaultFgColor()
}

// SetFgColor set the foreground color for the next text to be writen
func (t *Tty) SetFgColor(r, g, b int) {
	t.writeSequence("38;2;%d;%d;%dm", r, g, b)
}

// SetBgColor set the background color for the next text to be writen
func (t *Tty) SetBgColor(r, g, b int) {
	t.writeSequence("48;2;%d;%d;%dm", r, g, b)
}

func (t *Tty) getCursorPosition() (line, x int, err error) {
	t.writeSequence("6n")
	buf := make([]byte, 20)
	n, err := t.Read(buf)
	if err != nil {
		return 0, 0, err
	}

	sequence := string(buf[:n])

	// Remove escape
	pos := strings.Index(sequence, "\x1b[")
	if pos == -1 || sequence[pos+1] != []byte("[")[0] {
		return 0, 0, nil
	}
	sequence = sequence[pos+2:]

	// Remove R
	pos = strings.Index(sequence, "R")
	if pos == -1 {
		return 0, 0, nil
	}
	sequence = sequence[:pos]

	// Split by ;
	data := strings.Split(sequence, ";")
	if len(data) != 2 {
		return 0, 0, nil
	}

	// Convert
	line, err = strconv.Atoi(data[0])
	if err != nil {
		return 0, 0, nil
	}
	x, err = strconv.Atoi(data[1])
	if err != nil {
		return 0, 0, nil
	}

	return line, x, nil
}

func (t *Tty) clear() {
	t.writeSequence("J")
}

func (t *Tty) scrollUp(n int) {
	t.writeSequence(strconv.Itoa(n) + "S")
}
func (t *Tty) scrollDown(n int) {
	t.writeSequence(strconv.Itoa(n) + "T")
}

func (t *Tty) saveScreen() {
	t.Stdout.Write([]byte("\x1b[?47h"))
}
func (t *Tty) restoreScreen() {
	t.Stdout.Write([]byte("\x1b[?47l"))
}

// MoveTo makes the next text to be writen in a specific position
func (t *Tty) MoveTo(line, pos int) {
	t.writeSequence("%d;%dH", line, pos)
}
