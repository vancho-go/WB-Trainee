// Дана структура Human (с произвольным набором полей и методов).
// Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).

package main

import "fmt"

type Human struct {
	Name    string
	Surname string
}

type Action struct {
	Type string
	Human
}

func (h *Human) Greeting() {
	fmt.Println("Hello, my name is " + h.Name + " " + h.Surname)
}

func (a *Action) TypeInfo() {
	fmt.Printf("My type is %s\n", a.Type)
}

func main() {
	John_Doe := Human{"John", "Doe"}
	John_Doe.Greeting()

	Kick := Action{"Kick", Human{"Ivan", "Ivanov"}}
	Kick.TypeInfo()

	Kick.Greeting()
}
