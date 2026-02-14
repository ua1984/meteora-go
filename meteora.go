// Package meteora provides a Go client library for the Meteora REST APIs.
package meteora

import (
	"net/http"

	"github.com/ua1984/meteora-go/dammv1"
	"github.com/ua1984/meteora-go/dammv2"
	"github.com/ua1984/meteora-go/dlmm"
	"github.com/ua1984/meteora-go/dynamicvault"
	"github.com/ua1984/meteora-go/internal/httpclient"
	"github.com/ua1984/meteora-go/stake2earn"
)

const (
	defaultDLMMBaseURL         = "https://dlmm.datapi.meteora.ag"
	defaultDLMMLegacyBaseURL   = "https://dlmm-api.meteora.ag"
	defaultDAMMv2BaseURL       = "https://damm-v2.datapi.meteora.ag"
	defaultDAMMv1BaseURL       = "https://amm-v2.meteora.ag"
	defaultStake2EarnBaseURL   = "https://stake-for-fee-api.meteora.ag"
	defaultDynamicVaultBaseURL = "https://merv2-api.meteora.ag"
)

// Client provides access to all Meteora API services.
type Client struct {
	DLMM         *dlmm.Client
	DAMMv2       *dammv2.Client
	DAMMv1       *dammv1.Client
	Stake2Earn   *stake2earn.Client
	DynamicVault *dynamicvault.Client
}

// Option configures the Client.
type Option func(*options)

type options struct {
	httpClient          *http.Client
	dlmmBaseURL         string
	dlmmLegacyBaseURL   string
	dammv2BaseURL       string
	dammv1BaseURL       string
	stake2earnBaseURL   string
	dynamicVaultBaseURL string
}

// WithHTTPClient sets a custom http.Client for all API requests.
func WithHTTPClient(c *http.Client) Option {
	return func(o *options) { o.httpClient = c }
}

// WithDLMMBaseURL overrides the DLMM datapi base URL.
func WithDLMMBaseURL(u string) Option {
	return func(o *options) { o.dlmmBaseURL = u }
}

// WithDLMMLegacyBaseURL overrides the DLMM legacy API base URL.
func WithDLMMLegacyBaseURL(u string) Option {
	return func(o *options) { o.dlmmLegacyBaseURL = u }
}

// WithDAMMv2BaseURL overrides the DAMM v2 datapi base URL.
func WithDAMMv2BaseURL(u string) Option {
	return func(o *options) { o.dammv2BaseURL = u }
}

// WithDAMMv1BaseURL overrides the DAMM v1 API base URL.
func WithDAMMv1BaseURL(u string) Option {
	return func(o *options) { o.dammv1BaseURL = u }
}

// WithStake2EarnBaseURL overrides the Stake2Earn API base URL.
func WithStake2EarnBaseURL(u string) Option {
	return func(o *options) { o.stake2earnBaseURL = u }
}

// WithDynamicVaultBaseURL overrides the Dynamic Vault API base URL.
func WithDynamicVaultBaseURL(u string) Option {
	return func(o *options) { o.dynamicVaultBaseURL = u }
}

// New creates a new Meteora API client with the given options.
func New(opts ...Option) *Client {
	o := &options{
		dlmmBaseURL:         defaultDLMMBaseURL,
		dlmmLegacyBaseURL:   defaultDLMMLegacyBaseURL,
		dammv2BaseURL:       defaultDAMMv2BaseURL,
		dammv1BaseURL:       defaultDAMMv1BaseURL,
		stake2earnBaseURL:   defaultStake2EarnBaseURL,
		dynamicVaultBaseURL: defaultDynamicVaultBaseURL,
	}
	for _, opt := range opts {
		opt(o)
	}

	hc := o.httpClient

	return &Client{
		DLMM:         dlmm.NewClient(httpclient.New(o.dlmmBaseURL, hc), httpclient.New(o.dlmmLegacyBaseURL, hc)),
		DAMMv2:       dammv2.NewClient(httpclient.New(o.dammv2BaseURL, hc)),
		DAMMv1:       dammv1.NewClient(httpclient.New(o.dammv1BaseURL, hc)),
		Stake2Earn:   stake2earn.NewClient(httpclient.New(o.stake2earnBaseURL, hc)),
		DynamicVault: dynamicvault.NewClient(httpclient.New(o.dynamicVaultBaseURL, hc)),
	}
}
