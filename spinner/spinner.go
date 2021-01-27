//Package spinner Package for spinning progress markers. Like in place spinning sticks
package spinner

import (
	"fmt"
	"time"
)

var (
	defaultMessage    string   = "In Progress"
	defaultCharacters []string = []string{"-", "\\", "|", "/"}
	defaultDelay      int      = 200
)

//Default return spinner with message = "In progress" and spinning sticks. Delay is 200ms
func Default() *Spinner {
	return &Spinner{
		Message:    defaultMessage,
		Characters: defaultCharacters,
		Delay:      defaultDelay,
	}
}

//Spinner struct for the spinner entity.
type Spinner struct {
	stopChan chan struct{}
	//Progress message
	Message string
	//Spinning strings, Be creative here.
	Characters []string
	//Time in microseconds to wait between re-draws.
	Delay int
}

//Start Starts the spinning progress, by printing strings in s.Characters for given delay.
//
// This runs in a go routine and can be stopped by the stop function returned.
func (s *Spinner) Start() func() {
	s.stopChan = make(chan struct{}, 1)

	go func() {
		for {
			for _, val := range s.Characters {
				select {
				case <-s.stopChan:
					fmt.Printf("\r")
					return
				default:
					fmt.Printf("\r%s ... %s", s.Message, val)
					time.Sleep(200 * time.Millisecond)
				}
			}
		}
	}()
	return func() {
		s.stopChan <- struct{}{}
	}
}

//Start simplest function to start a spinner. Values are replaced with default values if 0 / nil / blank provided
func Start(message string, characters []string, delay int) func() {
	spinner := &Spinner{stopChan: make(chan struct{}, 1)}

	if message == "" {
		spinner.Message = defaultMessage
	} else {
		spinner.Message = message
	}

	if len(characters) == 0 {
		spinner.Characters = defaultCharacters
	} else {
		spinner.Characters = characters
	}

	if delay == 0 {
		spinner.Delay = defaultDelay
	} else {
		spinner.Delay = delay
	}

	spinner.Start()

	return func() {
		spinner.stopChan <- struct{}{}
	}
}
