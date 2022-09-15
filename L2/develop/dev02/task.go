package rle

import (
	"strings"
)

/*
=== Задача на распаковку ===
Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)
В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.
Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func checkIsNumber(s rune) bool {
	// the integer value out of the ASCII (or UTF-8) representation of the digit
	if s >= '0' && s <= '9' {
		return true
	} else {
		return false
	}
}
func Extract(input string) string {
	inputRune := []rune(input)
	var result []rune

	for len(inputRune) > 0 {
		if len(inputRune) == 1 {
			result = append(result, inputRune[0])
			break
		}
		let, num := inputRune[0], inputRune[1]
		if checkIsNumber(let) && checkIsNumber(num) {
			return "некорректная строка"
		}
		if checkIsNumber(num) {
			result = append(result, ([]rune(strings.Repeat(string(let), int(num-'0'))))...)
			inputRune = inputRune[2:]
		} else {
			result = append(result, let)
			inputRune = inputRune[1:]
		}
	}
	return string(result)
}
