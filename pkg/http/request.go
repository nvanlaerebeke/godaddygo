package http

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/oze4/godaddygo/internal/validator"
)

// RequestMethods holds all acceptable request methods
var RequestMethods = map[string]string{
	"GET":    "GET",
	"POST":   "POST",
	"DELETE": "DELETE",
	"PATCH":  "PATCH",
	"PUT":    "PUT",
}

// Request holds request data
type Request struct {
	APIKey    string
	APISecret string
	Method    string
	URL       string
	Host      string
	Body      []byte
}

// Do sends the http request
func (r *Request) Do() (bodyBytes []byte, err error) {
	valid := validator.Validate(r.Method, RequestMethods)
	if valid != true {
		return nil, errors.New("Invalid request method")
	}

	parsedURL, err := url.Parse(r.URL)
	if err != nil {
		return nil, err
	}

	req := &http.Request{
		URL:    parsedURL,
		Method: r.Method,
		Header: map[string][]string{
			"Authorization": {r.makeAuthString()},
			"Content-Type":  {"application/json"},
		},
		Body: ioutil.NopCloser(strings.NewReader(string(r.Body))),
	}

	result, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	bod, err := ioutil.ReadAll(result.Body)
	if err != nil {
		return nil, err
	}

	result.Body.Close()
	
	return bod, nil
}

func (r *Request) makeAuthString() string {
	return "sso " + r.APIKey + ":" + r.APISecret
}
