package cut

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

func createString(x string) *string {
	return &x
}

func TestCut(t *testing.T) {
	var tests = []struct {
		name     string
		f        *int
		d        *string
		s        *bool
		filename string
		want     []string
	}{
		{
			name:     "test1-f-flag",
			f:        createInt(1),
			d:        createString("\t"),
			s:        createBool(false),
			filename: "test.txt",
			want: []string{
				"Winter: white: snow: frost",
				"Spring: green: grass: warm",
				"Summer: colorful: blossom: hot",
				"Autumn: yellow: leaves: cool",
				"Hello Test"},
		},
		{
			name:     "test2-fs-flag",
			f:        createInt(2),
			d:        createString(":"),
			s:        createBool(false),
			filename: "test.txt",
			want: []string{
				" white",
				" green",
				" colorful",
				" yellow"},
		},
	}

	for _, tt := range tests {
		testname := tt.name
		t.Run(testname, func(t *testing.T) {
			answer := Cut(tt.f, tt.d, tt.s, tt.filename)
			if !reflect.DeepEqual(answer, tt.want) {
				t.Errorf("got %s, want %s", answer, tt.want)
			}
		})
	}
}
