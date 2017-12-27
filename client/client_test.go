package client

import (
	"net/http/cookiejar"
	"testing"
	"time"
)

func TestNewDefaultClient(t *testing.T) {
	if DefaultClient.httpClient.Timeout != defaultTimeout {
		t.Error()
		return
	}
	if DefaultClient.httpClient.Jar != nil {
		t.Error()
		return
	}
	if DefaultClient.httpClient.Transport != nil {
		t.Error()
		return
	}
}

func TestNew(t *testing.T) {
	_client := New(nil)
	if _client.httpClient.Timeout != defaultTimeout {
		t.Error()
		return
	}
	if _client.httpClient.Jar == nil {
		t.Error()
		return
	}
	if _client.httpClient.Transport != nil {
		t.Error()
		return
	}
}

func TestNewWithTimeout(t *testing.T) {
	_client := New(&Option{
		Timeout: 5 * time.Second,
	})
	if _client.httpClient.Timeout != 5*time.Second {
		t.Error()
		return
	}
	if _client.httpClient.Jar == nil {
		t.Error()
		return
	}
	if _client.httpClient.Transport != nil {
		t.Error()
		return
	}
}

func TestNewWithNoTimeout(t *testing.T) {
	_client := New(&Option{
		Timeout:   5 * time.Second,
		NoTimeout: true,
	})
	if _client.httpClient.Timeout != 0 {
		t.Error()
		return
	}
	if _client.httpClient.Jar == nil {
		t.Error()
		return
	}
	if _client.httpClient.Transport != nil {
		t.Error()
		return
	}
}

func TestNewWithCookie(t *testing.T) {
	jar, _ := cookiejar.New(nil)
	_client := New(&Option{
		Jar: jar,
	})
	if _client.httpClient.Timeout != defaultTimeout {
		t.Error()
		return
	}
	if _client.httpClient.Jar == nil {
		t.Error()
		return
	}
	if _client.httpClient.Transport != nil {
		t.Error()
		return
	}
}

func TestNewWithNoCookie(t *testing.T) {
	jar, _ := cookiejar.New(nil)
	_client := New(&Option{
		Jar:      jar,
		NoCookie: true,
	})
	if _client.httpClient.Timeout != defaultTimeout {
		t.Error()
		return
	}
	if _client.httpClient.Jar != nil {
		t.Error()
		return
	}
	if _client.httpClient.Transport != nil {
		t.Error()
		return
	}
}
