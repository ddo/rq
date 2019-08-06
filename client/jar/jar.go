package jar

import (
	"net/http"
	"net/http/cookiejar"

	"gopkg.in/ddo/go-dlog.v2"
)

var log = dlog.New("rq:jar", nil)

// Jar contains stdlib http cookie jar
type Jar struct {
	_jar http.CookieJar
}

// New returns new empty cookie jar
func New() *Jar {
	_jar, _ := cookiejar.New(nil)

	return &Jar{
		_jar: _jar,
	}
}

// GetHTTPJar returns stdlib http.CookieJar
func (j *Jar) GetHTTPJar() http.CookieJar {
	return j._jar
}

// Get returns all cookies by hostname
func (j *Jar) Get(hostname string) (cookies []*http.Cookie, err error) {
	u, err := parseHostname(hostname)
	if err != nil {
		return
	}

	cookies = j._jar.Cookies(u)
	log.Debug("cookies:", cookies)
	return
}

// GetByName returns a cookie by it's hostname and name
func (j *Jar) GetByName(hostname, name string) (cookie *http.Cookie, err error) {
	cookies, err := j.Get(hostname)
	if err != nil {
		return
	}

	for _, _cookie := range cookies {
		if _cookie.Name == name {
			cookie = _cookie
			log.Debug("cookie:", cookie)
			return
		}
	}

	return
}

// Set sets cookies by the hostname
func (j *Jar) Set(hostname string, cookies []*http.Cookie) (err error) {
	u, err := parseHostname(hostname)
	if err != nil {
		return
	}

	j._jar.SetCookies(u, cookies)
	return
}

// SetOne sets a cookie by the hostname
func (j *Jar) SetOne(hostname string, cookie *http.Cookie) (err error) {
	return j.Set(hostname, []*http.Cookie{cookie})
}

// Delete deletes a cookie by set max-age=-1
func (j *Jar) Delete(hostname string, name string) (err error) {
	return j.SetOne(hostname, &http.Cookie{
		Name:   name,
		MaxAge: -1,
	})
}

// Clear deletes all cookies by set max-age=-1
func (j *Jar) Clear(hostname string) (err error) {
	cookies, err := j.Get(hostname)
	if err != nil {
		return
	}

	for _, cookie := range cookies {
		err = j.Delete(hostname, cookie.Name)
		if err != nil {
			return
		}
	}

	return
}
