package function

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	_ "image/jpeg"
	png "image/png"

	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/machinebox/sdk-go/facebox"
)

// Handle a serverless request
func Handle(req []byte) string {
	c := facebox.New("http://192.168.192.131:8080")
	faces, err := c.Check(bytes.NewReader(req))
	if err != nil {
		return "boom"
	}

	img, _, err := image.Decode(bytes.NewReader(req))
	if err != nil {
		return fmt.Sprintf("unable to decode image, %s", err)
	}

	b := img.Bounds()
	m := image.NewRGBA(b)
	draw.Draw(m, b, img, b.Min, draw.Src)

	gc := draw2dimg.NewGraphicContext(m)
	gc.SetStrokeColor(color.Black)
	gc.SetFillColor(color.Transparent)
	gc.SetLineWidth(2.0)

	for _, f := range faces {
		gc.BeginPath()
		gc.MoveTo(float64(f.Rect.Left), float64(f.Rect.Top))
		gc.LineTo(float64(f.Rect.Left+f.Rect.Width), float64(f.Rect.Top))
		gc.LineTo(float64(f.Rect.Left+f.Rect.Width), float64(f.Rect.Top+f.Rect.Height))
		gc.LineTo(float64(f.Rect.Left), float64(f.Rect.Top+f.Rect.Height))
		gc.LineTo(float64(f.Rect.Left), float64(f.Rect.Top))
		gc.Close()
		gc.FillStroke()
	}

	buf := &bytes.Buffer{}
	err = png.Encode(buf, m)
	if err != nil {
		return fmt.Sprintf("Error encoding image %s", err)
	}
	//draw2dimg.SaveToPngFile("/tmp/stuff.png", m)

	return string(buf.Bytes())
}
