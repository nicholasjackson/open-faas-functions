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

func setupTests(
	t *testing.T,
	tweetReturn anaconda.Tweet,
	mediaReturn anaconda.Media,
	err error) (*TwitterPosterMock, *is.I) {
	mockedTwitterPoster := &TwitterPosterMock{
		PostTweetFunc: func(status string, v url.Values) (anaconda.Tweet, error) {
			return tweetReturn, err
		},
		UploadMediaFunc: func(base64String string) (anaconda.Media, error) {
			return mediaReturn, err
		},
	}

	// clear the main dependency creation and inject our own mock
	setup = func() {}
	client = mockedTwitterPoster

	is := is.New(t)

	return mockedTwitterPoster, is
}

func TestInvalidRequestReturnsBadRequest(t *testing.T) {
	_, is := setupTests(t, anaconda.Tweet{}, anaconda.Media{}, nil)

	resp := marshalResponse(Handle([]byte("Nic")))

	is.Equal(http.StatusBadRequest, resp.Code) // expected bad request
}

func TestEmptyTextReturnsBadRequest(t *testing.T) {
	_, is := setupTests(t, anaconda.Tweet{}, anaconda.Media{}, nil)

	resp := marshalResponse(Handle([]byte(`{ "Text": "" }`)))

	is.Equal(http.StatusBadRequest, resp.Code) // expected bad request
}

func TestValidMessageSendsTweet(t *testing.T) {
	mt, is := setupTests(t, anaconda.Tweet{}, anaconda.Media{}, nil)
	message := "Hey @Nic"

	resp := marshalResponse(Handle([]byte(fmt.Sprintf(`{ "Text": "%s" }`, message))))

	is.Equal(1, len(mt.PostTweetCalls()))            // expected post tweet to be called
	is.Equal(message, mt.PostTweetCalls()[0].Status) // expected tweet message to be correct
	is.Equal(http.StatusOK, resp.Code)               // expected bad request
}

func TestFailSendingTweetReturnsErrro(t *testing.T) {
	_, is := setupTests(t, anaconda.Tweet{}, anaconda.Media{}, fmt.Errorf("booom"))
	message := "Hey @Nic"

	resp := marshalResponse(Handle([]byte(fmt.Sprintf(`{ "Text": "%s" }`, message))))

	is.Equal("Tweet failed to send: booom", resp.Message) // expected error message
	is.Equal(http.StatusInternalServerError, resp.Code)   // expected internal server error
}

func TestNotUploadsImageIfMissing(t *testing.T) {
	mt, is := setupTests(t, anaconda.Tweet{}, anaconda.Media{}, nil)
	message := "Hey @Nic"

	marshalResponse(Handle([]byte(fmt.Sprintf(`{ "Text": "%s" }`, message))))

	is.Equal(0, len(mt.UploadMediaCalls())) // should not have called media upload
}

func TestUploadsImageIfPresent(t *testing.T) {
	mt, is := setupTests(t, anaconda.Tweet{}, anaconda.Media{}, nil)
	message := "Hey @Nic"

	marshalResponse(Handle([]byte(fmt.Sprintf(`{ "Text": "%s", "Image": "abc=" }`, message))))

	is.Equal(1, len(mt.UploadMediaCalls())) // should have called media upload
}

func TestAddsMediaToImageIfPresent(t *testing.T) {
	mt, is := setupTests(
		t,
		anaconda.Tweet{},
		anaconda.Media{
			MediaIDString: "123",
		},
		nil,
	)
	message := "Hey @Nic"

	marshalResponse(Handle([]byte(fmt.Sprintf(`{ "text": "%s", "image": "abc=" }`, message))))

	is.Equal("123", mt.PostTweetCalls()[0].V.Get("media_ids")) // should have added media id to tweet
}

func marshalResponse(rsp string) Response {
	r := Response{}
	json.Unmarshal([]byte(rsp), &r)
	return r
}
