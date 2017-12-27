package rq

import (
	"net/http"
	"net/url"
	"strings"
)

// ParseRequest parses Rq as http.Request
func (r *Rq) ParseRequest() (req *http.Request, err error) {
	// url
	u, err := url.Parse(r.URL)
	if err != nil {
		return
	}

	// query
	qs := u.Query()
	for key, values := range r.Query {
		for _, value := range values {
			qs.Add(key, value)
		}
	}
	u.RawQuery = qs.Encode()

	// body
	reader := r.Body
	if reader == nil && len(r.Form) > 0 {
		// set content type
		if _, ok := r.Header["Content-Type"]; !ok {
			r.Set("Content-Type", "application/x-www-form-urlencoded")
		}

		body := url.Values{}
		for key, values := range r.Form {
			for _, value := range values {
				body.Add(key, value)
			}
		}
		reader = strings.NewReader(body.Encode())
	}

	req, err = http.NewRequest(r.Method, u.String(), reader)
	if err != nil {
		return
	}

	// header
	for key, value := range r.Header {
		req.Header.Add(key, value)
	}

	return
}
