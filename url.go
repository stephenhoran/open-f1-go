package openf1go

import (
	"io"
	"net/http"
)

type Arg struct {
	Key   string
	Value string
}

func ArgBuilder(args []Arg) string {
	var url string

	if len(args) == 0 {
		return url
	}

	if len(args) > 0 {
		url += "?"
	}

	for i, arg := range args {
		if i != 0 {
			url += "&"
		}

		url += arg.Key + "=" + arg.Value
	}

	return url
}

func GetHTTPRequest(url string) ([]byte, error) {
	resp, err := http.Get(url)
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
