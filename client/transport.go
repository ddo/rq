package client

import (
	"net"
	"net/http"
	"net/url"
	"time"
)

// SetProxy sets client proxy by url string
func (c *Client) SetProxy(rawURL string) (err error) {
	log.Info("url:", rawURL)

	u, err := url.Parse(rawURL)
	if err != nil {
		log.Error(err)
		return
	}

	transport, ok := c.httpClient.Transport.(*http.Transport)
	if !ok {
		transport = nil
	}

	if transport == nil {
		transport = &http.Transport{
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
				DualStack: true,
			}).DialContext,
			MaxIdleConns:          100,
			IdleConnTimeout:       90 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		}
	}

	transport.Proxy = func(req *http.Request) (*url.URL, error) {
		return u, nil
	}

	log.Info("DONE:", u)
	return
}
