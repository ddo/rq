package client

import (
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
}
