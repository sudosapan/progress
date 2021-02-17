package progress

import (
	"time"
)

//example for Progress method of Spinnner
func ExampleSpinner_Start() {
	spinner := &Spinner{}
	spinner.Start()
	time.AfterFunc(10*time.Second, spinner.Stop)
}

func ExampleSpinner() {
	spinner := &Spinner{
		Message:    "Searching user details",
		Delay:      200,
		Characters: []rune{'-', '\\', '|', '/'},
	}

	spinner.Start()
	time.AfterFunc(10*time.Second, spinner.Stop)
}
