package utils_test

import (
	"testing"

	"github.com/phaus/platinum/utils"
)

func TestSimpleEncodeP(t *testing.T) {
	input := "fooobar\ntest==><\n"
	expected := "IylFpqzABE8gIIqkiRMri-420000"

	output := utils.EncodeP(input)

	if output != expected {
		t.Errorf("%s != %s", output, expected)
	}
}
