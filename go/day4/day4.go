package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fp, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scan := bufio.NewScanner(fp)
	scan.Split(bufio.ScanLines)
	var lines []string
	for scan.Scan() {
		lines = append(lines, scan.Text())
	}

	counter := 0

	/*
		Part 1:
		for i := 0; i < len(lines); i++ {
			for j := 0; j < len(lines[i]); j++ {
	*/

	for i := 1; i < len(lines)-1; i++ {
		for j := 1; j < len(lines[i])-1; j++ {
			topLeft := lines[i-1][j-1]
			topRight := lines[i-1][j+1]
			middle := lines[i][j]
			bottomLeft := lines[i+1][j-1]
			bottomRight := lines[i+1][j+1]
			mainDiag := string([]byte{topLeft, middle, bottomRight})
			otherDiag := string([]byte{topRight, middle, bottomLeft})

			if (mainDiag == "MAS" || mainDiag == "SAM") && (otherDiag == "MAS" || otherDiag == "SAM") {
				counter++
			}

			/*
				Part 1:
				// left right
				if j <= len(lines[i])-4 {
					if lines[i][j:j+4] == "XMAS" || lines[i][j:j+4] == "SAMX" {
						counter++
					}
				}

				// up down
				if i >= 3 {
					var sb strings.Builder
					for p := i; p >= i-3; p-- {
						sb.WriteByte(lines[p][j])
					}
					if sb.String() == "XMAS" || sb.String() == "SAMX" {
						counter++
					}
				}

				// main diagonal
				if j <= len(lines[i])-4 && i <= len(lines)-4 {
					var sb strings.Builder
					for p, r := i, j; r < j+4; p, r = p+1, r+1 {
						sb.WriteByte(lines[r][p])
					}
					if sb.String() == "XMAS" || sb.String() == "SAMX" {
						counter++
					}
				}

				// other diagonal
				if j >= 3 && i <= len(lines)-4 {
					var sb strings.Builder
					for p, r := i, j; r >= j-3; p, r = p+1, r-1 {
						sb.WriteByte(lines[r][p])
					}
					if sb.String() == "XMAS" || sb.String() == "SAMX" {
						counter++
					}
				}*/
		}
	}
	fmt.Println(counter)
}
