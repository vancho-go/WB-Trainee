package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strings"
)

/*
=== Утилита sort ===
Отсортировать строки (man sort)
Основное
Поддержать ключи
-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки
Дополнительное
Поддержать ключи
-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// var k int
	// var n, r, u bool
	// k := flag.Int("k", 0, "указание колонки для сортировки")
	// n := flag.Bool("n", false, "сортировать по числовому значению")
	r := flag.Bool("r", false, "сортировать в обратном порядке")
	// u := flag.Bool("u", false, "не выводить повторяющиеся строки")
	flag.Parse()
	filename := flag.Arg(0)

	data, err := os.Open(filename)
	defer data.Close()
	check(err)

	strs := [][]string{}
	scanner := bufio.NewScanner(data)

	for scanner.Scan() {
		strs = append(strs, strings.Split(scanner.Text(), " "))
	}

	if *r {
		sort.Slice(strs, func(i, j int) bool {
			return strs[0][i] < strs[0][j]
		})
	}

	fmt.Println(strs)

	// fmt.Println("file:", filename)
	// fmt.Println("k:", *k)
	// fmt.Println("n:", *n)
	// fmt.Println("r:", *r)
	// fmt.Println("u:", *u)
}
