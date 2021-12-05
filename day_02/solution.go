package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

type Command struct {
	direction string
	distance  int
}

func readCommands(filename string) []Command {
	file, err := os.Open(filename)
	check(err)
	defer file.Close()

	Scanner := bufio.NewScanner(file)

	var commands []Command
	for Scanner.Scan() {
		s := strings.Split(Scanner.Text(), " ")
		commands = append(commands, Command{s[0], toInt(s[1])})
	}
	return commands
}

func part1(commands []Command) int {
	horizontal := 0
	depth := 0
	for i := 0; i < len(commands); i++ {
		command := commands[i]
		switch command.direction {
		case "forward":
			horizontal += command.distance
		case "up":
			depth -= command.distance
		case "down":
			depth += command.distance
		}
	}
	return horizontal * depth
}

func part2(commands []Command) int {
	horizontal := 0
	depth := 0
	aim := 0
	for i := 0; i < len(commands); i++ {
		command := commands[i]
		switch command.direction {
		case "forward":
			horizontal += command.distance
			depth += command.distance * aim
		case "up":
			aim -= command.distance
		case "down":
			aim += command.distance
		}
	}
	return horizontal * depth
}

func main() {
	testInputData := readCommands("./day_02/test_input.txt")
	fmt.Println("Part 1 (test):", part1(testInputData))
	fmt.Println("Part 2 (test):", part2(testInputData))

	inputData := readCommands("./day_02/input.txt")
	fmt.Println("Part 1:", part1(inputData))
	fmt.Println("Part 2:", part2(inputData))
}
