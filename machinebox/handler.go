package function

import (
	"bytes"
	"fmt"
	"image/color"
	"image/draw"
	"io"

	"github.com/llgcode/draw2d"
	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/machinebox/sdk-go/facebox"
)

type deps func() *facebox.Client

//go:generate moq -out facebox_moq_test.go . FaceboxChecker

// FaceboxChecker defines checkiung facebox faces
type FaceboxChecker interface {
	Check(image io.Reader) ([]facebox.Face, error)
}

//go:generate moq -out drawable_moq_test.go . Drawable

// Drawable defines and interface for a drawing canvas
type Drawable interface {
	SetStrokeColor(color.Color)
	SetFillColor(color.Color)
	SetLineWidth(float64)
	MoveTo(x, y float64)
	LineTo(x, y float64)
	Close()
	FillStroke(paths ...*draw2d.Path)
}

//go:generate moq -out encoder_moq_test.go . ImageEncoderDecoder

// ImageEncoderDecoder allows the encoding and decoding of images
type ImageEncoderDecoder interface {
	Decode(img io.Reader) (draw.Image, error)
	Encode(b *bytes.Buffer, img draw.Image)
}

type graphicContextCreator func(img draw.Image) Drawable

var dependencies = func() (FaceboxChecker, graphicContextCreator, ImageEncoderDecoder) {
	return facebox.New("http://192.168.192.131:8080"),
		func(img draw.Image) Drawable {
			return draw2dimg.NewGraphicContext(img)
		},
		&ImageEncoderDecoderImpl{}
}

const (
	getFacesError    = "Error getting faces %s"
	decodeImageError = "Error decoding image %s"
)

// Handle a serverless request
func Handle(req []byte) string {
	c, newGraphicsContext, imgEncoderDecoder := dependencies()

	faces, err := c.Check(bytes.NewReader(req))
	if err != nil {
		return fmt.Sprintf(getFacesError, err)
	}

	img, err := imgEncoderDecoder.Decode(bytes.NewReader(req))
	if err != nil {
		return fmt.Sprintf(decodeImageError, err)
	}

	gc := newGraphicsContext(img)
	gc.SetStrokeColor(color.Black)
	gc.SetFillColor(color.Transparent)
	gc.SetLineWidth(2.0)

	for _, f := range faces {
		gc.MoveTo(float64(f.Rect.Left), float64(f.Rect.Top))
		gc.LineTo(float64(f.Rect.Left+f.Rect.Width), float64(f.Rect.Top))
		gc.LineTo(float64(f.Rect.Left+f.Rect.Width), float64(f.Rect.Top+f.Rect.Height))
		gc.LineTo(float64(f.Rect.Left), float64(f.Rect.Top+f.Rect.Height))
		gc.LineTo(float64(f.Rect.Left), float64(f.Rect.Top))
		gc.Close()
		gc.FillStroke()
	}

	buf := &bytes.Buffer{}
	imgEncoderDecoder.Encode(buf, img)

	return string(buf.Bytes())
}
