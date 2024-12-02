package main

import (
	"fmt"
	"os"

	"github.com/loganchaffee/advent-of-code/internal/day1"
)

func main() {
	args := os.Args

	if len(args) < 3 {
		fmt.Println("Please provide a program day and name e.g. day1 part1")
		return
	}

	dayAndPart := args[1] + args[2]

	switch dayAndPart {
	case "day1part1":
		fmt.Println(day1.Part1())
	case "day1part2":
		fmt.Println(day1.Part2())
	}
}
