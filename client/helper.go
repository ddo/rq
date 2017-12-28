package client

import (
	"net/url"
	"strconv"
	"strings"
	"time"
)

// helper
func humanizeNano(n time.Duration) string {
	var suffix string

	switch {
	case n > 1e9:
		n /= 1e9
		suffix = "s"
	case n > 1e6:
		n /= 1e6
		suffix = "ms"
	case n > 1e3:
		n /= 1e3
		suffix = "us"
	default:
		suffix = "ns"
	}

	return strconv.Itoa(int(n)) + suffix
}

func parseHostname(hostname string) (u *url.URL, err error) {
	// prefix scheme
	if !strings.HasPrefix(hostname, "http://") &&
		!strings.HasPrefix(hostname, "https://") {
		hostname = "https://" + hostname
	}
	log.Info("hostname:", hostname)

	u, err = url.Parse(hostname)
	if err != nil {
		log.Error("URL", err)
		return
	}

	return
}
