// Разработать программу, которая проверяет, что все символы в
// строке уникальные (true — если уникальные, false etc). Функция проверки должна быть регистронезависимой.
// Например:
// abcd — true
// abCdefAaf — false
// aabcd — false

package main

import (
	"fmt"
	"strings"
)

func checkUnique(s string) bool {
	lowerStr := strings.ToLower(s)
	res := map[string]int{}
	for _, char := range lowerStr {
		if _, ok := res[string(char)]; ok {
			return false
		} else {
			res[string(char)] += 1
		}
	}
	return true
}

func main() {
	s := "aabcd"
	fmt.Print(checkUnique(s))
}
