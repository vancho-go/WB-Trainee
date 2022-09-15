// Дана последовательность чисел: 2,4,6,8,10.
// Найти сумму их квадратов(22+32+42....) с использованием
// конкурентных вычислений.

package main

import (
	"fmt"
	"sync"
)

func square(n int) int {
	return n * n
}

func main() {
	wg := new(sync.WaitGroup)
	mu := new(sync.Mutex)
	sum := 0
	numbers := []int{2, 4, 6, 8, 10}

	for _, number := range numbers {
		wg.Add(1)
		number := number
		go func() {
			mu.Lock()
			defer wg.Done()
			defer mu.Unlock()
			sum = sum + square(number)
		}()
	}
	wg.Wait()
	fmt.Print(sum)
}
