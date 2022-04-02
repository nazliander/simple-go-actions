package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setupSuite(tb testing.TB) func(tb testing.TB) {
	log.Println("----- Setup suite")
	return func(tb testing.TB) {
		log.Println("----- Teardown suite")
	}
}

func setupTest(tb testing.TB, n string) (func(tb testing.TB), *httptest.ResponseRecorder) {
	log.Println("00 Setup Forecast Test")

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		tb.Fatal(err)
	}
	q := req.URL.Query()
	q.Add("n", fmt.Sprint(n))
	req.URL.RawQuery = q.Encode()
	nr := httptest.NewRecorder()
	h := http.HandlerFunc(forecast)
	h.ServeHTTP(nr, req)

	return func(tb testing.TB) {
		log.Println("01 Teardown Forecast Test")
	}, nr
}

func TestForecast(t *testing.T) {
	teardownSuite := setupSuite(t)
	defer teardownSuite(t)

	const epsilon = 0.00001

	table := []struct {
		name     string
		input    string
		expected float64
	}{
		{"same number", "2, 2, 2", 2.0},
		{"minus number", "-1, -2, -3", -2.0},
		{"zero", "-1,0,1", 0.0},
		{"big float", "0, 1, 1", 0.66666},
		{"big number", "202,78,9999", 3426.33333},
	}

	for _, tc := range table {
		t.Run(tc.name, func(t *testing.T) {
			teardownTest, rr := setupTest(t, tc.input)
			defer teardownTest(t)

			if status := rr.Code; status != http.StatusOK {
				t.Errorf("Handler returned wrong status code: got %v want %v",
					status, http.StatusOK)
			}

			actual := Response{}
			json.Unmarshal([]byte(rr.Body.String()), &actual)
			if math.Abs(tc.expected-actual.Result) > epsilon {
				t.Errorf("Expected %f, got %f", tc.expected, actual.Result)
			}
		})
	}
}
