package main

import (
	"fmt"
	//"github.com/davecheney/profile"
)

const totalMax = 1000000000

type stateFunc func() stateFunc

type Machine struct {
	total int64
}

func (m *Machine) initialize() stateFunc {
	m.total = 0
	return m.checkEven
}

func (m *Machine) checkEven() stateFunc {
	if m.total%2 == 0 {
		return m.addOne
	}
	return m.addThree
}

func (m *Machine) addOne() stateFunc {
	m.total += 1
	return m.checkFinished
}

func (m *Machine) addThree() stateFunc {
	m.total += 3
	return m.checkFinished
}

func (m *Machine) checkFinished() stateFunc {
	if m.total > totalMax {
		return m.finished
	}
	return m.checkEven
}

func (m *Machine) finished() stateFunc {
	fmt.Println(m.total)
	return nil
}

func main() {
	//defer profile.Start(profile.CPUProfile).Stop()
	m := new(Machine)
	fn := m.initialize
	for fn != nil {
		fn = fn()
	}
}
