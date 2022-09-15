// Реализовать пересечение двух неупорядоченных множеств.

package main

import (
	"fmt"
)

func getIntersection(a []string, b []string, mode byte) []string {
	m := make(map[string]byte)

	for _, k := range a {
		m[k] += 1
	}

	for _, k := range b {
		m[k] -= 1
	}

	result := []string{}

	for k, v := range m {
		if v == 0 {
			result = append(result, k)
		}
	}

	return result
}

func main() {

	a := []string{"a", "b", "c", "d"}
	b := []string{"c", "d", "e", "f"}

	fmt.Println("left & right", getIntersection(a, b, 3))
}
