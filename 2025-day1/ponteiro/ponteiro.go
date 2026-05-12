package ponteiro

const MAX = 100

var ponteiro int = 50

type GiroCallback func(int)

var onGiro GiroCallback

/**
 *     350 / 3 =
 */

func GetPonteiro() int {
	return ponteiro
}

func SetPonteiro(n int) {
	ponteiro = n
}

func GirarHorario(quantidade int) {
	ponteiro = ponteiro + quantidade
	if ponteiro >= MAX {
		ponteiro = ponteiro % MAX
	}

	if onGiro != nil {
		onGiro(ponteiro)
	}
}

func Reset() {
	ponteiro = 50
}

func GirarAntiHorario(quantidade int) {
	ponteiro = ponteiro - quantidade

	if ponteiro < 0 {
		ponteiro = MAX - (ponteiro*-1)%MAX
	}

	if onGiro != nil {
		onGiro(ponteiro)
	}
}

func OnGiro(callback GiroCallback) {
	onGiro = callback
}
