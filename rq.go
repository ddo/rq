package rq

import (
	"io"
)

// Rq contains nicer http request interface components
type Rq struct {
	URL    string `json:"url"`
	Method string `json:"method"`

	Query  map[string][]string `json:"query"`
	Form   map[string][]string `json:"form"`
	Header map[string][]string `json:"header"`

	Body io.Reader `json:"-"`
}

// New returms empty Rq object
func New(method, URL string) *Rq {
	return &Rq{
		URL:    URL,
		Method: method,

		Query:  map[string][]string{},
		Form:   map[string][]string{},
		Header: map[string][]string{},
	}
}

// Get is a shortcut of #New
func Get(URL string) *Rq {
	return New("GET", URL)
}

// Post is a shortcut of #New
func Post(URL string) *Rq {
	return New("POST", URL)
}

// Put is a shortcut of #New
func Put(URL string) *Rq {
	return New("PUT", URL)
}

// Delete is a shortcut of #New
func Delete(URL string) *Rq {
	return New("DELETE", URL)
}

// Head is a shortcut of #New
func Head(URL string) *Rq {
	return New("HEAD", URL)
}

// Set sets request header
func (r *Rq) Set(key string, value ...string) {
	r.Header[key] = append(r.Header[key], value...)
}

// UnSet unsets request header
func (r *Rq) UnSet(key string) {
	delete(r.Header, key)
}

// Qs sets request query
func (r *Rq) Qs(key string, value ...string) {
	r.Query[key] = append(r.Query[key], value...)
}

// UnQs unsets request query
func (r *Rq) UnQs(key string) {
	delete(r.Query, key)
}

// Send sets request form
// and also set the request content type as application/x-www-form-urlencoded
// if there is no content type header
func (r *Rq) Send(key string, value ...string) {
	r.Form[key] = append(r.Form[key], value...)
}

// UnSend unsets request form
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
