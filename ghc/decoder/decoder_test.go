package decoder_test

import (
	"testing"

	"github.com/adwd/gopherDojo/ghc/decoder"
)

func TestSelectDecoder(t *testing.T) {
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
		_, err := decoder.SelectDecoder(format.format)
		if err != nil {
			t.Errorf("no error expected, error = %v", err)
		}
	}

	table = []struct {
		format string
	}{
		{
			format: "jpegg",
		},
		{
			format: "go",
		},
		{
			format: "txt",
		},
	}

	for _, format := range table {
		_, err := decoder.SelectDecoder(format.format)
		if err == nil {
			t.Errorf("format: %v should return error", format.format)
		}
	}
}
