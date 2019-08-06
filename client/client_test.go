package client

import (
	"testing"
	"time"

	"github.com/ddo/rq"
	"github.com/ddo/rq/client/jar"
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
	if _client.httpClient.Jar != nil {
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
	if _client.httpClient.Jar != nil {
		t.Error()
		return
	}
	if _client.httpClient.Transport != nil {
		t.Error()
		return
	}
}

func TestNewWithCookie(t *testing.T) {
	jar := jar.New()
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
	_client := New(&Option{})
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

func TestDefaultRq(t *testing.T) {
	defaultRq := rq.New("", "http://ddo.me")
	defaultRq.Qs("_", "123456")
	defaultRq.Qs("qs", "2")
	defaultRq.Send("extra", "data")
	defaultRq.Set("User-Agent", "ddo/rq")

	r := rq.Post("https://example.com")
	r.Qs("qs", "1")
	r.Send("data", "1", "2")
	r.Set("Accept", "*/*")

	newRq := ApplyDefaultRq(defaultRq, r)

	if newRq.URL != "https://example.com" {
		t.Error()
		return
	}
	if newRq.Method != "POST" {
		t.Error()
		return
	}
	if newRq.Query["_"][0] != "123456" || newRq.Query["qs"][0] != "1" {
		t.Error()
		return
	}
	if newRq.Form["data"][0] != "1" || newRq.Form["data"][1] != "2" || newRq.Form["extra"][0] != "data" {
		t.Error()
		return
	}
	if newRq.Header["Accept"][0] != "*/*" || newRq.Header["User-Agent"][0] != "ddo/rq" {
		t.Error()
		return
	}
}
