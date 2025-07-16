package mocking

import "time"

type DefaultSleeper struct{}

func (s *DefaultSleeper) Sleep(){
	time.Sleep(1 * time.Second)
}
