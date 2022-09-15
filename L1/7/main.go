// Реализовать конкурентную запись данных в map.

package main

import (
	"fmt"
	"strconv"
	"sync"
)

type asyncMap struct {
	mx sync.Mutex
	m  map[string]int
}

func NewAsyncMap() *asyncMap {
	return &asyncMap{
		m: make(map[string]int),
	}
}

func (a *asyncMap) Get(key string) (int, bool) {
	a.mx.Lock()
	defer a.mx.Unlock()
	val, ok := a.m[key]
	return val, ok
}

func (a *asyncMap) Set(key string, value int) {
	a.mx.Lock()
	defer a.mx.Unlock()
	a.m[key] = value
}

func main() {
	asm := NewAsyncMap()
	asm.Set("first", 1)

	if el, ok := asm.Get("first"); ok {
		fmt.Println(el)
	}
	wg := new(sync.WaitGroup)
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			asm.Set(strconv.Itoa(i), i)
			wg.Done()
		}(i)
	}
	wg.Wait()

	fmt.Println(asm)
}
