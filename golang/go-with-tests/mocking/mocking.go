package mocking

import (
	"fmt"
	"io"
	"time"
)

const ( 
	CountdownStart = 3
	CountdownFinalWord = "!"
)

type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct {}

func (d DefaultSleeper) Sleep()  {
	time.Sleep(1 * time.Second)	
}

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep()  {
	s.Calls++
	
}

func Countdown(out io.Writer, sleeper Sleeper)  {
	for i := CountdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}
	fmt.Fprint(out, CountdownFinalWord)
	
}
