package hello_test

import (
	"go-action-test/hello"
	"testing"
)

func TestGreet(t *testing.T) {
	result := hello.Greet()
	if result != "Hello GitHub Actions" {
		t.Errorf("Greet() = %s; Expected Hello GitHub actions", result)
	}

}
