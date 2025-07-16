package main

import (
	"os"

	mocking "github.com/Daavipf/Learning-Golang/Mocking"
)

func main(){
	sleeper := &mocking.DefaultSleeper{}
	mocking.Countdown(os.Stdout, sleeper)
}
