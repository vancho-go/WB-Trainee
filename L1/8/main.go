// Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.
package main

import (
	"fmt"
	"strconv"
)

func setBitTo1(n int64, pos uint) int64 {
	n |= (1 << pos)
	return n
}

func setBitTo0(n int64, pos uint) int64 {
	n &^= (1 << pos)
	return n
}

func main() {
	x := int64(123)
	fmt.Printf("Before setting 2nd bit to 1: %s\n", strconv.FormatInt(x, 2))
	x = setBitTo1(x, 2)
	fmt.Printf("After setting 2nd bit to 1: %s\n", strconv.FormatInt(x, 2))
	x = setBitTo0(x, 1)
	fmt.Printf("After setting 1st bit to 0: %s\n", strconv.FormatInt(x, 2))
}
