package client

import (
	"bytes"
	"testing"

	"github.com/ddo/pick-json"

	"github.com/ddo/rq"
)

func TestDefaultSend(t *testing.T) {
	r := rq.Post("https://httpbin.org/post?k=1")

	r.Qs("k", "2", "3")
	r.Qs("v", "4")

	r.Set("User-Agent", "github.com/ddo/rq")
	r.Set("Content-Type", "application/x-www-form-urlencoded")

	r.Send("data", "data value 1", "data value 2")
	r.Send("extra", "data value 3")

	data, res, err := Send(r, true)
	if err != nil {
		t.Error()
		return
	}
	if res.StatusCode != 200 {
		t.Error()
		return
	}
	if !verifyHTTPBinRes(data) {
		t.Error()
		return
	}
}

func TestSend(t *testing.T) {
	_client := New(nil)

	r := rq.Post("https://httpbin.org/post?k=1")

	r.Qs("k", "2", "3")
	r.Qs("v", "4")

	r.Set("User-Agent", "github.com/ddo/rq")
	r.Set("Content-Type", "application/x-www-form-urlencoded")

	r.Send("data", "data value 1", "data value 2")
	r.Send("extra", "data value 3")

	data, res, err := _client.Send(r, true)
	if err != nil {
		t.Error()
		return
	}
	if res.StatusCode != 200 {
		t.Error()
		return
	}
	if !verifyHTTPBinRes(data) {
		t.Error()
		return
	}
}

func TestSendErr(t *testing.T) {
	_client := New(nil)

	r := rq.Post("httpbin.org/post?k=1")

	data, res, err := _client.Send(r, true)
	if err == nil {
		t.Error()
		return
	}
	if res != nil {
		t.Error()
		return
	}
	if data != nil {
		t.Error()
		return
	}
}

// helper
func verifyHTTPBinRes(data []byte) bool {
	strs := pickjson.PickString(bytes.NewReader(data), "User-Agent", 1)
	if len(strs) == 0 || strs[0] != "github.com/ddo/rq" {
		return false
	}
	strs = pickjson.PickString(bytes.NewReader(data), "Content-Type", 1)
	if len(strs) == 0 || strs[0] != "application/x-www-form-urlencoded" {
		return false
	}
	strs = pickjson.PickString(bytes.NewReader(data), "url", 1)
	if len(strs) == 0 || strs[0] != "https://httpbin.org/post?k=1&k=2&k=3&v=4" {
		return false
	}
	strs = pickjson.PickString(bytes.NewReader(data), "extra", 1)
	if len(strs) == 0 || strs[0] != "data value 3" {
		return false
	}
	return true
}
