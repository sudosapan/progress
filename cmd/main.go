package main

import (
	"time"

	"github.com/sudosapan/progress/spinner"
)

func main() {

	spin := &spinner.Spinner{
		Message:    "Showing example of spinner",
		Delay:      200,
		Characters: []string{"-", "\\", "|", "/"},
	}

	stopFunc := spin.Start()
	time.AfterFunc(2*time.Second, stopFunc)
	time.Sleep(5 * time.Second)
}
