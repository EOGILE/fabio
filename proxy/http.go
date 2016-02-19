package proxy

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

func newHTTPProxy(t *url.URL, tr http.RoundTripper, prefix string, stripPrefix bool) http.Handler {
	rp := httputil.NewSingleHostReverseProxy(t)
	rp.Transport = tr
	if stripPrefix {
		return http.StripPrefix(prefix, rp)
	} else {
		return rp
	}
}
