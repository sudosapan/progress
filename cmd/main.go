package main

import (
	"time"

	"github.com/sudosapan/progress"
)

func main() {

	spin := &progress.Spinner{
		Message:    "Showing example of spinner",
		Delay:      200,
		Characters: []rune{'-', '\\', '|', '/'},
	}
	spin.Start()
	time.AfterFunc(4*time.Second, spin.Stop)
	time.Sleep(5 * time.Second)

	spin = &progress.Spinner{}
	spin.Start()
	time.AfterFunc(4*time.Second, spin.Stop)
	time.Sleep(5 * time.Second)
}
