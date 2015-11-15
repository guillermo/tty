package tty

import (
	"testing"
	"time"
)

func TestTty(t *testing.T) {

	DefaultTty.ManualFlush = true
	DefaultTty.Open()
	// It will not open inside vim and return an error
	defer DefaultTty.Close()

	height, width, _ := DefaultTty.Dimensions()
	text := "The Editor"

	x := (width - len(text)) / 2
	y := height / 2

	spaces := make([]byte, len(text)+4)
	for i := range spaces {
		spaces[i] = []byte(" ")[0]
	}

	DefaultTty.SetBgColor(20, 133, 204)
	DefaultTty.SetFgColor(178, 18, 18)

	// First
	DefaultTty.MoveTo(y-1, x-2)
	DefaultTty.Write(spaces)
	DefaultTty.MoveTo(y, x-2)
	DefaultTty.Write([]byte("  " + text + "  "))
	DefaultTty.MoveTo(y+1, x-2)
	DefaultTty.Write(spaces)

	DefaultTty.MoveTo(y+2, x+1)
	DefaultTty.SetDefaultColor()
	DefaultTty.SetNormal()
	DefaultTty.SetBlink()
	DefaultTty.EraseLine()
	DefaultTty.Write([]byte("Normal"))
	DefaultTty.SetNoBlink()
	DefaultTty.Flush()
	time.Sleep(time.Second * 2)

	// Bold
	DefaultTty.SetBgColor(20, 133, 204)
	DefaultTty.SetFgColor(178, 18, 18)
	DefaultTty.SetBold()
	DefaultTty.MoveTo(y, x-2)
	DefaultTty.SetBlink()
	DefaultTty.Write([]byte("  " + text + "  "))

	DefaultTty.SetNormal()
	DefaultTty.MoveTo(y+2, x+1)
	DefaultTty.SetDefaultColor()
	DefaultTty.EraseLine()
	DefaultTty.Write([]byte("Bold"))
	DefaultTty.SetNoBlink()
	DefaultTty.Flush()
	time.Sleep(time.Second * 2)

	// Reverse
	DefaultTty.SetBgColor(20, 133, 204)
	DefaultTty.SetFgColor(178, 18, 18)
	DefaultTty.SetReverse()
	DefaultTty.MoveTo(y, x-2)
	DefaultTty.SetBlink()
	DefaultTty.Write([]byte("  " + text + "  "))

	DefaultTty.SetNormal()
	DefaultTty.MoveTo(y+2, x+1)
	DefaultTty.SetDefaultColor()
	DefaultTty.EraseLine()
	DefaultTty.Write([]byte("Reverse"))
	DefaultTty.SetNoBlink()
	DefaultTty.Flush()
	time.Sleep(time.Second * 2)

	// Underscore
	DefaultTty.SetBgColor(20, 133, 204)
	DefaultTty.SetFgColor(178, 18, 18)
	DefaultTty.SetUnderscore()
	DefaultTty.MoveTo(y, x-2)
	DefaultTty.SetBlink()
	DefaultTty.Write([]byte("  " + text + "  "))

	DefaultTty.SetNormal()
	DefaultTty.MoveTo(y+2, x+1)
	DefaultTty.SetDefaultColor()
	DefaultTty.EraseLine()
	DefaultTty.Write([]byte("Underscore"))
	DefaultTty.SetNoBlink()
	DefaultTty.Flush()
	time.Sleep(time.Second * 2)

}
