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
	scanner := bufio.NewScanner(fp)
	scanner.Split(bufio.ScanLines)
	safe_counter := 0
	for scanner.Scan() {
		report := scanner.Text()
		levels := strings.Split(report, " ")
		safe := false
		for i := 0; i < len(levels); i++ {
			sub := []string{}
			sub = append(sub, levels[:i]...)
			sub = append(sub, levels[i+1:]...)
			if is_safe(sub) {
				safe = true
				break
			}
		}
		if safe {
			safe_counter++
		}
	}
	fmt.Println(safe_counter)
}

func is_safe(levels []string) bool {
	firstLevel, _ := strconv.Atoi(levels[0])
	secondLevel, _ := strconv.Atoi(levels[1])
	slope := secondLevel - firstLevel
	for i := 0; i < len(levels)-1; i++ {
		currentLevel, _ := strconv.Atoi(levels[i])
		nextLevel, _ := strconv.Atoi(levels[i+1])
		diff := nextLevel - currentLevel
		if slope < 0 && diff > 0 {
			return false
		}
		if slope > 0 && diff < 0 {
			return false
		}
		if diff < 0 {
			diff = -diff
		}
		if diff > 3 || diff < 1 {
			return false
		}
	}
	return true
}
