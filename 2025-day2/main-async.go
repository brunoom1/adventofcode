package main

import (
	"bufio"
	"day2/idsCheck"
	"fmt"
	"os"
	"strings"
	"sync"
)

func async_main() {

	reader := bufio.NewReader(os.Stdin)

	content, err := reader.ReadString('\n')
	content = content[0 : len(content)-1] // remove \n

	if err != nil {
		fmt.Println(err.Error())
	}

	intervalos := strings.Split(content, ",")
	var invalidsIds = make([]int, 0, 100)

	var wg sync.WaitGroup
	var mutex sync.Mutex

	wg.Add(len(intervalos))
	for _, intervalo := range intervalos {
		go idsCheck.AsyncGetInvalidIds(&invalidsIds, &mutex, intervalo, &wg)
	}
	wg.Wait()

	fmt.Print(idsCheck.SomeIds(invalidsIds))
}
