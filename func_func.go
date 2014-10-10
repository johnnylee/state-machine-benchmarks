package smtest

import ()

type ffStateFunc func(*int64) ffStateFunc

func ffStateInitialize(total *int64) ffStateFunc {
	*total = 0
	return ffStateCheckEven
}

func ffStateCheckEven(total *int64) ffStateFunc {
	if *total%2 == 0 {
		return ffStateAddOne
	}
	return ffStateAddThree
}

func ffStateAddOne(total *int64) ffStateFunc {
	*total += 1
	return ffStateCheckFinished
}

func ffStateAddThree(total *int64) ffStateFunc {
	*total += 3
	return ffStateCheckFinished
}

func ffStateCheckFinished(total *int64) ffStateFunc {
	if *total > totalMax {
		return ffStateFinished
	}
	return ffStateCheckEven
}

func ffStateFinished(total *int64) ffStateFunc {
	return nil
}

func runFuncFunc() {
	var total int64
	fn := ffStateInitialize
	for fn != nil {
		fn = fn(&total)
	}
}
