package openf1go

import (
	"io"
	"net/http"
	"net/url"
)

type Arg struct {
	Key   string
	Value string
}

func UrlBuilder(s string, args []Arg) (*url.URL, error) {
	u, err := url.Parse(s)
	if err != nil {
		return nil, err
	}

	q := url.Values{}

	for _, arg := range args {
		q.Set(arg.Key, arg.Value)
	}

	u.RawQuery = q.Encode()

	return u, nil
}

func GetHTTPRequest(url *url.URL) ([]byte, error) {
	resp, err := http.Get(url.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, err
}
