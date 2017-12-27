package client

import (
	"net/url"
	"testing"

	"github.com/ddo/rq"
)

func TestCookie(t *testing.T) {
	_client := New(nil)

	r := rq.Get("https://httpbin.org/cookies/set")
	r.Qs("k2", "v2")
	r.Qs("k1", "v1")

	_, _, err := _client.Send(r, true)
	if err != nil {
		t.Error()
		return
	}

	u, _ := url.Parse("http://httpbin.org")
	cookies := _client.httpClient.Jar.Cookies(u)
	if len(cookies) != 2 {
		t.Error()
		return
	}
	for i := 0; i < len(cookies); i++ {
		if cookies[i].Name == "k1" && cookies[i].Value != "v1" {
			t.Error()
			return
		}
		if cookies[i].Name == "k2" && cookies[i].Value != "v2" {
			t.Error()
			return
		}
	}
}
