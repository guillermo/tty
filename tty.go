// Package tty provides primitives to communicate with the terminal
//
//   tty.Open()
//   defer tty.Close()
//
//   rows, cols, err := tty.Dimensions()
//
//   tty.MoveTo(2, 1)
//   tty.SetDefaultColor()
//   tty.SetBgColor(20, 133, 204)
//   tty.SetFgColor(178, 18, 18)
//   tty.SetBold()
//   tty.SetUnderscore()
//   tty.SetReverse()
//   tty.SetBlink()
//   tty.Write([]byte("Bold"))
//
package tty

import (
	"bytes"
	"os"
	"time"

	"github.com/kless/term"

	"github.com/guillermo/dreader"
)

func init() {
	DefaultTty = &Tty{
		InputDelay: time.Millisecond * 10,
		Stdin:      os.Stdin,
		Stdout:     os.Stdout,
		Fps:        60,
	}
}

// DefaultTty is automatically associated with the current Stdin and Stdout
var DefaultTty *Tty

// Tty represent a terminal
type Tty struct {
	// InputDelay is the amount of time that Read should wait to complete the read.
	// While sending escape sequences is not the same \027 that is the Esc key than
	// \027 ... that is ...
	// If InputDelay is 0, it will return immediately
	InputDelay time.Duration

	// Stdin is the file where the Read function will read from.
	// It defaults to os.Stdin
	Stdin *os.File

	// Stdout is the file where the Write function will write to.
	// It defaults to os.Stdout
	Stdout *os.File

	// Fps is the maximum number of times that the write function will be called in
	// Stdout during a second
	// If Fps is 0 Writes are sync
	Fps int

	// ManualFlush prevents Write to behave synchronously
	// If ManualFlush is set to true you should call Flush for your writes to be dump
	ManualFlush bool

	wbuf   *bytes.Buffer
	term   *term.Terminal
	r      *dreader.DelayedReader
	scroll int
}

// Read reads data from for Stdout
// If InputDelay is different than 0, a buffer will be use and a call to Read will
// not imply a call to Stdin.Read
func (t *Tty) Read(data []byte) (n int, e error) {
	if t.r == nil {
		t.r = dreader.New(t.Stdin, t.InputDelay)
	}
	return t.r.Read(data)
}

// Writes data to the Stdout
// If Fps is not 0 it will buffer the data
func (t *Tty) Write(data []byte) (n int, e error) {
	if t.ManualFlush {
		if t.wbuf == nil {
			t.wbuf = new(bytes.Buffer)
		}
		return t.wbuf.Write(data)
	}
	return t.Stdout.Write(data)
}

// Flush takes all the previous writes and flush them to Stdout
func (t *Tty) Flush() (n int, e error) {
	if t.wbuf == nil || t.wbuf.Len() == 0 {
		return 0, nil
	}

	m, err := t.wbuf.WriteTo(t.Stdout)
	return int(m), err
}

// Open will put the terminal in Raw mode and clean the scree
func (t *Tty) Open() (err error) {
	t.saveScreen()
	t.disableCursor()
	if t.ManualFlush {
		t.Flush()
	}
	if t.term == nil {
		t.term, err = term.New()
		if err != nil {
			return err
		}
	}

	err = t.term.RawMode()
	if err != nil {
		t.restoreScreen()
		return err
	}
	return nil
}

// Close will restore the terminal to the initial state
func (t *Tty) Close() error {
	if t.term != nil {
		t.term.Restore()
	}

	t.restoreScreen()
	t.enableCursor()
	if t.ManualFlush {
		t.Flush()
	}
	return nil
}

func (t *Tty) raw() (err error) {
	if t.term, err = term.New(); err != nil {
		return err
	}
	if err := t.term.RawMode(); err != nil {
		return err
	}
	return nil
}

// Dimensions returns terminal dimensions
func (t *Tty) Dimensions() (height, width int, err error) {
	if t.term == nil {
		t.term, err = term.New()
		if err != nil {
			return 25, 80, err
		}
	}
	height, width, err = t.term.GetSize()
	if height == 0 {
		height = 25
	}
	if width == 0 {
		width = 25
	}
	return height, width, err
}
