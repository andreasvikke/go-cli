package models

import "net/http"

type MockClient struct {
	DoFunc func(req *http.Request) (*http.Response, error)
}

func (c *MockClient) Do(req *http.Request) (*http.Response, error) {
	return c.DoFunc(req)
}
