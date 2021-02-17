//Package spinner Package for spinning progress markers. Like in place spinning sticks
package spinner

import (
	"fmt"
	"time"
)

var (
	defaultMessage    = "In Progress"
	defaultCharacters = []rune{'-', '\\', '|', '/'}
	defaultDelay      = 200
)

//Spinner struct for the spinner entity.
type Spinner struct {
	stopChan chan struct{}
	//Progress message
	Message string
	//Characters to cycle between.
	Characters []rune
	//Lesser value faster animation.
	Delay int
}

//Start Starts the spinning progress, by printing strings in s.Characters for given delay.
//
// Stop the spinner with Spinner.Stop.
func (s *Spinner) Start() {
	s.swapDefaults()
	s.stopChan = make(chan struct{}, 1)

	go func() {
		for {
			for _, val := range s.Characters {
				select {
				case <-s.stopChan:
					fmt.Printf("\r")
					return
				default:
					fmt.Printf("\r%s ... %c", s.Message, val)
					time.Sleep(time.Duration(s.Delay) * time.Millisecond)
				}
			}
		}
	}()
}

//Stop Function to stop the spinner. Does nothing spinner start has not been called yet.
func (s *Spinner) Stop() {
	if s.stopChan == nil {
		// Spinner not yet started
		return
	}
	s.stopChan <- struct{}{}
}

func (s *Spinner) swapDefaults() {
	if s.Message == "" {
		s.Message = defaultMessage
	}
	if len(s.Characters) == 0 {
		s.Characters = defaultCharacters
	}
	if s.Delay <= 0 {
		s.Delay = defaultDelay
	}
}
