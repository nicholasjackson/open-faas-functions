package function

import (
	"bytes"
	"os"
	"testing"

	"github.com/matryer/is"
)

func TestDecodeWithInvalidImageReturnsError(t *testing.T) {
	i := ImageEncoderDecoderImpl{}
	is := is.New(t)

	_, err := i.Decode(bytes.NewReader([]byte("s")))

	is.True(err != nil) // should not have returned an error
}

func TestDecodeWithValidImageReturnsImage(t *testing.T) {
	i := ImageEncoderDecoderImpl{}
	is := is.New(t)

	f, err := os.Open("./hashi.jpg")
	is.NoErr(err)

	img, err := i.Decode(f)

	is.True(img != nil) // should have returned an image
}

func TestEncodeWithValidImage(t *testing.T) {
	buf := &bytes.Buffer{}

	i := ImageEncoderDecoderImpl{}
	is := is.New(t)

	f, err := os.Open("./hashi.jpg")
	is.NoErr(err)

	img, err := i.Decode(f)
	is.NoErr(err)

	i.Encode(buf, img)

	is.True(len(buf.Bytes()) > 1) // should have encoded an image
}
