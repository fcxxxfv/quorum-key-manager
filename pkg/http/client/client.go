package httpclient

import (
	"net/http"

	"github.com/ConsenSysQuorum/quorum-key-manager/pkg/http/request"
	"github.com/ConsenSysQuorum/quorum-key-manager/pkg/http/response"
	"github.com/ConsenSysQuorum/quorum-key-manager/pkg/http/transport"
)

// Client is an HTTP client
type Client interface {
	// Do sends an HTTP request and returns an HTTP response, following
	// policy (such as redirects, cookies, auth)
	Do(*http.Request) (*http.Response, error)

	// CloseIdleConnections closes any connections on its Transport which
	// were previously connected from previous requests but are now
	// sitting idle in a "keep-alive" state. It does not interrupt any
	// connections currently in use.
	CloseIdleConnections()
}

// New creates a new HTTP client
func New(
	cfg *Config,
	trnsprt http.RoundTripper,
	preparer request.Preparer,
	modifier response.Modifier,
) (Client, error) {
	cfg = cfg.Copy().SetDefault()

	var err error
	if trnsprt == nil {
		trnsprt, err = transport.New(cfg.Transport)
		if err != nil {
			return nil, err
		}
	}

	c := &http.Client{
		Transport: trnsprt,
		Timeout:   cfg.Timeout.Duration,
	}

	if preparer == nil {
		preparer, err = request.Proxy(cfg.Request)
		if err != nil {
			return nil, err
		}
	}

	if modifier == nil {
		modifier = response.Proxy(cfg.Response)
	}

	return WithModifier(modifier)(WithPreparer(preparer)(c)), nil
}
