package day1

import "fmt"

func Part2() int {
	listA, listB, err := MakeLists("input.txt")
	if err != nil {
		fmt.Println("Unable to parse input file")
	}

	hashmap := map[int]int{}
	scores := []int{}

	for _, num := range listB {
		hashmap[num] = hashmap[num] + 1
	}

	for _, num := range listA {
		count := hashmap[num]

		scores = append(scores, num*count)
	}

	sum := 0
	for i := 0; i < len(scores); i++ {
		sum += scores[i]
	}

	return sum
}
