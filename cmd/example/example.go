package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/makssof/retrool"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Printf("Start at: %v\n", time.Now())

	// tryOptions := retrool.DefaultTryOptions
	tryOptions := &retrool.TryOptions{
		StartInterval:        1,
		Addition:             1,
		AdditionCoefficient:  0.5,
		MaxTries:             10,
		FailureDecisionMaker: retrool.DefaultDecisionMaker,
	}

	var lastExecTime *time.Time
	var n int
	startTime := time.Now()

	retrool.Try(tryOptions, func(i int) bool {
		now := time.Now()

		if lastExecTime == nil {
			lastExecTime = &now
		}

		n = rand.Intn(101)

		fmt.Printf("Try #%d | N=%d | Elapsed since last attempt: %v | "+
			"Total time elapsed: %v\n", i, n, now.Sub(*lastExecTime), now.Sub(startTime))

		*lastExecTime = time.Now()
		return n < 10
	})

	fmt.Printf("Finish at: %v\n", time.Now().Sub(startTime))
}
