package anagrams

import (
	"sort"
	"strings"
)

/*
=== Поиск анаграмм по словарю ===
Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.
Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func sortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func sortSlice(data []string) []string {
	sort.Slice(data, func(i, j int) bool {
		return data[i] <= data[j]
	})
	return data
}

func FindAnagrams(input *[]string) map[string][]string {
	// Map вида [гбав] абвг
	wordToItsSortedStr := make(map[string][]string)
	for _, word := range *input {
		// В нижний регистр
		word = strings.ToLower(word)
		// Проверяем есть ли уже такой ключ в map
		if _, ok := wordToItsSortedStr[sortString(word)]; ok {
			wordToItsSortedStr[sortString(word)] = append(wordToItsSortedStr[sortString(word)], word)
			// Если ключа еще нет, добавляем
		} else {
			wordToItsSortedStr[sortString(word)] = []string{word}
		}
	}

	result := make(map[string][]string)

	for _, v := range wordToItsSortedStr {
		if len(v) < 2 {
			continue
		} else {
			result[v[0]] = sortSlice(v[1:])
		}
	}
	return result
}
