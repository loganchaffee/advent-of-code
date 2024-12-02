package day1

import (
	"fmt"
	"math"
	"sort"
	"sync"
)

func Part1() int {
	listA, listB, err := MakeLists("input.txt")
	if err != nil {
		fmt.Println("Unable to parse input file")
	}

	var wg sync.WaitGroup

	sortArray := func(arr []int) {
		defer wg.Done()
		sort.Ints(arr)
	}

	wg.Add(2)
	go sortArray(listA)
	go sortArray(listB)
	wg.Wait()

	differences := []int{}

	for i := 0; i < len(listA); i++ {
		diff := int(math.Abs(float64(listA[i] - listB[i])))
		differences = append(differences, diff)
	}

	sum := 0
	for i := 0; i < len(differences); i++ {
		sum += differences[i]
	}

	return sum
}
