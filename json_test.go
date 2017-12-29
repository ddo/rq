package rq

import (
	"fmt"
	"testing"
)

var testJSON []byte

func TestJSONify(t *testing.T) {
	r := Put("https://ddo.me?k=1")

	r.Qs("k", "2", "3")
	r.Qs("v", "4")
	r.Qs("v", "5", "6")

	r.Send("data", "data value 1", "data value 2")
	r.Send("extra", "data value 3")

	data, err := r.JSONify()
	if err != nil {
		t.Error()
		return
	}
	if string(data) != `{"url":"https://ddo.me?k=1","method":"PUT","query":{"k":["2","3"],"v":["4","5","6"]},"form":{"data":["data value 1","data value 2"],"extra":["data value 3"]},"header":{}}` {
		t.Error()
		return
	}

	testJSON = data
}

func TestNewFromJSON(t *testing.T) {
	r, err := NewFromJSON(testJSON)
	if err != nil {
		fmt.Println(err)
		t.Error()
		return
	}
	if r == nil {
		t.Error()
		return
	}
	if r.URL != "https://ddo.me?k=1" {
		t.Error()
		return
	}
	if len(r.Query) != 2 || len(r.Form) != 2 || len(r.Header) != 0 {
		t.Error()
		return
	}
}

func TestNewFromJSONEmpty(t *testing.T) {
	r, err := NewFromJSON([]byte("{}"))
	if err != nil {
		fmt.Println(err)
		t.Error()
		return
	}
	if r != nil {
		t.Error()
		return
	}
}
