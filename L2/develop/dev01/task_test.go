package ntp_time

import (
	"testing"
	"time"
)

func TestGetCurrentTime(t *testing.T) {
	var tests = []struct {
		name string
		want time.Time
	}{
		{"first test", time.Now()},
	}

	for _, tt := range tests {

		testname := tt.name
		t.Run(testname, func(t *testing.T) {
			answer_time := GetCurretTime()
			if answer_time.Year() != tt.want.Year() || answer_time.Month() != tt.want.Month() || answer_time.Day() != tt.want.Day() {
				t.Errorf("got %s, want %s", answer_time, tt.want)
			}
		})
	}
}
