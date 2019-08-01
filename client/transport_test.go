package client

import (
	"net/http"
	"testing"
)

const (
	test_proxy_url = "https://127.0.0.1:8888"
)

func TestSetProxy(t *testing.T) {
	client := New(nil)
	if client == nil {
		t.Error()
		return
	}

	err := client.SetProxy(test_proxy_url)
	if err != nil {
		t.Error()
		return
	}

	transport := client.httpClient.Transport.(*http.Transport)
	proxy, err := transport.Proxy(nil)
	if err != nil {
		t.Error()
		return
	}
	if proxy.String() != test_proxy_url {
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
		Proxy: test_proxy_url,
	})
	if client == nil {
		t.Error()
		return
	}

	transport := client.httpClient.Transport.(*http.Transport)
	proxy, err := transport.Proxy(nil)
	if err != nil {
		t.Error()
		return
	}
	if proxy.String() != test_proxy_url {
		t.Error()
		return
	}
}

func TestUnSetProxy(t *testing.T) {
	client := New(&Option{
		Proxy: test_proxy_url,
	})
	if client == nil {
		t.Error()
		return
	}

	client.UnSetProxy()

	transport := client.httpClient.Transport.(*http.Transport)
	if transport.Proxy != nil {
		t.Error()
		return
	}
}
