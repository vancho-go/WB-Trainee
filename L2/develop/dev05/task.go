package grep

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

/*
=== Утилита grep ===
Реализовать утилиту фильтрации (man grep)
Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "conoriginalText" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func readFile(filename string) []string {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalln(err)
	}
	originalText := strings.Split(string(file), "\n")
	return originalText
}

// func printResult(n *bool, text []string, start, stop int) {
// 	if *n {
// 		for i := start; i < stop; i++ {
// 			fmt.Printf("%d:%s\n", i+1, text[i])
// 		}
// 	} else {
// 		for i := start; i < stop; i++ {
// 			fmt.Printf("%s\n", text[i])
// 		}
// 	}
// }

func makeResultSlice(n *bool, text []string, start, stop int) []string {
	var result []string
	if *n {
		for i := start; i < stop; i++ {
			result = append(result, fmt.Sprintf("%d:%s", i+1, text[i]))
		}
	} else {
		for i := start; i < stop; i++ {
			result = append(result, text[i])
		}
	}
	return result
}

func Grep(A, B, C *int, c, i, v, F, n *bool, filename, searchPhrase string) []string {
	// A := flag.Int("A", 0, "'after' печатать +N строк после совпадения")
	// B := flag.Int("B", 0, "'before' печатать +N строк до совпадения")
	// C := flag.Int("C", 0, "'context' (A+B) печатать ±N строк вокруг совпадения")
	// c := flag.Bool("c", false, "'count' (количество строк)")
	// i := flag.Bool("i", false, "'ignore-case' (игнорировать регистр)")
	// v := flag.Bool("v", false, "'invert' (вместо совпадения, исключать)")
	// F := flag.Bool("F", false, "'fixed', точное совпадение со строкой, не паттерн")
	// n := flag.Bool("n", false, "'line num', печатать номер строки")
	// flag.Parse()
	// filename := flag.Arg(len(flag.Args()) - 1)
	// searchPhrase := flag.Arg(len(flag.Args()) - 2)
	originalText := readFile(filename)

	linesNumber := len(originalText) - 1
	var result []string
	var start, stop int
	var searchText []string
	if *i {
		searchPhrase = strings.ToLower(searchPhrase)
		for i := 0; i < len(originalText); i++ {
			searchText = append(searchText, strings.ToLower(originalText[i]))
		}
	} else {
		searchText = originalText
	}

	for pos, line := range searchText {
		if *F {
			if line != searchPhrase {
				continue
			}
		} else {
			if !strings.Contains(line, searchPhrase) {
				continue
			}
		}

		start = pos
		stop = pos
		if *c {
			result = append(result, strconv.Itoa(len(originalText)))
			// fmt.Println(len(originalText))
			break
		}

		if *A != 0 {
			if pos+*A > linesNumber {
				stop = linesNumber
			} else {
				stop = pos + *A + 1
			}
		}

		if *B != 0 {
			if pos-*B < 0 {
				start = 0
			} else {
				start = pos - *B
			}
			if *A == 0 {
				stop = pos + 1
			}
		}

		if *C != 0 {
			if pos-*C < 0 {
				start = 0
			} else {
				start = pos - *C
			}

			if pos+*C > linesNumber {
				stop = linesNumber + 1
			} else {
				stop = pos + *C + 1
			}
		}

		if *A == 0 && *B == 0 && *C == 0 {
			stop += 1
		}

		if *v {
			// fmt.Println(originalText[:start], originalText[stop:])

			// printResult(n, originalText, 0, start)
			// printResult(n, originalText, stop, len(originalText))
			result = makeResultSlice(n, originalText, 0, start)
			result = append(result, makeResultSlice(n, originalText, stop, len(originalText))...)
			break
		}

		// fmt.Println(originalText[start : stop+1])
		// printResult(n, originalText, start, stop)
		// fmt.Println(start, stop)
		result = makeResultSlice(n, originalText, start, stop)
		break

	}
	return result
}
