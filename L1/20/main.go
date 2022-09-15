// Разработать программу, которая переворачивает слова в строке. Пример: «snow dog sun — sun dog snow».

package main

import (
	"fmt"
	"strings"
)

func revert2(s string) {
	words := strings.Split(s, " ")
	for i := len(words) - 1; i >= 0; i-- {
		fmt.Printf("%s ", words[i])
	}
}

func main() {
	revert2("snow dog sun")
}
