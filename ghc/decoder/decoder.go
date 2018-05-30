package decoder

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
)

type Decoder = func(r io.Reader) (image.Image, error)

var decoders = map[string]Decoder{
	"jpeg": jpeg.Decode,
	"jpg":  jpeg.Decode,
	"png":  png.Decode,
	"gif":  gif.Decode,
}

func SelectDecoder(format string) (Decoder, error) {
	decoder, ok := decoders[format]
	if !ok {
		return nil, fmt.Errorf("invalid input format")
	}

	return decoder, nil
}
