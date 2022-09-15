// Написать программу, которая конкурентно рассчитает значение квадратов чисел взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.

package main

import (
	"fmt"
	"sync"
)

func square(n int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(n * n)
}

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	wg := new(sync.WaitGroup)
	for _, num := range numbers {
		wg.Add(1)
		go square(num, wg)
	}
	wg.Wait()
}
