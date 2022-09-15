package anagrams

import (
	"reflect"
	"testing"
)

func TestFindAnagrams(t *testing.T) {
	var tests = []struct {
		name  string
		input *[]string
		want  map[string][]string
	}{
		{
			name:  "test1-empty",
			input: &[]string{""},
			want:  map[string][]string{},
		},
		{
			name:  "test2-one_word",
			input: &[]string{"пятка"},
			want:  map[string][]string{},
		},
		{
			name:  "test3-two_words",
			input: &[]string{"пятка", "пятак"},
			want: map[string][]string{
				"пятка": {"пятак"}},
		},
		{
			name:  "test4-many_words",
			input: &[]string{"пяТак", "листок", "пятка", "слиток", "тяпка", "Столик", "ничего", "акптя"},
			want: map[string][]string{
				"листок": {"слиток", "столик"},
				"пятак":  {"акптя", "пятка", "тяпка"}},
		},
	}

	for _, tt := range tests {
		testname := tt.name
		t.Run(testname, func(t *testing.T) {
			answer := FindAnagrams(tt.input)
			if !reflect.DeepEqual(answer, tt.want) {
				t.Errorf("got %s, want %s", answer, tt.want)
			}
		})
	}
}
