package rq

import (
	"io"
)

// Rq .
type Rq struct {
	URL    string `json:"URL"`
	Method string `json:"Method"`

	Query map[string][]string `json:"query"`

	Form map[string][]string `json:"form"`
	Body io.Reader           `json:"-"`

	Header map[string]string `json:"header"`
}

// New .
func New(method, URL string) *Rq {
	return &Rq{
		URL:    URL,
		Method: method,
		Query:  map[string][]string{},
		Form:   map[string][]string{},
		Header: map[string]string{},
	}
}

// Get .
func Get(URL string) *Rq {
	return New("GET", URL)
}

// Post .
func Post(URL string) *Rq {
	return New("POST", URL)
}

// Put .
func Put(URL string) *Rq {
	return New("PUT", URL)
}

// Delete .
func Delete(URL string) *Rq {
	return New("DELETE", URL)
}

// Head .
func Head(URL string) *Rq {
	return New("HEAD", URL)
}

// Set sets request header
func (r *Rq) Set(key, value string) {
	r.Header[key] = value
}

// UnSet unsets request header
func (r *Rq) UnSet(key string) {
	delete(r.Header, key)
}

// Qs sets request query
func (r *Rq) Qs(key string, value ...string) {
	r.Query[key] = value
}

// UnQs unsets request query
func (r *Rq) UnQs(key string, value ...string) {
	delete(r.Query, key)
}

// Send sets request form
// and also set the request content type as application/x-www-form-urlencoded
// if there is no content type header
func (r *Rq) Send(key string, value ...string) {
	r.Form[key] = value
}

// UnSend unsets request body
func (r *Rq) UnSend(key string) {
	delete(r.Form, key)
}

// SendRaw sets request body directly and override the #Send
func (r *Rq) SendRaw(reader io.Reader) {
	r.Body = reader
}

// UnSendRaw unsets request body
func (r *Rq) UnSendRaw() {
	r.Body = nil
}
