package client

import (
	"io/ioutil"
	"net/http"
	"time"

	"github.com/ddo/rq"
)

// Send sends the request and read it if read is true
// remember to close the res.Body if read is false
func (c *Client) Send(r *rq.Rq, read bool) (data []byte, res *http.Response, err error) {
	req, err := r.ParseRequest()
	if err != nil {
		log.Error(err)
		return
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
