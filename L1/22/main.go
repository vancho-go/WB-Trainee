// Разработать программу, которая перемножает, делит, складывает,
// вычитает две числовых переменных a,b, значение которых > 2^20.

package main

import (
	"fmt"
	"math/big"
)

func mul(x *big.Int, y *big.Int) *big.Int {
	res := new(big.Int).Mul(x, y)
	return res
}

func div(x *big.Int, y *big.Int) *big.Int {
	res := new(big.Int).Div(x, y)
	return res
}

func add(x *big.Int, y *big.Int) *big.Int {
	res := new(big.Int).Add(x, y)
	return res
}

func sub(x *big.Int, y *big.Int) *big.Int {
	res := new(big.Int).Sub(x, y)
	return res
}

func main() {
	q := new(big.Int)
	q.SetString("10000000000000000000", 10)

	w := new(big.Int)
	w.SetString("20000000000000000000", 10)

	fmt.Printf("%v * %v = %v\n", q, w, mul(q, w))
	fmt.Printf("%v / %v = %v\n", q, w, div(q, w))
	fmt.Printf("%v + %v = %v\n", q, w, add(q, w))
	fmt.Printf("%v - %v = %v\n", q, w, sub(q, w))
}
