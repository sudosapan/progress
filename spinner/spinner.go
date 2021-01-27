//Package spinner Package for spinning progress markers. Like in place spinning sticks
package spinner

import (
	"fmt"
	"time"
)

//Default return spinner with message = "In progress" and spinning sticks.
func Default() *Spinner {
	return &Spinner{
		Message:    "In Progress",
		Characters: []string{"-", "\\", "|", "/"},
		Delay:      200,
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
