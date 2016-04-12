package tty

// Clean cleans all the screen
func Clean() {
	DefaultTty.Clean()
}

// Close will restore the terminal to the initial state
func Close() error {
	return DefaultTty.Close()
}

// Dimensions returns terminal dimensions
func Dimensions() (height, width int, err error) {
	return DefaultTty.Dimensions()
}

// EraseLine removes all the text in the line
func EraseLine() {
	DefaultTty.EraseLine()
}

// Flush takes all the previous writes and flush them to Stdout
func Flush() (n int, e error) {
	return DefaultTty.Flush()
}

// MoveTo makes the next text to be writen in a specific position
func MoveTo(line, pos int) {
	DefaultTty.MoveTo(line, pos)
}

// Open will put the terminal in Raw mode and clean the scree
func Open() (err error) {
	return DefaultTty.Open()
}

// Read reads data from for Stdout
// If InputDelay is different than 0, a buffer will be use and a call to Read will
// not imply a call to Stdin.Read
func Read(data []byte) (n int, e error) {
	return DefaultTty.Read(data)
}

// SetBgColor set the background color for the next text to be writen
func SetBgColor(r, g, b int) {
	DefaultTty.SetBgColor(r, g, b)
}

// SetBlink makes next text to be writen to blink
func SetBlink() {
	DefaultTty.SetBlink()
}

// SetBold makes the next text to be writen to be bold
func SetBold() {
	DefaultTty.SetBold()
}

// SetDefaultBgColor makes the next text to be writen to use the default Background color
func SetDefaultBgColor() {
	DefaultTty.SetDefaultBgColor()
}

// SetDefaultColor is the same as calling SetDefaultFgColor and SetDefaultBgColor
func SetDefaultColor() {
	DefaultTty.SetDefaultColor()
}

// SetDefaultFgColor makes the next text to be writen to use the default Foreground color
func SetDefaultFgColor() {
	DefaultTty.SetDefaultFgColor()
}

// SetFgColor set the foreground color for the next text to be writen
func SetFgColor(r, g, b int) {
	DefaultTty.SetFgColor(r, g, b)
}

// SetNoBlink makes next text to be writen not to blink
func SetNoBlink() {
	DefaultTty.SetNoBlink()
}

// SetNormal disables Bold Underlink Reverse and Blink for the next text to be writen
func SetNormal() {
	DefaultTty.SetNormal()
}

// SetReverse makes the next text to be writen to be with reversed colors
func SetReverse() {
	DefaultTty.SetReverse()
}

// SetUnderscore makes the next text to be writen to have underscore
func SetUnderscore() {
	DefaultTty.SetUnderscore()
}

// Writes data to the Stdout
// If Fps is not 0 it will buffer the data
func Write(data []byte) (n int, e error) {
	return DefaultTty.Write(data)
}
