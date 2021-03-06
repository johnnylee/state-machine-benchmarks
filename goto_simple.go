package smtest

func runGotoSimple() {
	var total int64
	goto initialize

initialize:
	total = 0
	goto checkEven

checkEven:
	if total%2 == 0 {
		goto addOne
	}
	goto addThree

addOne:
	total += 1
	goto checkFinished

addThree:
	total += 3
	goto checkFinished

checkFinished:
	if total > totalMax {
		goto finished
	}
	goto checkEven

finished:
}
