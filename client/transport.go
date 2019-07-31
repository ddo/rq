package client

import (
	"errors"
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

	// handle nil transport
	client := c.httpClient
	if client.Transport == nil {

		// this default transport is from net/http DefaultTransport
		client.Transport = &http.Transport{
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

	// we need to conversion client.Transport which is RoundTripper interface to *http.Transport to set Proxy
	transport, ok := client.Transport.(*http.Transport)
	if !ok {
		err = errors.New("client.Transport (http.RoundTripper - interface) is not *http.Transport -> no Proxy")
		log.Error(err)
		return
	}

	transport.Proxy = http.ProxyURL(u)
	log.Info("DONE:", u)
	return
}
