package client

import (
	"net/http"
	"testing"
)

func TestSetProxy(t *testing.T) {
	client := New(nil)
	if client == nil {
		t.Error()
		return
	}

	err := client.SetProxy("https://127.0.0.1:8888")
	if err != nil {
		t.Error()
		return
	}

	proxy, err := client.httpClient.Transport.(*http.Transport).Proxy(nil)
	if err != nil {
		t.Error()
		return
	}
	if proxy.String() != "https://127.0.0.1:8888" {
		t.Error()
		return
	}
}

func TestSetProxyInvalidURL(t *testing.T) {
	client := New(nil)
	if client == nil {
		t.Error()
		return
	}

	err := client.SetProxy("127.0.0.1:8888")
	if err == nil {
		t.Error()
		return
	}
}

func TestSetProxyFromNew(t *testing.T) {
	client := New(&Option{
		Proxy: "https://127.0.0.1:8888",
	})
	if client == nil {
		t.Error()
		return
	}

	proxy, err := client.httpClient.Transport.(*http.Transport).Proxy(nil)
	if err != nil {
		t.Error()
		return
	}
	if proxy.String() != "https://127.0.0.1:8888" {
		t.Error()
		return
	}
}
