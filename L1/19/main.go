// Разработать программу, которая переворачивает подаваемую на ход строку (например: «главрыба — абырвалг»). Символы могут быть unicode.

package main

import "fmt"

func revert(s []rune) {
	for i := len(s) - 1; i >= 0; i-- {
		fmt.Print(string(s[i]))
	}
}

func main() {
	revert([]rune("главрыба"))
}
