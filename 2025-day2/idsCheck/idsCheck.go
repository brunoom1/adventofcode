package idsCheck

import (
	"strconv"
	"strings"
	"sync"
)

func GetIntervalParts(intervalo string) (int, int) {
	intervaloPartes := strings.Split(intervalo, "-")
	p1, err := strconv.Atoi(intervaloPartes[0])

	if err != nil {
		panic(err)
	}

	p2, err := strconv.Atoi(intervaloPartes[1])

	if err != nil {
		panic(err)
	}

	return p1, p2
}

func GetInvalidIds(intervalo string) []int {
	p1, p2 := GetIntervalParts(intervalo)

	invalidsIds := make([]int, 10)

	for i := p1; i <= p2; i++ {
		if IsInvalidId(i) {
			invalidsIds = append(invalidsIds, i)
		}
	}

	return invalidsIds
}

func IsInvalidId(id int) bool {
	idStr := strconv.Itoa(id)

	if len(idStr)%2 != 0 {
		return false
	}

	p1, p2 := idStr[0:len(idStr)/2], idStr[len(idStr)/2:]

	if p1 != p2 {
		return false
	}

	return true
}

func SomeIds(invalidsIds []int) int {
	var some int

	for _, id := range invalidsIds {
		some += id
	}

	return some
}

func AsyncGetInvalidIds(slc *[]int, mutex *sync.Mutex, intervalo string, wg *sync.WaitGroup) {
	ids := GetInvalidIds(intervalo)

	mutex.Lock()
	for _, id := range ids {
		*slc = append(*slc, id)
	}
	mutex.Unlock()

	wg.Done()
}
