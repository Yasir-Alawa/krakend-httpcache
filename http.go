// Package httpcache introduces an in-memory-cached http client into the KrakenD stack
package httpcache

import (
	"context"
	"net/http"
	"fmt"
	"github.com/Yasir-Alawa/httpcache"
	"github.com/luraproject/lura/config"
	"github.com/luraproject/lura/proxy"
	"github.com/luraproject/lura/transport/http/client"
)

// Namespace is the key to use to store and access the custom config data
const Namespace = "github.com/Yasir-Alawa/krakend-httpcache"

var (
	memTransport = httpcache.NewMemoryCacheTransport()
	memClient    = http.Client{Transport: memTransport}
)

// NewHTTPClient creates a HTTPClientFactory using an in-memory-cached http client
func NewHTTPClient(cfg *config.Backend) client.HTTPClientFactory {
	_, ok := cfg.ExtraConfig[Namespace]
	if !ok {
		return client.NewHTTPClient
	}
	return func(_ context.Context) *http.Client {
		return &memClient
	}
}

// BackendFactory returns a proxy.BackendFactory that creates backend proxies using
// an in-memory-cached http client
func BackendFactory(cfg *config.Backend) proxy.BackendFactory {
	return proxy.CustomHTTPProxyFactory(NewHTTPClient(cfg))
}
