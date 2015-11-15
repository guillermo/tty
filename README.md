# tty
Go package that abstracts a pty and provides common functions to deal with the terminal

```
PACKAGE DOCUMENTATION

package tty
    import "."

    Package tty provides primitives to communicate with the terminal

	  DefaultTty.Open()
	  defer DefaultTty.Close()

	  DefaultTty.MoveTo(y+2, x+1)
	  DefaultTty.SetDefaultColor()
	  DefaultTty.EraseLine()
	  DefaultTty.SetBgColor(20, 133, 204)
	  DefaultTty.SetFgColor(178, 18, 18)
	  DefaultTty.SetBold()
	  DefaultTty.SetUnderscore()
	  DefaultTty.SetReverse()
	  DefaultTty.Write([]byte("Bold"))
	  DefaultTty.SetNoBlink()
	  DefaultTty.Flush()

TYPES

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
    // contains filtered or unexported fields
}
    Tty represent a terminal

var DefaultTty *Tty
    DefaultTty is automatically associated with the current Stdin and Stdout

func (t *Tty) Clean()
    Clean cleans all the screen

func (t *Tty) Close() error
    Close will restore the terminal to the initial state

func (t *Tty) Dimensions() (height, width int, err error)
    Dimensions returns terminal dimensions

func (t *Tty) EraseLine()
    EraseLine removes all the text in the line

func (t *Tty) Flush() (n int, e error)
    Flush takes all the previous writes and flush them to Stdout

func (t *Tty) MoveTo(line, pos int)
    MoveTo makes the next text to be writen in a specific position

func (t *Tty) Open() (err error)
    Open will put the terminal in Raw mode and clean the scree

func (t *Tty) Read(data []byte) (n int, e error)
    Read reads data from for Stdout If InputDelay is different than 0, a
    buffer will be use and a call to Read will not imply a call to
    Stdin.Read

func (t *Tty) SetBgColor(r, g, b int)
    SetBgColor set the background color for the next text to be writen

func (t *Tty) SetBlink()
    SetBlink makes next text to be writen to blink

func (t *Tty) SetBold()
    SetBold makes the next text to be writen to be bold

func (t *Tty) SetDefaultBgColor()
    SetDefaultBgColor makes the next text to be writen to use the default
    Background color

func (t *Tty) SetDefaultColor()
    SetDefaultColor is the same as calling SetDefaultFgColor and
    SetDefaultBgColor

func (t *Tty) SetDefaultFgColor()
    SetDefaultFgColor makes the next text to be writen to use the default
    Foreground color

func (t *Tty) SetFgColor(r, g, b int)
    SetFgColor set the foreground color for the next text to be writen

func (t *Tty) SetNoBlink()
    SetNoBlink makes next text to be writen not to blink

func (t *Tty) SetNormal()
    SetNormal disables Bold Underlink Reverse and Blink for the next text to
    be writen

func (t *Tty) SetReverse()
    SetReverse makes the next text to be writen to be with reversed colors

func (t *Tty) SetUnderscore()
    SetUnderscore makes the next text to be writen to have underscore

func (t *Tty) Write(data []byte) (n int, e error)
    Writes data to the Stdout If Fps is not 0 it will buffer the data


```
