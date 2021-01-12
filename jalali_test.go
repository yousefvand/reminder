package main

import (
	"fmt"
	"testing"
)

func TestJalaliToGregorian(t *testing.T) {
	var tests = []struct {
		a, b, c             int
		wantY, wantM, wantD int
	}{
		{1372, 8, 11, 1993, 11, 2},
		{1355, 2, 28, 1976, 5, 18},
		{1395, 12, 30, 2017, 3, 20},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%d-%d-%d", tt.a, tt.b, tt.c)
		t.Run(testname, func(t *testing.T) {
			ansY, ansM, ansD := JalaliToGregorian(tt.a, tt.b, tt.c)
			if ansY != tt.wantY {
				t.Errorf("got %d, want %d", ansY, tt.wantY)
			}
			if ansM != tt.wantM {
				t.Errorf("got %d, want %d", ansM, tt.wantM)
			}
			if ansD != tt.wantD {
				t.Errorf("got %d, want %d", ansD, tt.wantD)
			}
		})
	}
}
