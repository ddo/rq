package client

import (
	"io/ioutil"
	"net/http"
	"time"

	"github.com/ddo/rq"
)

// Send sends the request and read it if read is true
// remember to close the res.Body if read is false
// Send also set content-type to application/x-www-form-urlencoded
// if form is available
func (c *Client) Send(r *rq.Rq, read bool) (data []byte, res *http.Response, err error) {
	// apply defaultRq
	newRq := ApplyDefaultRq(c.defaultRq, r)

	req, err := newRq.ParseRequest()
	if err != nil {
		log.Error(err)
		return
	}

	// set content type for form
	if req.Header.Get("content-type") == "" && len(newRq.Form) > 0 && newRq.Body == nil {
		req.Header.Set("content-type", "application/x-www-form-urlencoded")
	}

	if req.Header.Get("user-agent") == "" {
		req.Header.Set("user-agent", defaultUserAgent)
	}

	log.Info(req.Method, "\t>", req.URL.String())
	now := time.Now()

	res, err = c.httpClient.Do(req)
	if err != nil {
		log.Error("HTTP\t<", err, humanizeNano(time.Now().Sub(now)))
		return
	}
	log.Info(res.StatusCode, "\t<", res.Request.URL, humanizeNano(time.Now().Sub(now)))

	// stop here if read is false
	if !read {
		return
	}

	// read the response
	defer res.Body.Close()
	data, err = ioutil.ReadAll(res.Body)
	return
}

// Send is the wrapper of #Send but use the default client
func Send(r *rq.Rq, read bool) (data []byte, res *http.Response, err error) {
	return DefaultClient.Send(r, read)
}
