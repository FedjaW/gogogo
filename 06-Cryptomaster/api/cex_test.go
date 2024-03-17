package api_test

import (
	"testing"

	"leckomio.dev/go/crypto/api"
)

func TestAPICall(t *testing.T) {
	// arrange
	invalidInput := ""

	// act
	_, err := api.GetRate(invalidInput)

	// assert
	if err == nil {
		t.Error("Error was not found")
	}
}
