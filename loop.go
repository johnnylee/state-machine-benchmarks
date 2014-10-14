package smtest

const (
	STATE_INIT int = iota
	STATE_CHECK_EVEN
	STATE_ADD_ONE
	STATE_ADD_THREE
	STATE_CHECK_FINISHED
	STATE_FINISHED
)

func runLoop() {
	var total int64
	state := STATE_INIT

loop:
	for {
		switch state {
		case STATE_INIT:
			total = 0
			state = STATE_CHECK_EVEN
		case STATE_CHECK_EVEN:
			if total%2 == 0 {
				state = STATE_ADD_ONE
			} else {
				state = STATE_ADD_THREE
			}
		case STATE_ADD_ONE:
			total += 1
			state = STATE_CHECK_FINISHED
		case STATE_ADD_THREE:
			total += 3
			state = STATE_CHECK_FINISHED
		case STATE_CHECK_FINISHED:
			if total > totalMax {
				state = STATE_FINISHED
			} else {
				state = STATE_CHECK_EVEN
			}
		case STATE_FINISHED:
			break loop
		}
	}
}
