// Разработать программу, которая будет последовательно отправлять значения в канал, а с
// другой стороны канала — читать. По истечению N секунд программа должна завершаться.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func getSec() int {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter duration: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSuffix(input, "\n")
	duration, _ := strconv.Atoi(input)
	return duration
}

func getCh(msg chan int) {
	for number := range msg {
		fmt.Printf("New number: %d\n\n", number)
	}
}
func main() {
	seconds := getSec()

	numbersCh := make(chan int)
	defer close(numbersCh)

	start := time.Now()
	stop := time.Duration(seconds) * time.Second

	go getCh(numbersCh)

	for time.Since(start) < stop {
		numbersCh <- 2
		time.Sleep(time.Second)

	}
}
