package function

import (
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func TestFunctionalTest(t *testing.T) {
	f, err := os.Open("./zac.jpg")
	if err != nil {
		t.Fail()
	}

	data, err := ioutil.ReadAll(f)
	if err != nil {
		t.Fail()
	}

	out := Handle(data)
	log.Println(out)
}
