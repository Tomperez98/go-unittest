package server

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

type Test struct {
	Name          string
	Server        *httptest.Server
	Response      *Weather
	ExpectedError error
}

func TestGetWeahter(t *testing.T) {

	tests := []Test{
		{Name: "basic-request", Server: httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{ "city" : "Denver, CO", "forecast" : "sunny" }`))
		})), Response: &Weather{City: "Denver, CO", Forecast: "Sunny"}, ExpectedError: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			defer test.Server.Close()

			resp, _ := GetWeather(test.Server.URL)
			if !reflect.DeepEqual(resp, test.ExpectedError) {
				t.Errorf("Expected error FAILED: expected %v got %v\n", test.ExpectedError, resp)
			}
		})
	}

}
