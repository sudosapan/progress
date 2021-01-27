package spinner

import (
	"time"
)

//Using default spinner
func ExampleDefault() {
	// Executes in go routine
	stopFunc := Default().Start()
	// Perform task for which progress is marked.
	time.AfterFunc(10*time.Second, stopFunc)
}

func ExampleStart() {
	// Executes in go routine
	stopFunc := Start("Another spinner", []string{"-", "\\", "|", "/"}, 200)
	time.AfterFunc(2*time.Second, stopFunc)
	// /Perform task for which progress is marked.
	time.Sleep(3 * time.Second)
}

//example for Progress method of Spinnner
func ExampleSpinner_Start() {
	spinner := &Spinner{
		Message:    "Searching user details",
		Delay:      200,
		Characters: []string{"-", "\\", "|", "/"},
	}

	time.AfterFunc(10*time.Second, spinner.Start())
}
