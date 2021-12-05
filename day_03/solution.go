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

func readDiagnostics(filename string) []string {
	file, err := os.Open(filename)
	check(err)
	defer file.Close()

	Scanner := bufio.NewScanner(file)

	var diagnostics []string
	for Scanner.Scan() {
		diagnostics = append(diagnostics, Scanner.Text())
	}
	return diagnostics
}

func mostCommonBit(diagnostics []string, position int) int {
	totalOneBits := 0
	totalZeroBits := 0
	for _, diagnostic := range diagnostics {
		if diagnostic != "" {
			bitAtPosition := string(diagnostic)[position : position+1]
			if bitAtPosition == "1" {
				totalOneBits++
			} else {
				totalZeroBits++
			}
		}
	}
	if totalOneBits >= totalZeroBits {
		return 1
	} else {
		return 0
	}
}

func leastCommonBit(diagnostics []string, position int) int {
	if mostCommonBit(diagnostics, position) == 1 {
		return 0
	} else {
		return 1
	}
}

func diagnosticWithBitAtColumn(diagnostics []string, column int, bit int) []string {
	newDiagnostics := make([]string, 0)
	for _, diagnostic := range diagnostics {
		if diagnostic[column] == strconv.Itoa(bit)[0] {
			newDiagnostics = append(newDiagnostics, diagnostic)
		}
	}
	return newDiagnostics
}

func part1(diagnostics []string) int {
	gammaRate := ""
	epsilonRate := ""
	for i := 0; i < len(diagnostics[0]); i++ {
		gammaRate += strconv.Itoa(mostCommonBit(diagnostics, i))
		epsilonRate += strconv.Itoa(leastCommonBit(diagnostics, i))
	}
	gammaInt, _ := strconv.ParseInt(gammaRate, 2, 64)
	epsilonInt, _ := strconv.ParseInt(epsilonRate, 2, 64)

	return int(gammaInt) * int(epsilonInt)
}

func part2(diagnostics []string) int {
	oxygenValues := make([]string, len(diagnostics))
	copy(oxygenValues, diagnostics)
	co2Values := make([]string, len(diagnostics))
	copy(co2Values, diagnostics)
	length := len(diagnostics[0])

	for i := 0; i < length; i++ {
		oxygenValues = diagnosticWithBitAtColumn(oxygenValues, i, mostCommonBit(oxygenValues, i))
		if len(oxygenValues) == 1 {
			break
		}
	}

	for i := 0; i < length; i++ {
		co2Values = diagnosticWithBitAtColumn(co2Values, i, leastCommonBit(co2Values, i))
		if len(co2Values) == 1 {
			break
		}
	}

	oxygenRating := oxygenValues[0]
	co2Rating := co2Values[0]
	oxygenInt, _ := strconv.ParseInt(oxygenRating, 2, 64)
	co2Int, _ := strconv.ParseInt(co2Rating, 2, 64)

	return int(oxygenInt) * int(co2Int)
}

func main() {
	testInputData := readDiagnostics("./day_03/test_input.txt")
	fmt.Println("Part 1 (test):", part1(testInputData))
	fmt.Println("Part 2 (test):", part2(testInputData))

	inputData := readDiagnostics("./day_03/input.txt")
	fmt.Println("Part 1:", part1(inputData))
	fmt.Println("Part 2:", part2(inputData))
}
