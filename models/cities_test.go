package models_test

import (
	"learning-go/data"
	"learning-go/mocks"
	"learning-go/models"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestFilter(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockReader := mocks.NewMockDataReader(ctrl)

	tests := []struct {
		testName  string
		responses []data.Response
		want      []data.Response
	}{
		{
			testName: "List all",
			responses: []data.Response{{
				Id:          "bcn",
				Name:        "Barcelona (Spain)",
				HasBeach:    true,
				HasMountain: false,
				TempC:       30.5,
			}, {
				Id:          "nyc",
				Name:        "New York (United States of America)",
				HasBeach:    false,
				HasMountain: true,
				TempC:       28.7,
			}},
			want: []data.Response{{
				Id:          "bcn",
				Name:        "Barcelona (Spain)",
				HasBeach:    true,
				HasMountain: false,
				TempC:       30.5,
			}},
		},
	}

	for _, tc := range tests {
		t.Run(tc.testName, func(t *testing.T) {
			mockReader.EXPECT().ReadData().Return(tc.responses, nil).Times(1)

			cities, err := models.NewCities(mockReader)

			if err != nil {
				t.Fatal("Error creating cities: ", err)
			}

			result := cities.Filter(true, false)

			if len(result) != len(tc.want) {
				t.Fatalf("Expected %v, but got %v for results length", len(tc.want), len(result))
			}

			for i, w := range tc.want {
				if result[i].Name() != w.Name {
					t.Errorf("results[%v]: expected %v, but got %v", i, w, result[i])
				}
			}
		})
	}
}
