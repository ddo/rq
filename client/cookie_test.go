package client

import (
	"net/http"
	"testing"

	"github.com/ddo/rq"
)

var testCookieClient *Client

func TestCookieInit(t *testing.T) {
	testCookieClient = New(nil)

	r := rq.Get("https://httpbin.org/cookies/set")
	r.Qs("k2", "v2")
	r.Qs("k1", "v1")

	_, _, err := testCookieClient.Send(r, true)
	if err != nil {
		t.Error()
		return
	}
}

func TestGetCookies(t *testing.T) {
	cookies, err := testCookieClient.GetCookies("httpbin.org")
	if err != nil {
		t.Error()
		return
	}
	if len(cookies) != 2 {
		t.Error()
		return
	}

	cookies, err = testCookieClient.GetCookies("http://httpbin.org")
	if err != nil {
		t.Error()
		return
	}
	if len(cookies) != 2 {
		t.Error()
		return
	}
}

func TestGetCookie(t *testing.T) {
	cookie, err := testCookieClient.GetCookie("httpbin.org", "k1")
	if err != nil {
		t.Error()
		return
	}
	if cookie == nil || cookie.Name != "k1" || cookie.Value != "v1" {
		t.Error()
		return
	}
}

func TestGetCookieEmpty(t *testing.T) {
	cookie, err := testCookieClient.GetCookie("httpbin.org", "k11")
	if err != nil {
		t.Error()
		return
	}
	if cookie != nil {
		t.Error()
		return
	}
}

func TestSetCookies(t *testing.T) {
	cookies := []*http.Cookie{
		{Name: "k3", Value: "v3"},
		{Name: "k4", Value: "v4"},
		{Name: "k5", Value: "v5"},
	}
	err := testCookieClient.SetCookies("httpbin.org", cookies)
	if err != nil {
		t.Error()
		return
	}

	newCookies, err := testCookieClient.GetCookies("http://httpbin.org")
	if err != nil {
		t.Error()
		return
	}
	if len(newCookies) != 5 {
		t.Error()
		return
	}
}

func TestSetCookie(t *testing.T) {
	cookie := &http.Cookie{
		Name:  "k6",
		Value: "v6",
	}
	err := testCookieClient.SetCookie("httpbin.org", cookie)
	if err != nil {
		t.Error()
		return
	}

	newCookies, err := testCookieClient.GetCookies("http://httpbin.org")
	if err != nil {
		t.Error()
		return
	}
	if len(newCookies) != 6 {
		t.Error()
		return
	}
}

func TestSetCookieReplace(t *testing.T) {
	cookie := &http.Cookie{
		Name:  "k2",
		Value: "v22",
	}
	err := testCookieClient.SetCookie("httpbin.org", cookie)
	if err != nil {
		t.Error()
		return
	}

	newCookie, err := testCookieClient.GetCookie("http://httpbin.org", "k2")
	if err != nil {
		t.Error()
		return
	}
	if newCookie == nil || newCookie.Name != "k2" || newCookie.Value != "v22" {
		t.Error()
		return
	}
}

func TestDelCookie(t *testing.T) {
	err := testCookieClient.DelCookie("httpbin.org", "k1")
	if err != nil {
		t.Error()
		return
	}

	newCookies, err := testCookieClient.GetCookies("http://httpbin.org")
	if err != nil {
		t.Error()
		return
	}
	if len(newCookies) != 5 {
		t.Error()
		return
	}
}
