package confluence

import (
	"errors"
	"net/http"
	"net/url"
)

// Flags entity as config for user event tracking service
type API struct {
	endPoint        *url.URL
	Client          HTTPClient
	username, token string
}

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

func NewAPI(location string, username string, token string) (*API, error) {
	if len(location) == 0 {
		return nil, errors.New("url empty")
	}

	endpoint, err := url.ParseRequestURI(location)
	if err != nil {
		return nil, err
	}

	api := API{
		endPoint: endpoint,
		username: username,
		token:    token,
	}

	api.Client = &http.Client{}

	return &api, nil
}

func NewAPIWithClient(location string, client HTTPClient) (*API, error) {
	if len(location) == 0 {
		return nil, errors.New("url empty")
	}

	endpoint, err := url.ParseRequestURI(location)

	if err != nil {
		return nil, err
	}

	if client == nil {
		return nil, errors.New("empty http client")
	}

	a := new(API)
	a.endPoint = endpoint
	a.Client = client

	return a, nil
}
