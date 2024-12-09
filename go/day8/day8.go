package main

import (
	"bufio"
	"fmt"
	"os"
)

func isAntenna(b byte) bool {
	return (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z') || (b >= '0' && b <= '9')
}

func main() {
	fp, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(fp)
	scanner.Split(bufio.ScanLines)
	var m [][]byte
	for scanner.Scan() {
		line := scanner.Text()
		m = append(m, []byte(line))
	}
	antennas := map[byte][][2]int{}
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			if isAntenna(m[i][j]) {
				antennas[m[i][j]] = append(antennas[m[i][j]], [2]int{i, j})
			}
		}
	}
	uniqueAntinode := map[[2]int]bool{}
	for _, antennaPos := range antennas {
		for i := 0; i < len(antennaPos); i++ {
			for j := 0; j < len(antennaPos); j++ {
				if i != j {

					diff := [2]int{
						antennaPos[j][0] - antennaPos[i][0],
						antennaPos[j][1] - antennaPos[i][1],
					}

					for k := 0; ((antennaPos[j][0] + k*diff[0]) >= 0) && ((antennaPos[j][0] + k*diff[0]) < len(m)) && ((antennaPos[j][1] + k*diff[1]) >= 0) && ((antennaPos[j][1] + k*diff[1]) < len(m[0])); k++ {
						antinodePos := [2]int{
							antennaPos[j][0] + k*diff[0],
							antennaPos[j][1] + k*diff[1],
						}
						uniqueAntinode[antinodePos] = true
					}

					/*
						Part 1:
						diff := [2]int{
							antennaPos[j][0] - antennaPos[i][0],
							antennaPos[j][1] - antennaPos[i][1],
						}


						antinodePos := [2]int{
							antennaPos[j][0] + diff[0],
							antennaPos[j][1] + diff[1],
						}



						if antinodePos[0] >= len(m) || antinodePos[0] < 0 {
							continue
						}

						if antinodePos[1] >= len(m[0]) || antinodePos[1] < 0 {
							continue
						}

						uniqueAntinode[antinodePos] = true
					*/
				}
			}
		}
	}

	fmt.Println(len(uniqueAntinode))
}
