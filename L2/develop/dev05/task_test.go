package grep

import (
	"reflect"
	"testing"
)

func createInt(x int) *int {
	return &x
}

func createBool(x bool) *bool {
	return &x
}

func TestGrep(t *testing.T) {
	var tests = []struct {
		name                   string
		A, B, C                *int
		c, i, v, F, n          *bool
		filename, searchPhrase string
		want                   []string
	}{
		{
			name:         "test1-A-flag",
			A:            createInt(1),
			B:            createInt(0),
			C:            createInt(0),
			c:            createBool(false),
			i:            createBool(false),
			v:            createBool(false),
			F:            createBool(false),
			n:            createBool(false),
			filename:     "test.txt",
			searchPhrase: "CoUnt",
			want: []string{
				"-C - 'context' (A+B) печатать ±N строк вокруг совпадения -c - 'CoUnt' (количество строк)",
				"-i - 'ignore-case' (игнорировать регистр)"},
		},
		{
			name:         "test2-B-flag",
			A:            createInt(0),
			B:            createInt(1),
			C:            createInt(0),
			c:            createBool(false),
			i:            createBool(false),
			v:            createBool(false),
			F:            createBool(false),
			n:            createBool(false),
			filename:     "test.txt",
			searchPhrase: "CoUnt",
			want: []string{"-B - 'before' печатать +N строк до совпадения",
				"-C - 'context' (A+B) печатать ±N строк вокруг совпадения -c - 'CoUnt' (количество строк)",
			},
		},
		{
			name:         "test3-AB-flag",
			A:            createInt(1),
			B:            createInt(1),
			C:            createInt(0),
			c:            createBool(false),
			i:            createBool(false),
			v:            createBool(false),
			F:            createBool(false),
			n:            createBool(false),
			filename:     "test.txt",
			searchPhrase: "CoUnt",
			want: []string{"-B - 'before' печатать +N строк до совпадения",
				"-C - 'context' (A+B) печатать ±N строк вокруг совпадения -c - 'CoUnt' (количество строк)",
				"-i - 'ignore-case' (игнорировать регистр)",
			},
		},
		{
			name:         "test4-C-flag",
			A:            createInt(0),
			B:            createInt(0),
			C:            createInt(1),
			c:            createBool(false),
			i:            createBool(false),
			v:            createBool(false),
			F:            createBool(false),
			n:            createBool(false),
			filename:     "test.txt",
			searchPhrase: "CoUnt",
			want: []string{"-B - 'before' печатать +N строк до совпадения",
				"-C - 'context' (A+B) печатать ±N строк вокруг совпадения -c - 'CoUnt' (количество строк)",
				"-i - 'ignore-case' (игнорировать регистр)",
			},
		},
		{
			name:         "test5-Ac-flag",
			A:            createInt(1),
			B:            createInt(0),
			C:            createInt(0),
			c:            createBool(true),
			i:            createBool(false),
			v:            createBool(false),
			F:            createBool(false),
			n:            createBool(false),
			filename:     "test.txt",
			searchPhrase: "CoUnt",
			want:         []string{"9"},
		},
		{
			name:         "test6-Ai-flag",
			A:            createInt(1),
			B:            createInt(0),
			C:            createInt(0),
			c:            createBool(false),
			i:            createBool(true),
			v:            createBool(false),
			F:            createBool(false),
			n:            createBool(false),
			filename:     "test.txt",
			searchPhrase: "count",
			want: []string{
				"-C - 'context' (A+B) печатать ±N строк вокруг совпадения -c - 'CoUnt' (количество строк)",
				"-i - 'ignore-case' (игнорировать регистр)"},
		},
		{
			name:         "test7-Av-flag",
			A:            createInt(1),
			B:            createInt(0),
			C:            createInt(0),
			c:            createBool(false),
			i:            createBool(false),
			v:            createBool(true),
			F:            createBool(false),
			n:            createBool(false),
			filename:     "test.txt",
			searchPhrase: "CoUnt",
			want: []string{
				"Утилита grep",
				"Реализовать утилиту фильтрации по аналогии с консольной утилитой (man grep — смотрим описание и основные параметры).",
				"Реализовать поддержку утилитой следующих ключей:",
				"-A - 'after' печатать +N строк после совпадения",
				"-B - 'before' печатать +N строк до совпадения",
				"-v - 'invert' (вместо совпадения, исключать)",
				"-F - 'fixed', точное совпадение со строкой, не паттерн -n - 'line num', напечатать номер строки"},
		},
		{
			name:         "test7-AF-flag",
			A:            createInt(1),
			B:            createInt(0),
			C:            createInt(0),
			c:            createBool(false),
			i:            createBool(false),
			v:            createBool(false),
			F:            createBool(true),
			n:            createBool(false),
			filename:     "test.txt",
			searchPhrase: "Утилита grep",
			want: []string{
				"Утилита grep",
				"Реализовать утилиту фильтрации по аналогии с консольной утилитой (man grep — смотрим описание и основные параметры).",
			},
		},
		{
			name:         "test8-An-flag",
			A:            createInt(1),
			B:            createInt(0),
			C:            createInt(0),
			c:            createBool(false),
			i:            createBool(false),
			v:            createBool(false),
			F:            createBool(false),
			n:            createBool(true),
			filename:     "test.txt",
			searchPhrase: "CoUnt",
			want: []string{
				"6:-C - 'context' (A+B) печатать ±N строк вокруг совпадения -c - 'CoUnt' (количество строк)",
				"7:-i - 'ignore-case' (игнорировать регистр)",
			},
		},
		{
			name:         "test9-no-flags-flag",
			A:            createInt(0),
			B:            createInt(0),
			C:            createInt(0),
			c:            createBool(false),
			i:            createBool(false),
			v:            createBool(false),
			F:            createBool(false),
			n:            createBool(false),
			filename:     "test.txt",
			searchPhrase: "CoUnt",
			want: []string{
				"-C - 'context' (A+B) печатать ±N строк вокруг совпадения -c - 'CoUnt' (количество строк)",
			},
		},
	}

	for _, tt := range tests {
		testname := tt.name
		t.Run(testname, func(t *testing.T) {
			answer := Grep(tt.A, tt.B, tt.C, tt.c, tt.i, tt.v, tt.F, tt.n, tt.filename, tt.searchPhrase)
			if !reflect.DeepEqual(answer, tt.want) {
				t.Errorf("got %s, want %s", answer, tt.want)
			}
		})
	}
}
