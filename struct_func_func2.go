package smtest

type sff2StateFunc func(*Machine2) sff2StateFunc

type Machine2 struct {
	total int64
}

func (m *Machine2) initialize() sff2StateFunc {
	m.total = 0
	return (*Machine2).checkEven
}

func (m *Machine2) checkEven() sff2StateFunc {
	if m.total%2 == 0 {
		return (*Machine2).addOne
	}
	return (*Machine2).addThree
}

func (m *Machine2) addOne() sff2StateFunc {
	m.total += 1
	return (*Machine2).checkFinished
}

func (m *Machine2) addThree() sff2StateFunc {
	m.total += 3
	return (*Machine2).checkFinished
}

func (m *Machine2) checkFinished() sff2StateFunc {
	if m.total > totalMax {
		return (*Machine2).finished
	}
	return (*Machine2).checkEven
}

func (m *Machine2) finished() sff2StateFunc {
	return nil
}

func runStructFuncFunc2() {
	m := new(Machine2)
	fn := (*Machine2).initialize
	for fn != nil {
		fn = fn(m)
	}
}
