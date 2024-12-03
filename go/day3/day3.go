package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	re := regexp.MustCompile(`mul\(\d+,\d+\)|don\'t\(\)|do\(\)`)
	occurs := re.FindAll(content, len(content))
	sum := 0
	do := true
	for _, occur := range occurs {
		occur := string(occur)
		fmt.Println(occur)
		if occur == "do()" {
			do = true
		} else if occur == "don't()" {
			do = false
		} else {
			if do {
				s := strings.Split(occur, ",")
				n1, _ := strconv.Atoi(s[0][4:])
				n2, _ := strconv.Atoi(s[1][:len(s[1])-1])
				sum += n1 * n2
			}
		}
	}
	fmt.Println(sum)
}
