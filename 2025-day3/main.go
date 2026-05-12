package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func readBank() []string {
	bank := make([]string, 0)

	reader := bufio.NewReader(os.Stdin)

	for {
		line, err := reader.ReadString('\n')

		if err != nil {
			break
		}

		line = line[0 : len(line)-1]
		bank = append(bank, line)
	}

	return bank
}

func main() {
	bankBatery := readBank()
	var totalBatery int

	for _, batery := range bankBatery {
		var n1Pos int
		var n1, n2 int

		for i := 0; i < len(batery)-1; i++ {
			n, err := strconv.Atoi(batery[i : i+1])

			if err != nil {
				panic(fmt.Sprintf("Não foi possível converter numero %s", batery[i:i+1]))
			}

			if n > n1 {
				n1 = n
				n1Pos = i
			}
		}

		for j := n1Pos + 1; j < len(batery); j++ {

			n, err := strconv.Atoi(batery[j : j+1])

			if err != nil {
				panic("Não foi possível converter numero")
			}

			if n > n2 {
				n2 = n
			}
		}
		// fmt.Printf("%d:%d, %d:%d\n", n1Pos, n1, n2Pos, n2)
		totalBatery = totalBatery + (n1*10 + n2)
	}

	fmt.Print(totalBatery)

}
