package main

import (
	"bytes"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/DATA-DOG/godog"
)

func main() { /* usual main func */ }

var imageData []byte
var responseData []byte

func iHaveAValidImage() error {
	f, err := os.Open("./good.jpg")
	if err != nil {
		return err
	}

	defer f.Close()

	imageData, err = ioutil.ReadAll(f)
	if err != nil {
		return err
	}

	return nil
}

func iCallMyFunction() error {
	server := os.Getenv("FAAS_SERVER")

	resp, err := http.Post(fmt.Sprintf("http://%s:8081", server), "application/jpeg", bytes.NewReader(imageData))
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Expected status 200, got %d", resp.StatusCode)
	}
	defer resp.Body.Close()

	responseData, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func iExpectItToReturnAValidImage() error {
	_, t, err := image.Decode(bytes.NewReader(responseData))
	if err != nil {
		return err
	}

	if t == image.ErrFormat.Error() {
		return fmt.Errorf("Expected image, got %s", string(responseData))
	}

	return nil
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^I have a valid image$`, iHaveAValidImage)
	s.Step(`^I call my function$`, iCallMyFunction)
	s.Step(`^I expect it to return a valid image$`, iExpectItToReturnAValidImage)
}
