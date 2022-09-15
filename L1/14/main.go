// Разработать программу, которая в рантайме способна определить тип переменной: int, string, bool, channel из переменной типа interface{}.

package main

import "fmt"

func checkType(typee interface{}) {
	fmt.Printf("Value: %v Type: %T\n", typee, typee)
}

func main() {
	checkType("1")
	checkType(1)
	checkType(1.1)
}
