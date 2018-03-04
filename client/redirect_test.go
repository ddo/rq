package client

import (
	"testing"

	"github.com/ddo/rq"
)

func TestRedirectAfterN(t *testing.T) {
	_client := New(&Option{
		CheckRedirect: RedirectAfterN(5),
	})
	if _client.httpClient.CheckRedirect == nil {
		t.Error()
		return
	}

	r := rq.Get("https://httpbin.org/redirect/3")
	_, res, err := _client.Send(r, true)
	if err != nil {
		t.Error()
		return
	}
	if res.Request.URL.String() != "https://httpbin.org/get" {
		t.Error()
		return
	}

	// exceeded
	r = rq.Get("https://httpbin.org/redirect/5")
	_, res, err = _client.Send(r, true)
	if err != nil {
		t.Error()
		return
	}
	if res.Request.URL.String() != "https://httpbin.org/relative-redirect/1" {
		t.Error()
		return
	}
}

func TestRedirectDefault(t *testing.T) {
	_client := New(nil)
	if _client.httpClient.CheckRedirect == nil {
		t.Error()
		return
	}

	r := rq.Get("https://httpbin.org/redirect/6")
	_, res, err := _client.Send(r, true)
	if err != nil {
		t.Error()
		return
	}
	if res.Request.URL.String() != "https://httpbin.org/get" {
		t.Error()
		return
	}
}

func TestRedirectNoRedirect(t *testing.T) {
	_client := New(&Option{
		CheckRedirect: NoCheckRedirect,
	})
	if _client.httpClient.CheckRedirect == nil {
		t.Error()
		return
	}

	r := rq.Get("https://httpbin.org/redirect/1")
	_, res, err := _client.Send(r, true)
	if err != nil {
		t.Error()
		return
	}
	if res.Request.URL.String() != "https://httpbin.org/redirect/1" {
		t.Error()
		return
	}
}
