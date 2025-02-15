package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
	"time"
)

//go:embed input.test.txt
var input string

const blinks = 4

func main() {
	stones := strings.Split(input, " ")
	ticker := time.NewTicker(time.Second * 1)

	var line [][]string
	for _, stone := range stones {
		line = append(line, []string{stone})
	}

	for c := 0; c < blinks; c++ {
		for i := 0; i < len(line); i++ {
			var newSubline []string
			for _, stone := range line[i] {
				cleanStone := removeLeadingZeros(stone)
				stoneNumber, _ := strconv.Atoi(cleanStone)
				if stoneNumber == 0 {
					newSubline = append(newSubline, "1")
					continue
				}
				if len(cleanStone)%2 == 0 {
					newSubline = append(newSubline, removeLeadingZeros(cleanStone[:len(cleanStone)/2]), " ")
					newSubline = append(newSubline, removeLeadingZeros(cleanStone[len(cleanStone)/2:]), " ")
					continue
				}
				newSubline = append(newSubline, strconv.Itoa(stoneNumber*2024))
			}
			fmt.Println(i, "---", newSubline)
			line[i] = newSubline
		}
		<-ticker.C
	}

	sum := 0

	for _, stones := range line {
		sum += len(stones)
	}

	fmt.Println(sum)
}

func removeLeadingZeros(num string) string {
	cleanNum := "0"
	for i := 0; i < len(num); i++ {
		if num[i] != '0' {
			cleanNum = num[i:]
			break
		}
	}
	return cleanNum
}
