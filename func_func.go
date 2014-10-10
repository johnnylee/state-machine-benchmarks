package main

import (
	"fmt"
	//"github.com/davecheney/profile"
)

const totalMax = 1000000000

type stateFunc func(*int64) stateFunc

func stateInitialize(total *int64) stateFunc {
	*total = 0
	return stateCheckEven
}

func stateCheckEven(total *int64) stateFunc {
	if *total%2 == 0 {
		return stateAddOne
	}
	return stateAddThree
}

func stateAddOne(total *int64) stateFunc {
	*total += 1
	return stateCheckFinished
}

func stateAddThree(total *int64) stateFunc {
	*total += 3
	return stateCheckFinished
}

func stateCheckFinished(total *int64) stateFunc {
	if *total > totalMax {
		return stateFinished
	}
	return stateCheckEven
}

func stateFinished(total *int64) stateFunc {
	fmt.Println(*total)
	return nil
}

func main() {
	//defer profile.Start(profile.CPUProfile).Stop()
	var total int64
	fn := stateInitialize
	for fn != nil {
		fn = fn(&total)
	}
}
