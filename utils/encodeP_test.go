package utils_test

import (
	"testing"

	"github.com/phaus/platinum/utils"
)

func TestSimpleEncodeP(t *testing.T) {
	input := "Bob->Alice : hello"
	expected := "SyfFKj2rKt3CoKnELR1Io4ZDoSa70000"

	output := utils.EncodeP(input)

	if output != expected {
		t.Errorf("%s != %s", output, expected)
	}
}
