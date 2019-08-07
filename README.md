# rq [![Build Status][semaphoreci-img]][semaphoreci-url] [![Go Report][goreport-img]][goreport-url] [![codecov][codecov-img]][codecov-url]

A nicer interface for golang stdlib HTTP client

## Documents

-   rq: [here][godoc-url]
-   client: [here][godoc-client-url]
-   jar: [here][godoc-jar-url]

[godoc-img]: https://img.shields.io/badge/godoc-Reference-brightgreen.svg
[godoc-url]: https://godoc.org/github.com/ddo/rq
[godoc-client-url]: https://godoc.org/github.com/ddo/rq/client
[godoc-jar-url]: https://godoc.org/github.com/ddo/rq/client/jar
[semaphoreci-img]: https://semaphoreci.com/api/v1/ddo/rq/branches/master/shields_badge.svg
[semaphoreci-url]: https://semaphoreci.com/ddo/rq
[goreport-img]: https://goreportcard.com/badge/github.com/ddo/rq
[goreport-url]: https://goreportcard.com/report/github.com/ddo/rq
[codecov-img]: https://codecov.io/gh/ddo/rq/branch/master/graph/badge.svg
[codecov-url]: https://codecov.io/gh/ddo/rq

## Why?

Because golang HTTP client is a pain in the a...

## Features

-   Compatible with golang `http` stdlib: `http.Request`, `http.Response` and `http.Cookie`
-   Step by step to build your **request**
-   Better HTTP **client**
-   Better **cookie jar**
-   **Import/export** allow we save/transfer requests in JSON
-   **Default setting**: example default `User-Agent` or `Accept-Language`

## Installation

```sh
go get -u github.com/ddo/rq
```

## Getting started

### Simple

```go
import "net/http"
import "github.com/ddo/rq"

r := rq.Get("https://httpbin.org/get")

// query https://httpbin.org/get?q=1&q=2&q=3&_=123456
r.Qs("q", "1", "2")
r.Qs("q", "3")
r.Qs("_", "123456")

// send with golang default HTTP client
res, err := http.DefaultClient.Do(r.ParseRequest())
defer res.Body.Close()
```

### Custom client

In case you did not know that golang default `http.Client` has **no timeout**.
use **rq/client** which has `180s` timeout by default

```go
import "github.com/ddo/rq"
import "github.com/ddo/rq/client"

r := rq.Post("https://httpbin.org/post")

// query
r.Qs("_", "123456")

// Form
r.Send("data", "data value")
r.Send("extra", "extra value")

// use default rq client
// true to tell #Send to read all the response boby when return
data, res, err := client.Send(r, true)
// no need to close res.Body
// read = false -> you need to call res.Body when done reading
```

### Headers

```go
r := rq.Post("https://httpbin.org/post")

r.Set("Content-Type", "application/json")
r.Set("User-Agent", "ddo/rq")
```

### Raw body

```go
r := rq.Post("https://httpbin.org/post")

r.SendRaw(strings.NewReader("raw data binary or json"))
```

## Client [![Doc][godoc-img]][godoc-client-url]

### Default

```go
// by default timeout = 3min
// no cookie jar
// and stops after 10 consecutive requests (10 redirects)
customClient := client.New(nil)
```

### Custom Options

```go
import "github.com/ddo/rq/client/jar"

cookieJar := jar.New()

// custom timeout = 10s and cookie jar
customClient := client.New(&Option{
    Timeout: time.Second * 10,
    jar: cookieJar,
})
```

### Default settings

```go
// set default User-Agent
defaultRq := rq.Get("")
defaultRq.Set("User-Agent", "github.com/ddo/rq")

customClient := client.New(&Option{
    DefaultRq: defaultRq,
})

// from now all the requests called via this customClient
// gonna have the User-Agent header = "github.com/ddo/rq"
// if User-Agent header in request is not set
```

## Redirect

-   Default `client` stops after 10 consecutive requests
-   Or you can use `client.NoRedirect` to disable redirect

```go
client.New(&Option{
    CheckRedirect: client.NoCheckRedirect,
})
```

## Cookies [![Doc][godoc-img]][godoc-jar-url]

```go
import "github.com/ddo/rq/client/jar"

cookieJar := jar.New()

customClient := client.New(&client.Option{
    Jar: cookieJar,
})

// get all cookies by hostname
cookies, err := cookieJar.Get("httpbin.org")

// get a cookie by hostname and name
cookie, err := cookieJar.GetByName("httpbin.org", "cookiename").

// set cookies
err := cookieJar.Set("httpbin.org", cookies)

// set a cookie
err := cookieJar.SetOne("httpbin.org", cookie)

// clear the cookie jar
err := cookieJar.Clear("httpbin.org")

// delete a cookie by it's name
err := cookieJar.Delete("httpbin.org", "cookiename")
```

## Debug

Set env `DLOG=*` to enable logger to see request activities

## TODO

List here [#1](https://github.com/ddo/rq/issues/1)