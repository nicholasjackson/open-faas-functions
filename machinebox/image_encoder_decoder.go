package function

import (
	"bytes"
	"image"
	"image/draw"
	_ "image/jpeg"
	png "image/png"
	"io"
)

// ImageEncoderDecoderImpl implements ImageEncoderDecoder
type ImageEncoderDecoderImpl struct{}

// Decode a stream into a draw.Image
func (i *ImageEncoderDecoderImpl) Decode(r io.Reader) (draw.Image, error) {
	src, _, err := image.Decode(r)
	if err != nil {
		return nil, err
	}

	b := src.Bounds()
	img := image.NewRGBA(b)
	draw.Draw(img, b, src, b.Min, draw.Src)

	return img, nil
}

// Encode a image into png format and write to a buffer
func (i *ImageEncoderDecoderImpl) Encode(b *bytes.Buffer, img draw.Image) {
	png.Encode(b, img)
}
