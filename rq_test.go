package rq

import (
	"strings"
	"testing"
)

func TestNew(t *testing.T) {
	rq := New("method", "url")
	if rq.URL != "url" {
		t.Error()
		return
	}
	if rq.Method != "method" {
		t.Error()
		return
	}
}

func TestGet(t *testing.T) {
	rq := Get("url")
	if rq.URL != "url" {
		t.Error()
		return
	}
	if rq.Method != "GET" {
		t.Error()
		return
	}
}

func TestPost(t *testing.T) {
	rq := Post("url")
	if rq.URL != "url" {
		t.Error()
		return
	}
	if rq.Method != "POST" {
		t.Error()
		return
	}
}

func TestPut(t *testing.T) {
	rq := Put("url")
	if rq.URL != "url" {
		t.Error()
		return
	}
	if rq.Method != "PUT" {
		t.Error()
		return
	}
}

func TestDelete(t *testing.T) {
	rq := Delete("url")
	if rq.URL != "url" {
		t.Error()
		return
	}
	if rq.Method != "DELETE" {
		t.Error()
		return
	}
}

func TestHead(t *testing.T) {
	rq := Head("url")
	if rq.URL != "url" {
		t.Error()
		return
	}
	if rq.Method != "HEAD" {
		t.Error()
		return
	}
}

func TestSet(t *testing.T) {
	rq := Get("url")
	rq.Set("User-Agent", "github.com/ddo/rq", "ddo")
	rq.Set("User-Agent", "ddo.me")
	if rq.Header["User-Agent"][0] != "github.com/ddo/rq" {
		t.Error()
		return
	}
	if rq.Header["User-Agent"][1] != "ddo" {
		t.Error()
		return
	}
	if rq.Header["User-Agent"][2] != "ddo.me" {
		t.Error()
		return
	}

	rq.UnSet("User-Agent")
	if _, ok := rq.Header["User-Agent"]; ok {
		t.Error()
		return
	}
}

func TestQs(t *testing.T) {
	rq := Get("url")
	rq.Qs("qs", "1", "22", "333")
	if !sameArr(rq.Query["qs"], []string{"1", "22", "333"}) {
		t.Error()
		return
	}

	rq.UnQs("qs")
	if _, ok := rq.Query["qs"]; ok {
		t.Error()
		return
	}
}

func TestSend(t *testing.T) {
	rq := Post("url")
	rq.Send("data", "1", "22", "333")
	if !sameArr(rq.Form["data"], []string{"1", "22", "333"}) {
		t.Error()
		return
	}

	rq.UnSend("data")
	if _, ok := rq.Form["data"]; ok {
		t.Error()
		return
	}
}

func TestSendRaw(t *testing.T) {
	data := strings.NewReader("data")

	rq := Post("url")
	rq.SendRaw(data)
	if rq.Body == nil {
		t.Error()
		return
	}

	rq.UnSendRaw()
	if rq.Body != nil {
		t.Error()
		return
	}
}

// helper
func sameArr(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
