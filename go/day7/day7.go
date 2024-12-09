package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	fp, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(fp)
	scanner.Split(bufio.ScanLines)

	sum := 0
	for scanner.Scan() {
		eq := scanner.Text()
		eqSplit := strings.Split(eq, ": ")
		testValue, _ := strconv.Atoi(eqSplit[0])
		var values []int
		for _, value := range strings.Split(strings.Join(eqSplit[1:], ""), " ") {
			valueInt, _ := strconv.Atoi(value)
			values = append(values, valueInt)
		}
		possibleValues := possible(values)
		if slices.Contains(possibleValues, testValue) {
			sum += testValue
		}
	}

	fmt.Println(sum)
}

func possible(values []int) []int {
	if len(values) == 1 {
		return values
	}
	concat := strconv.Itoa(values[0]) + strconv.Itoa(values[1])
	concatInt, _ := strconv.Atoi(concat)

	additionPossibilities := possible(append([]int{values[0] + values[1]}, values[2:]...))
	multiplicationPossibilities := possible(append([]int{values[0] * values[1]}, values[2:]...))

	// Part 2
	concatenationPossiblities := possible(append([]int{concatInt}, values[2:]...))

	return append(
		additionPossibilities,
		append(
			multiplicationPossibilities,
			concatenationPossiblities...,
		)...,
	)
}
