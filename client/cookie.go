package client

import (
	"net/http"
)

// GetCookies returns cookies by hostname
func (c *Client) GetCookies(hostname string) (cookies []*http.Cookie, err error) {
	// skip if no jar
	if c.httpClient.Jar == nil {
		return
	}

	u, err := parseHostname(hostname)
	if err != nil {
		return
	}

	cookies = c.httpClient.Jar.Cookies(u)
	log.Debug("cookies:", cookies)
	return
}

// GetCookie returns a cookie by cookie name
func (c *Client) GetCookie(hostname, name string) (cookie *http.Cookie, err error) {
	cookies, err := c.GetCookies(hostname)
	if err != nil {
		return
	}

	for i := 0; i < len(cookies); i++ {
		if cookies[i].Name == name {
			cookie = cookies[i]
			log.Debug("cookie:", cookie)
			return
		}
	}
	return
}

// SetCookies sets client cookies with the hostname
func (c *Client) SetCookies(hostname string, cookies []*http.Cookie) (err error) {
	// skip if no jar
	if c.httpClient.Jar == nil {
		return
	}

	u, err := parseHostname(hostname)
	if err != nil {
		return
	}

	c.httpClient.Jar.SetCookies(u, cookies)
	return
}

// SetCookie is a wrapper of #SetCookies but only a cookie
func (c *Client) SetCookie(hostname string, cookie *http.Cookie) (err error) {
	return c.SetCookies(hostname, []*http.Cookie{cookie})
}

// DelCookie is a wrapper of #SetCookies to remove a cookie
func (c *Client) DelCookie(hostname string, name string) (err error) {
	return c.SetCookie(hostname, &http.Cookie{
		Name:   name,
		MaxAge: -1,
	})
}
