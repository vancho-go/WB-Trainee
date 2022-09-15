// Удалить i-ый элемент из слайса.
package main

import "fmt"

func removeFromArray(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}

func main() {
	slice := []string{"first", "second", "third"}
	fmt.Print(removeFromArray(slice, 0))
}
