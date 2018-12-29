// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 244.

// Countdown implements the countdown for a rocket launch.
package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	abort := make(chan struct{})

	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		abort <- struct{}{}
	}()

	go func() {
		tick := time.Tick(1 * time.Second)
		for countdown := 10; countdown > 0; countdown-- {
			fmt.Println(countdown)
			<-tick
		}
	}()

	fmt.Println("Commencing countdown.  Press return to abort.")

	x := <-abort
  if x == struct{}{} {
    fmt.Println("Launch aborted!")
    return
  }

	fmt.Println("Lift off!")
}
