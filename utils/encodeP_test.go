package utils_test

import (
	"testing"

	"github.com/phaus/platinum/utils"
)

func TestSimpleEncodeP(t *testing.T) {
	input := "ABCDEFGHIJ"
	expected := "StHoTd5rS_Vmz080"

	output := utils.EncodeP(input)

	if output != expected {
		t.Errorf("%s != %s", output, expected)
	}
}
