package smtest

func gfStateInitialize() int64 {
	return 0
}

func gfStateCheckEven(total int64) bool {
	return total%2 == 0
}

func gfStateAddOne(total int64) int64 {
	return total + 1
}

func gfStateAddThree(total int64) int64 {
	return total + 3
}

func gfStateCheckFinished(total int64) bool {
	return total > totalMax
}

func runGotoFunctions() {
	var total int64
	goto initialize

initialize:
	total = gfStateInitialize()
	goto checkEven

checkEven:
	if gfStateCheckEven(total) {
		goto addOne
	}
	goto addThree

addOne:
	total = gfStateAddOne(total)
	goto checkFinished

addThree:
	total = gfStateAddThree(total)
	goto checkFinished

checkFinished:
	if gfStateCheckFinished(total) {
		goto finished
	}
	goto checkEven

finished:
}
