package letscloud

import (
	"log"
	"net/http"
	"time"

	"github.com/letscloud-community/letscloud-go/httpclient"
)

// Version of the SDK
const Version = "1.2.0"

// LetsCloud represents a wrapper client for our LetsCloud API
type LetsCloud struct {
	debug     bool
	requester Requester
}

// Requester defines the API that will be used for sending HTTP Requests to the letscloud API
type Requester interface {
	NewRequest(method, url string, data interface{}) (*http.Request, error)
	SendRequest(req *http.Request) ([]byte, error)
	SetTimeout(d time.Duration)
	SetAPIKey(t string)
	SetBaseURL(url string)
	APIKey() string
}

// Option is a function that can be used to set options for the LetsCloud client
type Option func(*LetsCloud)

// WithTimeout sets the timeout for the HTTP client
func WithTimeout(timeout time.Duration) Option {
	return func(lc *LetsCloud) {
		lc.requester.SetTimeout(timeout)
	}
}

// WithBaseURL sets the base URL for the HTTP client
func WithBaseURL(baseURL string) Option {
	return func(lc *LetsCloud) {
		lc.requester.SetBaseURL(baseURL)
	}
}

// WithDebug enables or disables debug mode
func WithDebug(debug bool) Option {
	return func(lc *LetsCloud) {
		lc.debug = debug
	}
}

// New creates a new instance of LetsCloud with the provided API key and options
func New(apiKey string, opts ...Option) (*LetsCloud, error) {
	cl := httpclient.NewHttpClient(apiKey)
	lc := &LetsCloud{requester: cl}

	// Apply options
	for _, opt := range opts {
		opt(lc)
	}

	return lc, nil
}

// SetTimeout sets timeout for http client
func (c *LetsCloud) SetTimeout(d time.Duration) error {
	if d < 0 {
		return ErrInvalidTimeout
	}

	c.requester.SetTimeout(d)

	return nil
}

// APIKey get the API Key
func (c LetsCloud) APIKey() string {
	return c.requester.APIKey()
}

// SetAPIKey sets the API Key
func (c *LetsCloud) SetAPIKey(ak string) error {
	if ak == "" {
		return ErrInvalidToken
	}

	if c.requester == nil {
		c.requester = httpclient.NewHttpClient(ak)
	}

	c.requester.SetAPIKey(ak)

	return nil
}

// debugLog logs debug messages if debug mode is enabled
func (c *LetsCloud) debugLog(message string) {
	if c.debug {
		log.Println("[DEBUG]", message)
	}
}
