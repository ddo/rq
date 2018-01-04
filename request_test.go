package rq

import (
	"io/ioutil"
	"strings"
	"testing"
)

func TestParseRequest(t *testing.T) {
	rq := Put("https://ddo.me?k=1")

	rq.Qs("k", "2", "3")
	rq.Qs("v", "4")

	rq.Set("User-Agent", "github.com/ddo/rq")

	rq.Send("data", "data value 1", "data value 2")
	rq.Send("extra", "data value 3")

	req, err := rq.ParseRequest()
	if err != nil {
		t.Error()
		return
	}
	if req == nil {
		t.Error()
		return
	}
	if req.Method != "PUT" {
		t.Error()
		return
	}
	if req.URL.String() != "https://ddo.me?k=1&k=2&k=3&v=4" {
		t.Error()
		return
	}
	if req.Header.Get("user-agent") != "github.com/ddo/rq" {
		t.Error()
		return
	}

	// body
	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		t.Error()
		return
	}
	if string(data) != "data=data+value+1&data=data+value+2&extra=data+value+3" {
		t.Error()
		return
	}
}

func TestParseRequestRaw(t *testing.T) {
	rq := Put("https://ddo.me?k=1")

	rq.Qs("k", "2", "3")
	rq.Qs("v", "4")

	rq.Set("User-Agent", "github.com/ddo/rq")

	rq.Send("data", "data value 1", "data value 2")
	rq.Send("extra", "data value 3")

	rq.SendRaw(strings.NewReader("rawdata"))

	req, err := rq.ParseRequest()
	if err != nil {
		t.Error()
		return
	}
	if req == nil {
		t.Error()
		return
	}

	// body
	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		t.Error()
		return
	}
	if string(data) != "rawdata" {
		t.Error()
		return
	}
}

func TestParseRequestInvalidURL(t *testing.T) {
	rq := Get(":")

	req, err := rq.ParseRequest()
	if err == nil {
		t.Error()
		return
	}
	if req != nil {
		t.Error()
		return
	}
}

func TestParseRequestInvalidMethod(t *testing.T) {
	rq := Get("https://ddo.me")
	rq.Method = ":"

	req, err := rq.ParseRequest()
	if err == nil {
		t.Error()
		return
	}
	if req != nil {
		t.Error()
		return
	}
}
