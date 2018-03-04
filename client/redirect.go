package client

import (
	"net/http"
)

// RedirectAfterN returns http.Client.CheckRedirect func
func RedirectAfterN(n int) func(req *http.Request, via []*http.Request) error {
	return func(req *http.Request, via []*http.Request) error {
		log.Info(req.Method, "\t>>", req.URL.String())

		if len(via) >= n {
			return http.ErrUseLastResponse
		}
		return nil
	}
}

// DefaultCheckRedirect returns http.Client.CheckRedirect func with defaultRedirects
var DefaultCheckRedirect = RedirectAfterN(defaultRedirects)

// NoCheckRedirect returns http.Client.CheckRedirect func with no redirect
var NoCheckRedirect = RedirectAfterN(0)
