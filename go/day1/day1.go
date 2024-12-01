package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	fp, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer fp.Close()
	scanner := bufio.NewScanner(fp)
	scanner.Split(bufio.ScanLines)
	var leftSide []int
	var rightSide []int
	for scanner.Scan() {
		entry := strings.Split(scanner.Text(), "   ")
		left, _ := strconv.Atoi(entry[0])
		right, _ := strconv.Atoi(entry[1])
		leftSide = append(leftSide, left)
		rightSide = append(rightSide, right)
	}
	slices.Sort(leftSide)
	slices.Sort(rightSide)
	sum := 0
	/*
		Part 1:
		for i := 0; i < len(leftSide); i++ {
			dist := rightSide[i] - leftSide[i]
			if dist < 0 {
				dist = -dist
			}
			sum += dist
		}
	*/
	for i := 0; i < len(leftSide); i++ {
		partial := 0
		for j := 0; j < len(rightSide); j++ {
			if leftSide[i] == rightSide[j] {
				partial++
			}
		}
		sum += partial * leftSide[i]
	}
	fmt.Println(sum)
}
