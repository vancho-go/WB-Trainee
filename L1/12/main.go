// Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.

package main

import "fmt"

type set map[string]struct{}

func (s set) add(str string) {
	s[str] = struct{}{}
}

func main() {
	strs := set{}
	test := []string{"cat", "cat", "dog", "cat", "tree"}
	for _, str := range test {
		strs.add(str)
	}
	fmt.Println(strs)
}
