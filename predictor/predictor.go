package predictor

import (
	"errors"
	"strconv"
	"strings"
)

func SplitParser(s string) ([]string, error) {

	splittedStrNrs := strings.Split(s, ",")
	if len(splittedStrNrs) != 3 {
		var emptyArr []string
		return emptyArr, errors.New("Parser error")
	}
	for i := range splittedStrNrs {
		splittedStrNrs[i] = strings.TrimSpace(splittedStrNrs[i])
	}
	return splittedStrNrs, nil

}

func MovingAverage(splittedStrNrs []string) float64 {

	sum := 0.0
	for _, x := range splittedStrNrs {
		parsedFloat, err := strconv.ParseFloat(x, 64)
		if err != nil {
			return 0
		}
		sum = sum + parsedFloat
	}
	return sum / 3.0

}
