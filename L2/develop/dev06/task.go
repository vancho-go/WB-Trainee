package cut

import (
	"io/ioutil"
	"log"
	"strings"
)

/*
=== Утилита cut ===
Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные
Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем
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

func Cut(f *int, d *string, s *bool, filename string) []string {
	// f := flag.Int("f", 0, "'fields' - выбрать поля (колонки)")
	// d := flag.String("d", "\t", "'delimiter' - использовать другой разделитель")
	// s := flag.Bool("s", false, "'separated' - только строки с разделителем")
	// flag.Parse()
	// filename := flag.Arg(len(flag.Args()) - 1)
	originalText := readFile(filename)
	var midResult [][]string
	for _, line := range originalText {
		curLine := strings.Split(line, *d)
		if *s {
			if len(curLine) > 1 {
				midResult = append(midResult, curLine)
				continue
			} else {
				continue
			}
		}
		midResult = append(midResult, curLine)
	}

	var result []string
	for _, line := range midResult {
		if *f != 0 {
			if *f-1 <= len(line)-1 {
				// fmt.Println(line[*f-1])
				result = append(result, line[*f-1])
			}
		} else {
			// fmt.Println(line)
			result = append(result, line...)
		}
	}
	return result
}
