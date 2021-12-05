package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func toInt(s string) int {
	result, err := strconv.Atoi(s)
	check(err)
	return result
}

func readDepths(filename string) []int {
	file, err := os.Open(filename)
	check(err)
	defer file.Close()

	Scanner := bufio.NewScanner(file)

	var numbers []int
	for Scanner.Scan() {
		numbers = append(numbers, toInt(Scanner.Text()))
	}
	return numbers
}

func countSlidingWindowIncrease(depths []int, windowSize int) int {
	count := 0
	for i := 0; i < len(depths)-windowSize; i++ {
		if depths[i+windowSize] > depths[i] {
			count++
		}
	}
	return count
}

func part1(inputData []int) int {
	return countSlidingWindowIncrease(inputData, 1)
}

func part2(inputData []int) int {
	return countSlidingWindowIncrease(inputData, 3)
}

func main() {
	testInputData := readDepths("./day_01/test_input.txt")
	fmt.Println("Part 1 (test):", part1(testInputData))
	fmt.Println("Part 2 (test):", part2(testInputData))

	inputData := readDepths("./day_01/input.txt")
	fmt.Println("Part 1:", part1(inputData))
	fmt.Println("Part 2:", part2(inputData))
}
