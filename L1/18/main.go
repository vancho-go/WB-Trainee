// Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной
// среде. По завершению программа должна выводить итоговое значение счетчика.

package main

import (
	"fmt"
	"strconv"
	"sync"
)

func counterAdd(counter *int, ch chan int, wg *sync.WaitGroup) {
	for range ch {
		*counter++
	}
	wg.Done()
}
func anotherGoroutine(id string, ch chan int, wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		fmt.Printf("Doing something good, worker %s\n", id)
		ch <- 1
	}
	wg.Done()
}

func main() {
	worker := 3
	ch := make(chan int, worker)
	wg1 := new(sync.WaitGroup)
	wg2 := new(sync.WaitGroup)
	wg2.Add(1)
	counter := 0
	go counterAdd(&counter, ch, wg2)

	for i := 0; i < worker; i++ {
		wg1.Add(1)
		go anotherGoroutine(strconv.Itoa(i), ch, wg1)
	}
	wg1.Wait()
	close(ch)
	wg2.Wait()
	fmt.Println("Total: ", counter)
}
