package httpclient

import (
	"io"
	"net/http"
	"net/url"
)

func NewApiHttpRequest(method string, url string, queryParams url.Values, body io.Reader) (*http.Request, error) {

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.URL.RawQuery = queryParams.Encode()

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	return req, nil
}

func DoRequest(r *http.Request) (*http.Response, error) {
	c := http.Client{}
	return c.Do(r)
}
