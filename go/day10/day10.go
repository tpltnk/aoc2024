package main

import (
	"bufio"
	"fmt"
	"os"
)

var m [][]byte

func main() {
	fp, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(fp)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		m = append(m, []byte(line))
	}
	var startPositions [][2]int
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[0]); j++ {
			if m[i][j] == '0' {
				startPositions = append(startPositions, [2]int{i, j})
			}
		}
	}

	sum := 0

	for _, spos := range startPositions {
		fmt.Printf("Trailhead at (%d, %d)\n", spos[0], spos[1])
		partial, found := findPaths(spos[0], spos[1], 0)
		sum += partial
		fmt.Println("Final have:", found)
		fmt.Printf("End trailhead at (%d, %d) score is %d\n", spos[0], spos[1], partial)
	}

	fmt.Println(sum)
}

func suitableNeighbors(i, j int) [][2]int {
	height := int(m[i][j] - '0')
	var neigh [][2]int
	if i+1 < len(m) && int(m[i+1][j]-'0')-height == 1 {
		neigh = append(neigh, [2]int{i + 1, j})
	}
	if i-1 >= 0 && int(m[i-1][j]-'0')-height == 1 {
		neigh = append(neigh, [2]int{i - 1, j})
	}
	if j+1 < len(m[0]) && int(m[i][j+1]-'0')-height == 1 {
		neigh = append(neigh, [2]int{i, j + 1})
	}
	if j-1 >= 0 && int(m[i][j-1]-'0')-height == 1 {
		neigh = append(neigh, [2]int{i, j - 1})
	}
	return neigh
}

func findPaths(i, j, k int) (int, [][2]int) {
	if k == 9 {
		return 1, [][2]int{{i, j}}
	}
	sum := 0
	var foundCoords [][2]int
	for _, neigh := range suitableNeighbors(i, j) {
		partial, found := findPaths(neigh[0], neigh[1], k+1)
		alreadyFound := false

		for i := 0; i < len(foundCoords); i++ {
			for j := 0; j < len(found); j++ {
				if foundCoords[i][0] == found[j][0] && foundCoords[i][1] == found[j][1] {
					alreadyFound = true
					goto out
				}
			}
		}
	out:

		if !alreadyFound {
			foundCoords = append(foundCoords, found...)
			sum += partial
		}
	}
	return sum, foundCoords
}
