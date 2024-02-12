package backend

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/khulnasoft/kengine-plugin-sdk-go/backend/proxy"
	"github.com/khulnasoft/kengine-plugin-sdk-go/backend/useragent"
	"github.com/khulnasoft/kengine-plugin-sdk-go/experimental/featuretoggles"
)

const (
	AppURL               = "GF_APP_URL"
	ConcurrentQueryCount = "GF_CONCURRENT_QUERY_COUNT"
)

type configKey struct{}

// KengineConfigFromContext returns Khulnasoft config from context.
func KengineConfigFromContext(ctx context.Context) *KengineCfg {
	v := ctx.Value(configKey{})
	if v == nil {
		return NewKengineCfg(nil)
	}

	cfg := v.(*KengineCfg)
	if cfg == nil {
		return NewKengineCfg(nil)
	}

	return cfg
}

// WithKengineConfig injects supplied Khulnasoft config into context.
func WithKengineConfig(ctx context.Context, cfg *KengineCfg) context.Context {
	ctx = context.WithValue(ctx, configKey{}, cfg)
	return ctx
}

type KengineCfg struct {
	config map[string]string
}

func NewKengineCfg(cfg map[string]string) *KengineCfg {
	return &KengineCfg{config: cfg}
}

func (c *KengineCfg) Get(key string) string {
	return c.config[key]
}

func (c *KengineCfg) FeatureToggles() FeatureToggles {
	features, exists := c.config[featuretoggles.EnabledFeatures]
	if !exists || features == "" {
		return FeatureToggles{}
	}

	fs := strings.Split(features, ",")
	enabledFeatures := make(map[string]struct{}, len(fs))
	for _, f := range fs {
		enabledFeatures[f] = struct{}{}
	}

	return FeatureToggles{
		enabled: enabledFeatures,
	}
}

func (c *KengineCfg) Equal(c2 *KengineCfg) bool {
	if c == nil && c2 == nil {
		return true
	}
	if c == nil || c2 == nil {
		return false
	}

	if len(c.config) != len(c2.config) {
		return false
	}
	for k, v1 := range c.config {
		if v2, ok := c2.config[k]; !ok || v1 != v2 {
			return false
		}
	}
	return true
}

type FeatureToggles struct {
	// enabled is a set-like map of feature flags that are enabled.
	enabled map[string]struct{}
}

// IsEnabled returns true if feature f is contained in ft.enabled.
func (ft FeatureToggles) IsEnabled(f string) bool {
	_, exists := ft.enabled[f]
	return exists
}

type Proxy struct {
	clientCfg *proxy.ClientCfg
}

func (c *KengineCfg) proxy() (Proxy, error) {
	if v, exists := c.config[proxy.PluginSecureSocksProxyEnabled]; exists && v == strconv.FormatBool(true) {
		var (
			allowInsecure = false
			err           error
		)

		if v := c.Get(proxy.PluginSecureSocksProxyAllowInsecure); v != "" {
			allowInsecure, err = strconv.ParseBool(c.Get(proxy.PluginSecureSocksProxyAllowInsecure))
			if err != nil {
				return Proxy{}, fmt.Errorf("parsing %s, value must be a boolean: %w", proxy.PluginSecureSocksProxyAllowInsecure, err)
			}
		}

		return Proxy{
			clientCfg: &proxy.ClientCfg{
				ClientCert:    c.Get(proxy.PluginSecureSocksProxyClientCert),
				ClientKey:     c.Get(proxy.PluginSecureSocksProxyClientKey),
				RootCA:        c.Get(proxy.PluginSecureSocksProxyRootCACert),
				ProxyAddress:  c.Get(proxy.PluginSecureSocksProxyProxyAddress),
				ServerName:    c.Get(proxy.PluginSecureSocksProxyServerName),
				AllowInsecure: allowInsecure,
			},
		}, nil
	}

	return Proxy{}, nil
}

func (c *KengineCfg) AppURL() (string, error) {
	url, ok := c.config[AppURL]
	if !ok {
		return "", fmt.Errorf("app URL not set in config. A more recent version of Khulnasoft may be required")
	}
	return url, nil
}

func (c *KengineCfg) ConcurrentQueryCount() (int, error) {
	count, ok := c.config[ConcurrentQueryCount]
	if !ok {
		return 0, fmt.Errorf("ConcurrentQueryCount not set in config")
	}
	i, err := strconv.Atoi(count)
	if err != nil {
		return 0, fmt.Errorf("ConcurrentQueryCount cannot be converted to integer")
	}
	return i, nil
}

type userAgentKey struct{}

// UserAgentFromContext returns user agent from context.
func UserAgentFromContext(ctx context.Context) *useragent.UserAgent {
	v := ctx.Value(userAgentKey{})
	if v == nil {
		return useragent.Empty()
	}

	ua := v.(*useragent.UserAgent)
	if ua == nil {
		return useragent.Empty()
	}

	return ua
}

// WithUserAgent injects supplied user agent into context.
func WithUserAgent(ctx context.Context, ua *useragent.UserAgent) context.Context {
	ctx = context.WithValue(ctx, userAgentKey{}, ua)
	return ctx
}
