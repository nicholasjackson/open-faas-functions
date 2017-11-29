package function

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/ChimeraCoder/anaconda"
)

// Request defines the input into the function
type Request struct {
	Text string
}

// Response defines the response from the function
type Response struct {
	Code    int
	Message string
}

//go:generate moq -out mocks_test.go . TwitterPoster

// TwitterPoster defines the behaviour for posting a tweet
type TwitterPoster interface {
	PostTweet(status string, v url.Values) (tweet anaconda.Tweet, err error)
}

var consumerToken = os.Getenv("twitter_consumer_key")
var consumerSecret = os.Getenv("twitter_consumer_secret")
var accessToken = os.Getenv("twitter_access_token")
var accessTokenSecret = os.Getenv("twitter_access_token_secret")
var client TwitterPoster

type setupFunction func()

var setup = setupDeps

// Handle a serverless request
func Handle(req []byte) string {
	setup()

	r, err := marshalRequest(req)
	if err != nil {
		return createResponse(http.StatusBadRequest, "Invalid request message")
	}

	if r.Text == "" {
		return createResponse(http.StatusBadRequest, "Empty message")
	}

	_, err = client.PostTweet(r.Text, nil)
	if err != nil {
		return createResponse(
			http.StatusInternalServerError,
			fmt.Sprintf("Tweet failed to send: %s", err),
		)
	}

	return createResponse(http.StatusOK, "Tweet sent")
}

func setupDeps() {
	anaconda.SetConsumerKey(consumerToken)
	anaconda.SetConsumerSecret(consumerSecret)
	client = anaconda.NewTwitterApi(accessToken, accessTokenSecret)
}

func marshalRequest(req []byte) (Request, error) {
	r := Request{}
	return r, json.Unmarshal(req, &r)
}

func createResponse(code int, message string) string {
	r := Response{Code: code, Message: message}

	rstring, _ := json.Marshal(r)

	return string(rstring)
}
