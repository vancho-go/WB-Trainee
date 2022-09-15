// Реализовать собственную функцию sleep

package main

import (
	"fmt"
	"time"
)

func ownSleep(t int) {
	d := time.Second * time.Duration(t)
	select {
	case <-time.After(d):
	}
}

func main() {
	fmt.Println("Before sleep")
	ownSleep(3)
	fmt.Println("After sleep")
}
