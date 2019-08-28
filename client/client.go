package client

import (
	"net/http"
	"time"

	"gopkg.in/ddo/go-dlog.v2"

	"github.com/ddo/rq"
	"github.com/ddo/rq/client/jar"
)

const (
	defaultTimeout   = 3 * time.Minute
	defaultRedirects = 10

	defaultUserAgent = "github.com/ddo/rq"
)

var log = dlog.New("rq:client", nil)

// Client contains stdlib http client and other custom client settings
type Client struct {
	httpClient *http.Client

	defaultRq *rq.Rq // default url, method, qs, form and headers if they are nil
}

// Option contains client settings
type Option struct {
	Timeout time.Duration // if 0, default timeout will be applied

	Jar *jar.Jar // if Jar is nil, there no cookie jar

	Proxy     string // Proxy option is the sugar syntax for Transport.Proxy
	Transport http.RoundTripper

	CheckRedirect func(req *http.Request, via []*http.Request) error

	DefaultRq *rq.Rq
}

// New returns new client which init with provided options
func New(opt *Option) *Client {
	if opt == nil {
		opt = &Option{}
	}

	timeout := opt.Timeout
	var jar http.CookieJar
	proxy := opt.Proxy
	transport := opt.Transport
	checkRedirect := opt.CheckRedirect
	defaultRq := opt.DefaultRq

	if opt.Timeout == 0 {
		timeout = defaultTimeout
	}

	if opt.Jar != nil {
		jar = opt.Jar.GetHTTPJar()
	}

	if checkRedirect == nil {
		checkRedirect = DefaultCheckRedirect
	}

	log.Info("timeout:", timeout)
	log.Info("jar:", jar)
	log.Info("proxy:", proxy)
	log.Info("transport:", transport)
	log.Info("checkRedirect:", checkRedirect)

	client := &Client{
		httpClient: &http.Client{
			Timeout:       timeout,
			Jar:           jar,
			Transport:     transport,
			CheckRedirect: checkRedirect,
		},
		defaultRq: defaultRq,
	}

	if proxy != "" {
		err := client.SetProxy(proxy)
		if err != nil {
			return nil
		}
	}

	return client
}

// DefaultClient has default timeout and stdlib default transport
// no cookie management
var DefaultClient = New(nil)

// ApplyDefaultRq overrides Rq properties with default value if key is not set
func ApplyDefaultRq(defaultRq, _rq *rq.Rq) (newRq *rq.Rq) {
	newRq = &rq.Rq{}

	// copy
	*newRq = *_rq

	// return as is _rq if no defaultRq
	if defaultRq == nil {
		return
	}

	// url
	if newRq.URL == "" {
		newRq.URL = defaultRq.URL
	}

	// query
	for k, v := range defaultRq.Query {
		if _, ok := newRq.Query[k]; !ok {
			newRq.Query[k] = v
		}
	}

	// form
	for k, v := range defaultRq.Form {
		if _, ok := newRq.Form[k]; !ok {
			newRq.Form[k] = v
		}
	}

	// header
	for k, v := range defaultRq.Header {
		if _, ok := newRq.Header[k]; !ok {
			newRq.Header[k] = v
		}
	}
	return
}
