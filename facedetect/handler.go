package function

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"image"
	"image/color"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"gocv.io/x/gocv"
)

// Request is base64 encoded image

// Response for the function
type Response struct {
	Faces  []image.Rectangle
	Bounds image.Rectangle
}

// Handle a serverless request
func Handle(req []byte) string {

	var data []byte

	data, err := base64.StdEncoding.DecodeString(string(req))
	if err != nil {
		data = req
	}

	typ := http.DetectContentType(data)
	if typ != "image/jpeg" && typ != "image/png" {
		return "Only jpeg or png images, either raw uncompressed bytes or base64 encoded are acceptable inputs, you uploaded: " + typ
	}

	tmpfile, err := ioutil.TempFile("/tmp", "image")
	if err != nil {
		log.Fatal(err)
	}

	defer os.Remove(tmpfile.Name()) // clean up

	io.Copy(tmpfile, bytes.NewBuffer(data))

	// fetch and process the query string
	var output string
	query, err := url.ParseQuery(os.Getenv("Http_Query"))
	if err == nil {
		output = query.Get("output")
	}

	faceProcessor := NewFaceProcessor()
	faces, bounds := faceProcessor.DetectFaces(tmpfile.Name())

	resp := Response{
		Faces:  faces,
		Bounds: bounds,
	}

	j, err := json.Marshal(resp)
	if err != nil {
		return fmt.Sprintf("Error encoding output: %s", err)
	}

	// do we need to create and output an image?
	if output == "image" {
		data, err := faceProcessor.DrawFaces(tmpfile.Name(), faces)
		if err != nil {
			return fmt.Sprint("Error creating image output: %s", err)
		}

		return string(data)
	}

	// return the coordinates
	return string(j)
}

// BySize allows sorting images by size
type BySize []image.Rectangle

func (s BySize) Len() int {
	return len(s)
}
func (s BySize) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s BySize) Less(i, j int) bool {
	return s[i].Size().X > s[j].Size().X && s[i].Size().Y > s[j].Size().Y
}

var yellow = color.RGBA{255, 255, 0, 0}

// FaceProcessor detects the position of a face from an input image
type FaceProcessor struct {
	faceclassifier  *gocv.CascadeClassifier
	eyeclassifier   *gocv.CascadeClassifier
	glassclassifier *gocv.CascadeClassifier
}

// NewFaceProcessor creates a new face processor loading any dependent settings
func NewFaceProcessor() *FaceProcessor {
	// load classifier to recognize faces
	classifier1 := gocv.NewCascadeClassifier()
	classifier1.Load("/home/app/cascades/haarcascade_frontalface_default.xml")

	classifier2 := gocv.NewCascadeClassifier()
	classifier2.Load("/home/app/cascades/haarcascade_eye.xml")

	classifier3 := gocv.NewCascadeClassifier()
	classifier3.Load("/home/app/cascades/haarcascade_eye_tree_eyeglasses.xml")

	return &FaceProcessor{
		faceclassifier:  &classifier1,
		eyeclassifier:   &classifier2,
		glassclassifier: &classifier3,
	}
}

// DetectFaces detects faces in the image and returns an array of rectangle
func (fp *FaceProcessor) DetectFaces(file string) (faces []image.Rectangle, bounds image.Rectangle) {
	img := gocv.IMRead(file, gocv.IMReadColor)
	defer img.Close()

	bds := image.Rectangle{Min: image.Point{}, Max: image.Point{X: img.Cols(), Y: img.Rows()}}
	//	gocv.CvtColor(img, img, gocv.ColorRGBToGray)
	//	gocv.Resize(img, img, image.Point{}, 0.6, 0.6, gocv.InterpolationArea)

	// detect faces
	tmpfaces := fp.faceclassifier.DetectMultiScaleWithParams(
		img, 1.03, 3, 0, image.Point{X: 10, Y: 10}, image.Point{X: 200, Y: 200},
	)

	fcs := make([]image.Rectangle, 0)

	if len(tmpfaces) > 0 {
		// draw a rectangle around each face on the original image
		for _, f := range tmpfaces {
			// detect eyes
			faceImage := img.Region(f)

			eyes := fp.eyeclassifier.DetectMultiScaleWithParams(
				faceImage, 1.03, 3, 0, image.Point{X: 0, Y: 0}, image.Point{X: 100, Y: 100},
			)

			glasses := fp.glassclassifier.DetectMultiScaleWithParams(
				faceImage, 1.03, 3, 0, image.Point{X: 0, Y: 0}, image.Point{X: 100, Y: 100},
			)

			if len(eyes) > 0 || len(glasses) > 0 {
				fcs = append(fcs, f)
			}
		}

		return fcs, bds
	}

	return nil, bds
}

// DrawFaces adds a rectangle to the given image with the face location
func (fp *FaceProcessor) DrawFaces(file string, faces []image.Rectangle) ([]byte, error) {
	img := gocv.IMRead(file, gocv.IMReadColor)
	defer img.Close()

	for _, r := range faces {
		gocv.Rectangle(img, r, yellow, 1)
	}

	filename := fmt.Sprintf("/tmp/%d.jpg", time.Now().UnixNano())
	gocv.IMWrite(filename, img)
	defer os.Remove(filename) // clean up

	return ioutil.ReadFile(filename)
}
