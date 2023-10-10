package confluence

import (
	"errors"
	"net/http"
	"net/url"
)

// Flags entity as config for user event tracking service
type API struct {
	endPoint        *url.URL
	Client          *http.Client
	username, token string
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
