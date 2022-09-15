// Разработать конвейер чисел. Даны два канала: в первый пишутся
// числа (x) из массива, во второй — результат операции x*2, после
// чего данные из второго канала должны выводиться в stdout.

package main

import (
	"fmt"
	"sync"
)

func writeToFirstCh(numbers []int, ch chan int) {
	defer close(ch)
	for _, elem := range numbers {
		ch <- elem
	}
}

func readFromFirstAndWriteToSecondCh(inputCh chan int, outputCh chan int) {
	defer close(inputCh)
	for elem := range inputCh {
		outputCh <- elem
	}
}

func printFromSecondCh(inputCh chan int, wg *sync.WaitGroup) {
	defer close(inputCh)
	for elem := range inputCh {
		fmt.Println(elem * 2)
		wg.Done()
	}
}

func main() {
	firstCh := make(chan int)
	secondCh := make(chan int)
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	wg := new(sync.WaitGroup)
	wg.Add(len(numbers))

	go writeToFirstCh(numbers, firstCh)
	go readFromFirstAndWriteToSecondCh(firstCh, secondCh)
	go printFromSecondCh(secondCh, wg)

	wg.Wait()
}
