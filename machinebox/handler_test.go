package function

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"io"
	"testing"

	"github.com/llgcode/draw2d"
	"github.com/machinebox/sdk-go/facebox"
	"github.com/matryer/is"
)

/*
func TestFunctionalTest(t *testing.T) {
	f, err := os.Open("./hashi.jpg")
	if err != nil {
		t.Fail()
	}

	data, err := ioutil.ReadAll(f)
	if err != nil {
		t.Fail()
	}

	out := Handle(data)
	fmt.Println(out)
}
*/

func setup(t *testing.T, faces []facebox.Face, facesError error, imgError error) (*FaceboxCheckerMock, *DrawableMock, *ImageEncoderDecoderMock, *is.I) {
	mockedFaceboxChecker := &FaceboxCheckerMock{
		CheckFunc: func(image io.Reader) ([]facebox.Face, error) {
			return faces, facesError
		},
	}

	mockedDrawable := &DrawableMock{
		CloseFunc: func() {
		},
		FillStrokeFunc: func(paths ...*draw2d.Path) {
		},
		LineToFunc: func(x float64, y float64) {
		},
		MoveToFunc: func(x float64, y float64) {
		},
		SetFillColorFunc: func(in1 color.Color) {
		},
		SetLineWidthFunc: func(in1 float64) {
		},
		SetStrokeColorFunc: func(in1 color.Color) {
		},
	}

	mockedImageEncoderDecoder := &ImageEncoderDecoderMock{
		DecodeFunc: func(img io.Reader) (draw.Image, error) {
			return image.NewNRGBA(image.Rect(0, 0, 0, 0)), imgError
		},
		EncodeFunc: func(b *bytes.Buffer, img draw.Image) {
		},
	}

	dependencies = func() (FaceboxChecker, graphicContextCreator, ImageEncoderDecoder) {
		return mockedFaceboxChecker,
			func(img draw.Image) Drawable {
				return mockedDrawable
			},
			mockedImageEncoderDecoder
	}

	return mockedFaceboxChecker, mockedDrawable, mockedImageEncoderDecoder, is.New(t)
}

func TestCallFaceboxCheck(t *testing.T) {
	m, _, _, is := setup(t, nil, nil, nil)

	Handle([]byte(""))

	is.Equal(1, len(m.CheckCalls())) // should have called check
}

func TestFaceboxCheckWithErrorReturnsError(t *testing.T) {
	err := fmt.Errorf("boom")
	_, _, _, is := setup(t, nil, err, nil)

	out := Handle([]byte(""))

	is.Equal(fmt.Sprintf(getFacesError, err), out) // output error should have been returned by handler
}

func TestWhenImageDecodeReturnsErrorReturnsError(t *testing.T) {
	_, _, _, is := setup(t, nil, nil, image.ErrFormat)

	out := Handle([]byte(""))

	is.Equal(fmt.Sprintf(decodeImageError, image.ErrFormat), out) // output error should have been returned by handler
}

func TestSetsUpGC(t *testing.T) {
	_, gc, _, is := setup(t, nil, nil, nil)

	Handle([]byte(""))

	is.Equal(1, len(gc.SetStrokeColorCalls())) // should call set stroke color
	is.Equal(1, len(gc.SetFillColorCalls()))   // should call set fill color
	is.Equal(1, len(gc.SetLineWidthCalls()))   // should call set fill color
}

func TestDrawsFaceRectangle(t *testing.T) {
	faces := []facebox.Face{
		facebox.Face{
			Rect: facebox.Rect{
				Left:   10,
				Top:    20,
				Width:  100,
				Height: 200,
			},
		},
	}
	_, gc, _, is := setup(t, faces, nil, nil)

	Handle([]byte(""))

	is.Equal(1, len(gc.MoveToCalls()))                           // should have called move to
	is.Equal(float64(faces[0].Rect.Left), gc.MoveToCalls()[0].X) // should have set x to 10
	is.Equal(float64(faces[0].Rect.Top), gc.MoveToCalls()[0].Y)  // should have set y to 20

	is.Equal(4, len(gc.LineToCalls()))            // should have called line to 4 times
	is.Equal(float64(110), gc.LineToCalls()[0].X) // should have set x to 110
	is.Equal(float64(20), gc.LineToCalls()[0].Y)  // should have set y to 20
	is.Equal(float64(110), gc.LineToCalls()[1].X) // should have set x to 110
	is.Equal(float64(220), gc.LineToCalls()[1].Y) // should have set y to 20
	is.Equal(float64(10), gc.LineToCalls()[2].X)  // should have set x to 10
	is.Equal(float64(220), gc.LineToCalls()[2].Y) // should have set y to 20
	is.Equal(float64(10), gc.LineToCalls()[3].X)  // should have set x to 110
	is.Equal(float64(20), gc.LineToCalls()[3].Y)  // should have set y to 20

	is.Equal(1, len(gc.CloseCalls()))      // should have called close once
	is.Equal(1, len(gc.FillStrokeCalls())) // should have called fill stroke once
	// complete tests
}
