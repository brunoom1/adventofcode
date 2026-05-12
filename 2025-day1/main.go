package main

import (
	"adventofcode/ponteiro"
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	line = line[0 : len(line)-1]

	strings := strings.Split(line, " ")

	countZeros := 0

	for _, comando := range strings {

		comandoDirecao := comando[0:1]
		comandoValor := comando[1:]
		comandoValotInt, _ := strconv.Atoi(comandoValor)

		switch comandoDirecao {
		case "R":
			for i := 0; i < comandoValotInt; i++ {
				ponteiro.GirarHorario(1)
			}
			if ponteiro.GetPonteiro() == 0 {
				countZeros++
			}
		case "L":
			for i := 0; i < comandoValotInt; i++ {
				ponteiro.GirarAntiHorario(1)
			}
			if ponteiro.GetPonteiro() == 0 {
				countZeros++
			}
		}
	}

	print(countZeros)
}
