package confluence_test

import (
	"confluence-poc/src/confluence"
	"confluence-poc/src/models"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func TestHandler_GetById(t *testing.T) {
	client := &models.MockClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			if req.URL.String() == "https://test.atlassian.net/wiki/rest/api/content/1234" {
				return &http.Response{
					StatusCode: 200,
					Body:       ioutil.NopCloser(strings.NewReader(`{"id":"1234"}`)),
				}, nil
			}

			return nil, errors.New(
				"Error from web server",
			)
		},
	}

	api, err := confluence.NewAPIWithClient("https://test.atlassian.net/wiki/rest/api", client)
	if err != nil {
		t.Error(err)
	}

	value := api.GetContentByID([]string{"4321"}, models.ContentQuery{})
	t.Log(value.Errors)
	if len(value.Errors) == 0 {
		t.Errorf("Expected errors to be more than 0, got %d", len(value.Errors))
	}

	value = api.GetContentByID([]string{"1234"}, models.ContentQuery{})
	t.Log(value.Data)
	if len(value.Data) == 0 {
		t.Errorf("Expected data to be more than 0, got %d", len(value.Data))
	}
}
