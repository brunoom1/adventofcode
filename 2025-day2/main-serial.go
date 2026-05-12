package main

import (
	"bufio"
	"day2/idsCheck"
	"fmt"
	"os"
	"strings"
)

func serial_main() {

	reader := bufio.NewReader(os.Stdin)

	content, err := reader.ReadString('\n')
	content = content[0 : len(content)-1] // remove \n

	if err != nil {
		fmt.Println(err.Error())
	}

	intervalos := strings.Split(content, ",")

	var resultSomeIds int

	for _, intervalo := range intervalos {
		invalidsIds := idsCheck.GetInvalidIds(intervalo)
		resultSomeIds += idsCheck.SomeIds(invalidsIds)
	}

	fmt.Print(resultSomeIds)
}
