package mocking

import (
	"fmt"
	"io"
	"os"
	"time"
)

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := 3; i > 0; i-- {
		fmt.Fprintf(out, "%d\n", i)
		sleeper.Sleep()
	}
	fmt.Fprintf(out, "Go!")
}

type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct{}

func (ds DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

func main() {
	sleeper := &ConfigurableSleeper{3 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}
