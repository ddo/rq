package jar

import (
	"net/http"
	"reflect"
	"testing"
)

const (
	testHostname = "test.com"
)

var testCookies = []*http.Cookie{
	{Name: "k1", Value: "v1"},
	{Name: "k2", Value: "v2"},
}

func TestNew(t *testing.T) {
	j := New()
	if j == nil {
		t.Error()
		return
	}

	if j._jar == nil {
		t.Error()
		return
	}
}

func TestGetHTTPJar(t *testing.T) {
	j := New()
	if j == nil {
		t.Error()
		return
	}

	if j._jar != j.GetHTTPJar() {
		t.Error()
		return
	}
}

func TestGet(t *testing.T) {
	j := New()

	err := j.Set(testHostname, testCookies)
	if err != nil {
		t.Error()
		return
	}

	cookies, err := j.Get(testHostname)
	if err != nil {
		t.Error()
		return
	}

	if !reflect.DeepEqual(cookies, testCookies) {
		t.Error()
		return
	}
}

func TestGetByName(t *testing.T) {
	j := New()

	err := j.Set(testHostname, testCookies)
	if err != nil {
		t.Error()
		return
	}

	cookie, err := j.GetByName(testHostname, testCookies[0].Name)
	if err != nil {
		t.Error()
		return
	}

	if !reflect.DeepEqual(cookie, testCookies[0]) {
		t.Error()
		return
	}
}

func TestSet(t *testing.T) {
	j := New()

	err := j.Set(testHostname, testCookies)
	if err != nil {
		t.Error()
		return
	}

	cookies, err := j.Get(testHostname)
	if err != nil {
		t.Error()
		return
	}

	if !reflect.DeepEqual(cookies, testCookies) {
		t.Error()
		return
	}
}

func TestSetOne(t *testing.T) {
	j := New()

	err := j.SetOne(testHostname, testCookies[0])
	if err != nil {
		t.Error()
		return
	}

	cookies, err := j.Get(testHostname)
	if err != nil {
		t.Error()
		return
	}

	if !reflect.DeepEqual(cookies[0], testCookies[0]) {
		t.Error()
		return
	}
}

func TestDelete(t *testing.T) {
	j := New()

	err := j.Set(testHostname, testCookies)
	if err != nil {
		t.Error()
		return
	}

	err = j.Delete(testHostname, testCookies[0].Name)
	if err != nil {
		t.Error()
		return
	}

	cookies, err := j.Get(testHostname)
	if err != nil {
		t.Error()
		return
	}

	if !reflect.DeepEqual(cookies[0], testCookies[1]) {
		t.Error()
		return
	}
}

func TestClear(t *testing.T) {
	j := New()

	err := j.Set(testHostname, testCookies)
	if err != nil {
		t.Error()
		return
	}

	err = j.Clear(testHostname)
	if err != nil {
		t.Error()
		return
	}

	cookies, err := j.Get(testHostname)
	if err != nil {
		t.Error()
		return
	}

	if len(cookies) != 0 {
		t.Error()
		return
	}
}
