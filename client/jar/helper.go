package jar

import (
	"net/url"
	"strings"
)

func parseHostname(hostname string) (u *url.URL, err error) {
	// prefix scheme
	if !strings.HasPrefix(hostname, "http://") &&
		!strings.HasPrefix(hostname, "https://") {
		hostname = "https://" + hostname
	}
	// log.Info("hostname:", hostname)

	u, err = url.Parse(hostname)
	if err != nil {
		log.Error("URL", err)
		return
	}

	return
}
