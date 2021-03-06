package letscloud

import (
	"net/http"
	"time"

	"github.com/letscloud-community/letscloud-go/httpclient"
)

const (
	baseURL = "https://core.letscloud.io/api"
)

// LetsCloud represents a wrapper client for our LetsCloud API
type LetsCloud struct {
	debug     bool
	requester Requester
}

//Requester defines the API that will be used for sending HTTP Requests to the letscloud API
type Requester interface {
	NewRequest(method, url string, data interface{}) (*http.Request, error)
	SendRequest(req *http.Request) ([]byte, error)
	SetTimeout(d time.Duration)
	SetAPIKey(t string)
	APIKey() string
}

//SetTimeout sets timeout for http client
func (c *LetsCloud) SetTimeout(d time.Duration) error {
	if d < 0 {
		return ErrInvalidTimeout
	}

	c.requester.SetTimeout(d)

	return nil
}

//APIKey get the API Key
func (c LetsCloud) APIKey() string {
	return c.requester.APIKey()
}

//SetAPIKey sets the API Key
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

//New instantiate a new LetsCloud Client
func New(apiKey string) (*LetsCloud, error) {
	cl := httpclient.NewHttpClient(apiKey)

	return &LetsCloud{requester: cl}, nil
}
