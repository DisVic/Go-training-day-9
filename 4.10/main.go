package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int
	_, _ = fmt.Scan(&n)
	var helper int
	var str string
	str = strconv.Itoa(n)
	for len(str) != 1 {
		str = strconv.Itoa(n)
		for n != 0 {
			helper += n % 10
			n /= 10
		}
		n = helper
		helper = 0
	}
	fmt.Println(n)
}
