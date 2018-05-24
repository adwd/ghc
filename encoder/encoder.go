package encoder

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
)

type Encoder = func(w io.Writer, img image.Image) error

func SelectEncoder(format string) (Encoder, error) {
	switch format {
	case "jpeg", "jpg":
		return func(w io.Writer, img image.Image) error {
			return jpeg.Encode(w, img, nil)
		}, nil
	case "png":
		return png.Encode, nil
	case "gif":
		return func(w io.Writer, img image.Image) error {
			return gif.Encode(w, img, nil)
		}, nil
	default:
		return nil, fmt.Errorf("invalid output format")
	}

}
