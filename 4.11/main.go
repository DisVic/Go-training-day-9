package main

import (
	"fmt"
	"strconv"
)

func CountLetters(str string) bool {
	for i := 0; i < len(str); i++ {
		for j := i + 1; j < len(str); j++ {
			if str[i] == str[j] {
				return false
			}
		}
	}

	return true
}

func main() {
	var year int
	_, _ = fmt.Scan(&year)
	for nextYear := year; nextYear <= 9500; nextYear++ {
		str := strconv.Itoa(nextYear)
		if nextYear > year && CountLetters(str) == true {
			fmt.Println(nextYear)
			break
		}
	}
}
