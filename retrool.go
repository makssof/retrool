package retrool

import (
	"time"
)

// The TryOptions structure describes how attempts should be made to perform an action.
// There are default options: DefaultTryOptions
type TryOptions struct {
	// StartInterval is the initial delay after the first failed
	// attempt (if the first attempt fails). The base number to which
	// Addition will be added with multiplication by AdditionCoefficient.
	StartInterval float32

	// With each new attempt, Addition (multiplied by AdditionCoefficient) will be added to StartInterval
	Addition float32

	// With each new attempt, Addition multiplied by AdditionCoefficient will be added to StartInterval.
	AdditionCoefficient float32

	// How many attempts to make before it is considered completely failed
	MaxTries int

	// The FailureDecisionMaker function takes as input the result of executing the passed function
	// and determines whether the result is a failure. If yes, it returns false, otherwise true.
	// If this function returns false, the attempts will continue (if the counter does not exceed MaxTries)
	//
	// There is a default function DefaultDecisionMaker. If no other function is specified, this one will be used.
	// It expects a boolean value as input and returns true if the input value is true
	FailureDecisionMaker func(interface{}) bool
}

var (
	DefaultTryOptions = &TryOptions{
		StartInterval:        1,
		Addition:             1,
		AdditionCoefficient:  0.5,
		MaxTries:             5,
		FailureDecisionMaker: DefaultDecisionMaker,
	}
)

// DefaultDecisionMaker is the default function to detect failed attempt
// It expects a boolean value as input and returns true if the input value is true
func DefaultDecisionMaker(res interface{}) bool {
	boolVal, isBool := res.(bool)

	return isBool && boolVal
}

// The Try behavior is described via parameters. See TryOptions
func Try(opts *TryOptions, fnc func(int) bool) bool {
	var tryResult bool
	var currentAddition float32 = 0
	tryCounter := 1
	decisionMaker := DefaultDecisionMaker

	if opts.FailureDecisionMaker != nil {
		decisionMaker = opts.FailureDecisionMaker
	}

	tryResult = fnc(tryCounter)

	if decisionMaker(tryResult) {
		return true
	}

	for {
		if tryCounter > opts.MaxTries-1 {
			return false
		}

		msToSleep := (opts.StartInterval + currentAddition*opts.AdditionCoefficient) * 1000 // sec2ms
		timeToSlip := time.Millisecond * time.Duration(msToSleep)
		currentAddition += opts.Addition

		time.Sleep(timeToSlip)

		tryCounter++

		tryResult = fnc(tryCounter)

		if decisionMaker(tryResult) {
			return true
		}
	}
}
