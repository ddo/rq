# rq
A nicer interface for golang stdlib HTTP client

[![Doc][godoc-img]][godoc-url] rq

[![Doc][godoc-img]][godoc-client-url] client

[godoc-img]: https://img.shields.io/badge/godoc-Reference-brightgreen.svg?style=flat-square
[godoc-url]: https://godoc.org/gopkg.in/ddo/rq.v0
[godoc-client-url]: https://godoc.org/gopkg.in/ddo/rq.v0/client

## why?
because golang HTTP client is a pain in the a...

## features

* compatible golang ``http`` stdlib as ``http.Request``, ``http.Response`` and ``http.Cookie``
* step by step to build your **request**
* provide the easier way to work with **cookies**
* **import/export** allow we save/transfer requests as json ***SOON***
* **default setting**: example default ``user-agent`` or ``accept-language`` ***SOON***

## installation

```sh
go get -u gopkg.in/ddo/rq.v0
```

## getting started

simple

```go
import "github.com/ddo/rq"

r := rq.Get("https://httpbin.org/get")

// query https://httpbin.org/get?q=1&q=2&_=123456
r.Qs("q", "1")
r.Qs("q", "2")
r.Qs("_", "123456")

// send with golang default HTTP client
res, err := http.DefaultClient.Do(r.ParseRequest())
defer res.Body.Close()
```

in case you did not know that golang default HTTP client has **no timeout**.
use **rq/client** which has ``180s`` timeout by default

```go
import "github.com/ddo/rq"
import "github.com/ddo/rq/client"

r := rq.Post("https://httpbin.org/post")

// query
r.Qs("_", "123456")

// Form
r.Send("data", "data value")

// use default rq client
// true to tell #Send to read all the response boby when return
data, res, err := client.Send(r, true)
// no need to close res.Body
// read = false -> you need to call res.Body when done reading
```

set headers

```go
r := rq.Post("https://httpbin.org/post")

r.Set("Content-Type", "application/json")
r.Set("User-Agent", "ddo/rq")
```

raw body

```go
r := rq.Post("https://httpbin.org/post")

r.SendRaw(strings.NewReader("raw data binary or json"))
```

## client

```go
// by default timeout = 3min and has a cookie jar
customClient := client.New(nil)

// or custom timeout = 10s and no cookie jar
customClient := client.New(&Option{
    Timeout: time.Second * 10,
    NoCookie: true,
})
```

## cookies

```go
cookies, err := client.GetCookies("httpbin.org")
cookie, err := client.GetCookie("httpbin.org", "cookiename").

err := client.SetCookies("httpbin.org", cookies)
err := client.SetCookie("httpbin.org", cookie)

err := client.DelCookie("httpbin.org", "cookiename")
```

## TODO

check at #1