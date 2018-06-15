package encoder_test

import (
	"testing"

	"github.com/adwd/gopherDojo/ghc/encoder"
)

func TestSelectEncoder(t *testing.T) {
	table := []struct {
		format string
	}{
		{
			format: "jpeg",
		},
		{
			format: "jpg",
		},
		{
			format: "png",
		},
		{
			format: "gif",
		},
	}

	for _, format := range table {
		_, err := encoder.SelectEncoder(format.format)
		if err != nil {
			t.Errorf("format: %s", format.format)
		}
	}
}
