// К каким негативным последствиям может привести данный
// фрагмент кода, и как это исправить? Приведите корректный пример реализации.

// var justString string

// func someFunc() {
// 	v := createHugeString(1 << 10)
// 	justString = v[:100]
// }

// func main() {
// 	someFunc()
// }

package main

import (
	"fmt"
	"math/rand"
)

// Проблема 1: не реализована функция someFunc => не скомпилируется => необходимо ее реализовать
// Проблема 2: строка v может быть меньше 100 => выйдем за границы массива => паника => нужно обработать этот случай

// Возможное решение
func createHugeString(size int) string {
	res := make([]byte, size)
	for n := range res {
		res[n] = byte(rand.Intn(122-48) + 48)
	}
	return string(res)
}

func main() {
	v := createHugeString(1 << 10) // 1024

	if len(v)-1 >= 100 {
		justString := v[:100]
		fmt.Println(justString)
	} else {
		fmt.Println("Index out of range")
	}
}
