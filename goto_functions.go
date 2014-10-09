package main

import "fmt"

const totalMax = 1000000000

func stateInitialize() int64 {
	return 0
}

func stateCheckEven(total int64) bool {
	return total%2 == 0
}

func stateAddOne(total int64) int64 {
	return total + 1
}

func stateAddThree(total int64) int64 {
	return total + 3
}

func stateCheckFinished(total int64) bool {
	return total > totalMax
}

func main() {
	var total int64

	goto initialize

initialize:
	total = stateInitialize()
	goto checkEven

checkEven:
	if stateCheckEven(total) {
		goto addOne
	}
	goto addThree

addOne:
	total = stateAddOne(total)
	goto checkFinished

addThree:
	total = stateAddThree(total)
	goto checkFinished

checkFinished:
	if stateCheckFinished(total) {
		goto finished
	}
	goto checkEven

finished:
	fmt.Println(total)
}
