package client

import (
	"net/http"
	"net/http/cookiejar"
	"time"

	"gopkg.in/ddo/go-dlog.v2"

	"github.com/ddo/rq"
)

const (
	defaultTimeout   = 3 * time.Minute
	defaultRedirects = 10
)

var log = dlog.New("rq:client", nil)

// Client contains stdlib http client and other custom client settings
type Client struct {
	httpClient *http.Client

	defaultRq *rq.Rq // default url, method, qs, form and headers if they are nil
}

// Option contains client settings
type Option struct {
	Timeout   time.Duration
	NoTimeout bool // if NoTimeout is false Timeout will be set as default

	Jar      http.CookieJar
	NoCookie bool // if NoCookie is true Jar will be skip

	Proxy     string // Proxy option is the sugar syntax for Transport.Proxy
	Transport http.RoundTripper

	CheckRedirect func(req *http.Request, via []*http.Request) error

	DefaultRq *rq.Rq
}

// New returns new client which init with provided options
// cookie management is enable by default
func New(opt *Option) *Client {
	if opt == nil {
		opt = &Option{}
	}

	timeout := opt.Timeout
	jar := opt.Jar
	proxy := opt.Proxy
	transport := opt.Transport
	checkRedirect := opt.CheckRedirect
	defaultRq := opt.DefaultRq

	if opt.Timeout == 0 && !opt.NoTimeout {
		timeout = defaultTimeout
	}
	if opt.NoTimeout {
		timeout = 0
	}

	if opt.Jar == nil && !opt.NoCookie {
		jar, _ = cookiejar.New(nil)
	}
	if opt.NoCookie {
		jar = nil
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
var DefaultClient = New(&Option{NoCookie: true})

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
