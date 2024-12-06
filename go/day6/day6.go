package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	fp, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(fp)
	scanner.Split(bufio.ScanLines)
	// 0, 90, 180, 270 deg
	carets := map[byte][2]int{
		'^': {-1, 0},
		'>': {0, 1},
		'v': {1, 0},
		'<': {0, -1},
	}
	turn := map[byte]byte{
		'^': '>',
		'>': 'v',
		'v': '<',
		'<': '^',
	}
	cp := [2]int{}
	var cc byte
	var mo [][]byte
	lineCount := 0
	for scanner.Scan() {
		row := scanner.Text() // IMPORTANT: use .Text(), because .Bytes() does no allocation
		for caret := range carets {
			colCount := strings.Index(string(row), string(caret))
			if colCount != -1 {
				cp[0] = lineCount
				cp[1] = colCount
				cc = caret
			}
		}
		mo = append(mo, []byte(row))
		lineCount++
	}

	height := len(mo)
	width := len(mo[0])
	/*
		Part 1:
		pos := map[[2]int]bool{
			currPos: true,
		}*/

	counter := 0

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if mo[i][j] == '#' || mo[i][j] == '^' {
				fmt.Println("skipped", counter, i, j)
				continue
			}
			posTurn := [][2]int{}
			posTurnCaret := []byte{}
			m := make([][]byte, len(mo))
			for k, line := range mo {
				m[k] = make([]byte, len(line))
				copy(m[k], line)
			}
			m[i][j] = '#'
			currPos := [2]int{cp[0], cp[1]}
			currCaret := cc
			isLooping := false
			// ticker := time.NewTicker(time.Millisecond * 50)
			start := time.Now()
			for time.Since(start) < time.Second*2 {
				m[currPos[0]][currPos[1]] = 'X'
				newPos := [2]int{
					currPos[0] + carets[currCaret][0],
					currPos[1] + carets[currCaret][1],
				}
				if !(newPos[0] >= 0 && newPos[0] < height && newPos[1] >= 0 && newPos[1] < width) {
					break
				}
				if m[newPos[0]][newPos[1]] == '#' {
					currCaret = turn[currCaret]
					posTurnIndex := -1
					for pti, pt := range posTurn {
						if pt[0] == currPos[0] && pt[1] == currPos[1] {
							posTurnIndex = pti
						}
					}
					if posTurnIndex != -1 && posTurnCaret[posTurnIndex] == currCaret {
						isLooping = true
						break
					}
					posTurn = append(posTurn, currPos)
					posTurnCaret = append(posTurnCaret, currCaret)
				} else {
					currPos = newPos
				}
				m[currPos[0]][currPos[1]] = currCaret
				/*out, _ := os.Create("output.txt")
				for _, line := range m {
					out.Write(append(line, '\n'))
				}
				out.Sync()
				out.Close()
				<-ticker.C*/
				// Part 1:
				// pos[currPos] = true
			}
			if time.Since(start) > time.Second*2 {
				isLooping = true
			}
			if isLooping {
				counter++
				fmt.Println("looping", counter, i, j, currPos, string(currCaret))
			} else {
				fmt.Println("xxxxxxx", counter, i, j, currPos, string(currCaret))
			}
		}
	}

	fmt.Println(counter)

	// Part 1
	// fmt.Println(len(pos))
}
