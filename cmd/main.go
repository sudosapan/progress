package main

import (
	"time"

	"github.com/sudosapan/progress/spinner"
)

func main() {

	spin := &spinner.Spinner{
		Message:    "Showing example of spinner",
		Delay:      200,
		Characters: []rune{'-', '\\', '|', '/'},
	}
	spin.Start()
	time.AfterFunc(4*time.Second, spin.Stop)
	time.Sleep(5 * time.Second)

	spin = &spinner.Spinner{}
	spin.Start()
	time.AfterFunc(4*time.Second, spin.Stop)
	time.Sleep(5 * time.Second)
}
