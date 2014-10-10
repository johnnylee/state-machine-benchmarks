package smtest

type sffStateFunc func() sffStateFunc

type Machine struct {
	total int64
}

func (m *Machine) initialize() sffStateFunc {
	m.total = 0
	return m.checkEven
}

func (m *Machine) checkEven() sffStateFunc {
	if m.total%2 == 0 {
		return m.addOne
	}
	return m.addThree
}

func (m *Machine) addOne() sffStateFunc {
	m.total += 1
	return m.checkFinished
}

func (m *Machine) addThree() sffStateFunc {
	m.total += 3
	return m.checkFinished
}

func (m *Machine) checkFinished() sffStateFunc {
	if m.total > totalMax {
		return m.finished
	}
	return m.checkEven
}

func (m *Machine) finished() sffStateFunc {
	return nil
}

func runStructFuncFunc() {
	m := new(Machine)
	fn := m.initialize
	for fn != nil {
		fn = fn()
	}
}
