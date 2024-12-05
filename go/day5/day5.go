package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fp, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	pors := [][2]int{}
	scanner := bufio.NewScanner(fp)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		por := scanner.Text()
		if por == "" {
			break
		}
		posSplit := strings.Split(por, "|")
		p1, _ := strconv.Atoi(posSplit[0])
		p2, _ := strconv.Atoi(posSplit[1])
		pors = append(pors, [2]int{p1, p2})
	}

	sum := 0
	for scanner.Scan() {
		update := scanner.Text()
		updateSplit := strings.Split(update, ",")
		u := []int{}
		for _, page := range updateSplit {
			p, _ := strconv.Atoi(page)
			u = append(u, p)
		}
		// part 1: isRightOrder := true
		wasRightOrder := true
		for i := 0; i < len(u); i++ {
			for j := 0; j < len(u); j++ {
				if i != j {
					for k := 0; k < len(pors); k++ {
						if pors[k][0] == u[i] && pors[k][1] == u[j] {
							if i > j {
								wasRightOrder = false
								u[i], u[j] = u[j], u[i]
								/*
									part 1:
									isRightOrder = false
									goto end
								*/
							}
						}

						if pors[k][0] == u[j] && pors[k][1] == u[i] {
							if i < j {
								wasRightOrder = false
								u[i], u[j] = u[j], u[i]
								/*
									isRightOrder = false
									goto end
								*/
							}
						}
					}
				}
			}
		}

		/*
				part 1:
			end:
				if isRightOrder {
		*/
		if !wasRightOrder {
			mid := u[(len(u)-1)/2]
			sum += mid
		}

		/*}*/
	}

	fmt.Println(sum)
}
