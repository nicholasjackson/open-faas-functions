package function

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"testing"

	"github.com/ChimeraCoder/anaconda"
	"github.com/matryer/is"
)

func setupTests(t *testing.T, tweetReturn anaconda.Tweet, err error) (*TwitterPosterMock, *is.I) {
	mockedTwitterPoster := &TwitterPosterMock{
		PostTweetFunc: func(status string, v url.Values) (anaconda.Tweet, error) {
			return tweetReturn, err
		},
	}

	// clear the main dependency creation and inject our own mock
	setup = func() {}
	client = mockedTwitterPoster

	is := is.New(t)

	return mockedTwitterPoster, is
}

func TestInvalidRequestReturnsBadRequest(t *testing.T) {
	_, is := setupTests(t, anaconda.Tweet{}, nil)

	resp := marshalResponse(Handle([]byte("Nic")))

	is.Equal(http.StatusBadRequest, resp.Code) // expected bad request
}

func TestEmptyTextReturnsBadRequest(t *testing.T) {
	_, is := setupTests(t, anaconda.Tweet{}, nil)

	resp := marshalResponse(Handle([]byte(`{ "Text": "" }`)))

	is.Equal(http.StatusBadRequest, resp.Code) // expected bad request
}

func TestValidMessageSendsTweet(t *testing.T) {
	mt, is := setupTests(t, anaconda.Tweet{}, nil)
	message := "Hey @Nic"

	resp := marshalResponse(Handle([]byte(fmt.Sprintf(`{ "Text": "%s" }`, message))))

	is.Equal(1, len(mt.PostTweetCalls()))            // expected post tweet to be called
	is.Equal(message, mt.PostTweetCalls()[0].Status) // expected tweet message to be correct
	is.Equal(http.StatusOK, resp.Code)               // expected bad request
}

func TestFailSendingTweetReturnsErrro(t *testing.T) {
	_, is := setupTests(t, anaconda.Tweet{}, fmt.Errorf("booom"))
	message := "Hey @Nic"

	resp := marshalResponse(Handle([]byte(fmt.Sprintf(`{ "Text": "%s" }`, message))))

	is.Equal("Tweet failed to send: booom", resp.Message) // expected error message
	is.Equal(http.StatusInternalServerError, resp.Code)   // expected internal server error
}

func marshalResponse(rsp string) Response {
	r := Response{}
	json.Unmarshal([]byte(rsp), &r)
	return r
}
