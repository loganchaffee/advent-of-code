package day1

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"
)

func MakeLists(filename string) ([]int, []int, error) {
	// Open file
	file, err := os.Open("./internal/day1/inputs/" + filename)
	if err != nil {
		return nil, nil, errors.New(err.Error())
	}
	defer file.Close()

	listA := []int{}
	listB := []int{}

	// Read input file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		nums := strings.Split(line, "   ")

		numA, err := strconv.Atoi(nums[0])
		if err != nil {
			panic("Something went wrong")
		}
		listA = append(listA, numA)

		numB, err := strconv.Atoi(nums[1])
		if err != nil {
			panic("Something went wrong")
		}
		listB = append(listB, numB)
	}

	return listA, listB, nil
}
