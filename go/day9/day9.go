package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strconv"
)

//go:embed "input.txt"
var dm string

func main() {
	var m []byte
	var p []int

	id := 0
	for i := 0; i < len(dm); i++ {
		c := int(dm[i] - '0')
		if i%2 == 0 {
			// file
			ms := slices.Repeat([]byte(strconv.Itoa(id)), c)
			p = append(p, slices.Repeat([]int{id}, len(ms))...)
			m = append(m, ms...)
			id++
		} else {
			// free space
			ms := slices.Repeat([]byte{'.'}, c)
			p = append(p, slices.Repeat([]int{-1}, len(ms))...)
			m = append(m, ms...)
		}
	}

	for i := len(m) - 1; i >= 0; i-- {
		if m[i] == '.' {
			continue
		}
		for j := 0; j < len(m); j++ {
			if m[j] == '.' && i > j {
				m[j], m[i] = m[i], m[j]
				p[j], p[i] = p[i], p[j]
				break
			}
		}
		allGood := true
		for j := 0; j <= i; j++ {
			if m[j] == '.' {
				allGood = false
			}
		}
		if allGood {
			break
		}
	}

	var checksum int64 = 0
	for i := 0; i < len(m); i++ {
		if m[i] != '.' {
			checksum += int64(i * p[i])
		}
	}

	fmt.Println(checksum)
}
