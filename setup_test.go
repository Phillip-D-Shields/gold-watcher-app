package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
	"testing"

	"fyne.io/fyne/v2/test"
)

var testApp Config

func TestMain(m *testing.M) {
	a := test.NewApp()
	testApp.App = a
	testApp.HTTPClient = client
	os.Exit(m.Run())
}

var jsonToReturn = `
{
  "ts": 1729581055439,
  "tsj": 1729581046058,
  "date": "Oct 22nd 2024, 03:10:46 am NY",
  "items": [
    {
      "curr": "USD",
      "xauPrice": 2735.5925,
      "xagPrice": 34.1855,
      "chgXau": 14.1325,
      "chgXag": 0.4535,
      "pcXau": 0.5193,
      "pcXag": 1.3444,
      "xauClose": 2721.46,
      "xagClose": 33.732
    }
  ]
}
`

type RoundTripFunc func(req *http.Request) *http.Response

func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

func NewTestClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: fn,
	}
}

var client = NewTestClient(func(req *http.Request) *http.Response {
	return &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(bytes.NewBufferString(jsonToReturn)),
		Header:     make(http.Header),
	}
})
