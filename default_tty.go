package tty

func Clean() {
	DefaultTty.Clean()
}
func Close() error {
	return DefaultTty.Close()
}
func Dimensions() (height, width int, err error) {
	return DefaultTty.Dimensions()
}
func EraseLine() {
	DefaultTty.EraseLine()
}
func Flush() (n int, e error) {
	return DefaultTty.Flush()
}
func MoveTo(line, pos int) {
	DefaultTty.MoveTo(line, pos)
}
func Open() (err error) {
	return DefaultTty.Open()
}
func Read(data []byte) (n int, e error) {
	return DefaultTty.Read(data)
}
func SetBgColor(r, g, b int) {
	DefaultTty.SetBgColor(r, g, b)
}
func SetBlink() {
	DefaultTty.SetBlink()
}
func SetBold() {
	DefaultTty.SetBold()
}
func SetDefaultBgColor() {
	DefaultTty.SetDefaultBgColor()
}
func SetDefaultColor() {
	DefaultTty.SetDefaultColor()
}
func SetDefaultFgColor() {
	DefaultTty.SetDefaultFgColor()
}
func SetFgColor(r, g, b int) {
	DefaultTty.SetFgColor(r, g, b)
}
func SetNoBlink() {
	DefaultTty.SetNoBlink()
}
func SetNormal() {
	DefaultTty.SetNormal()
}
func SetReverse() {
	DefaultTty.SetReverse()
}
func SetUnderscore() {
	DefaultTty.SetUnderscore()
}
func Write(data []byte) (n int, e error) {
	return DefaultTty.Write(data)
}
