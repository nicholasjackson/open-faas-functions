package function

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"

	"github.com/matryer/is"
)

func TestDetectsFacesInImageAndReturnsImage(t *testing.T) {
	is := is.New(t)
	input := "group.jpg"
	defer setEnv("Http_Query", "output=image")()

	bytes, err := ioutil.ReadFile(input)
	is.NoErr(err) // Error should be nil

	resp := Handle(bytes)

	ioutil.WriteFile("./out.jpg", []byte(resp), os.ModePerm)
}

func TestDetectsFacesInImageAndReturnsJSON(t *testing.T) {
	is := is.New(t)
	input := "group.jpg"

	bytes, err := ioutil.ReadFile(input)
	is.NoErr(err) // Error should be nil

	j := &Response{}
	resp := Handle(bytes)
	json.Unmarshal([]byte(resp), j)

	is.Equal(true, len(j.Faces) > 0)
}

func TestDetectsFacesInImageAndReturnsJSONWithImage(t *testing.T) {
	is := is.New(t)
	input := "group.jpg"
	defer setEnv("Http_Query", "output=json_image")()

	bytes, err := ioutil.ReadFile(input)
	is.NoErr(err) // Error should be nil

	j := &Response{}
	resp := Handle(bytes)
	json.Unmarshal([]byte(resp), j)

	is.Equal(true, len(j.ImageBase64) > 0)
}

func setEnv(key, value string) func() {
	os.Setenv(key, value)

	return func() {
		os.Unsetenv(key)
	}
}
