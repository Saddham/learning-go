package models_test

import (
	"learning-go/models"
	"testing"
)

func TestCity(t *testing.T) {
	expectedName := "Test City"
	temperature := 30
	city := models.NewCity(expectedName, float64(temperature), true, true)

	t.Run("name", func(t *testing.T) {
		got := city.Name()

		if got != expectedName {
			t.Errorf("Expected '%v', but got '%v'", expectedName, got)
		}
	})

	t.Run("tempC", func(t *testing.T) {
		// t.Helper() => Marks this function as non test func

		got := city.TempC()

		if got != float64(temperature) {
			t.Errorf("Expected '%v', but got '%v'", temperature, got)
		}
	})
}
