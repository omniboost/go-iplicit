package iplicit

import (
	"log"
	"net/url"
	"os"
	"testing"
)

var client *Client

func TestMain(m *testing.M) {
	client = NewClient(nil)
	client.Username = os.Getenv("IPLICIT_USERNAME")
	client.UserApiKey = os.Getenv("IPLICIT_USERAPIKEY")
	client.Domain = os.Getenv("IPLICIT_DOMAIN")

	baseURLString := os.Getenv("IPLICIT_BASE_URL")
	if baseURLString != "" {
		baseURL, err := url.Parse(baseURLString)
		if err != nil {
			log.Fatal(err)
		}
		client.BaseURL = baseURL
	}
	if os.Getenv("DEBUG") != "" {
		client.Debug = true
	}
	client.DisallowUnknownFields = false

	os.Exit(m.Run())
}
