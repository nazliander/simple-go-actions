package predictor

import (
	"testing"

	"github.com/goccy/go-reflect"
)

func TestSplitParser(t *testing.T) {

	table := []struct {
		name     string
		input    string
		expected []string
	}{
		{"same number", "2,2, 2", []string{"2", "2", "2"}},
		{"zero", "-1,0,1", []string{"-1", "0", "1"}},
	}

	for _, tc := range table {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := SplitParser(tc.input)
			if !reflect.DeepEqual(tc.expected, actual) || (err != nil) {
				t.Errorf("Expected %s, got %s", tc.expected, actual)
			}
		})
	}
}

func TestMovingAverage(t *testing.T) {

	table := []struct {
		name     string
		input    []string
		expected float64
	}{
		{"basic test", []string{"2", "4", "6"}, 4.0},
		{"minus", []string{"-2.34", "-3.16", "-6.50"}, -4.0},
	}

	for _, tc := range table {
		t.Run(tc.name, func(t *testing.T) {
			actual := MovingAverage(tc.input)
			if !reflect.DeepEqual(tc.expected, actual) {
				t.Errorf("Expected %f, got %f", tc.expected, actual)
			}
		})
	}
}
