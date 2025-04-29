package httpclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	defaultTimeout = time.Second * 60
	defaultBaseURL = "https://core.letscloud.io/api"
)

type httpClient struct {
	apiKey  string
	baseURL string
	httpcl  *http.Client
}

func (h *httpClient) APIKey() string {
	return h.apiKey
}

func (h *httpClient) SetAPIKey(t string) {
	h.apiKey = t
}

func (h *httpClient) SetBaseURL(url string) {
	h.baseURL = url
}

func (h *httpClient) SetTimeout(d time.Duration) {
	h.httpcl.Timeout = d
}

func (h *httpClient) NewRequest(method, endpoint string, data interface{}) (*http.Request, error) {
	if h.apiKey == "" {
		return nil, errors.New("no api key found. provide your api-key")
	}

	if h.baseURL == "" {
		h.baseURL = defaultBaseURL
	}

	b, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, h.baseURL+endpoint, bytes.NewReader(b))
	if err != nil {
		return nil, err
	}

	req.Header.Add("api-token", h.apiKey)

	if data != nil {
		req.Header.Add("Content-Type", "application/json")
	}

	return req, nil
}

func (h *httpClient) SendRequest(req *http.Request) ([]byte, error) {
	resp, err := h.httpcl.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusUnauthorized {
		return nil, errors.New("401 unauthorized: please check your api key")
	}

	if resp.StatusCode != http.StatusOK {
		return ioutil.ReadAll(resp.Body)
	}

	return ioutil.ReadAll(resp.Body)
}

func (h *httpClient) Do(req *http.Request) (*http.Response, error) {
	return h.httpcl.Do(req)
}

// NewHttpClient creates a new instance of httpClient
func NewHttpClient(apiKey string) *httpClient {
	return &httpClient{
		apiKey:  apiKey,
		baseURL: defaultBaseURL,
		httpcl:  &http.Client{Timeout: defaultTimeout},
	}
}
