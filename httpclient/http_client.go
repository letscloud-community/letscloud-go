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
	defaultTimeout = time.Second * 10
)

type httpClient struct {
	apiKey string
	httpcl *http.Client
}

func (h *httpClient) APIKey() string {
	return h.apiKey
}

func (h *httpClient) NewRequest(method, url string, data interface{}) (*http.Request, error) {
	if h.apiKey == "" {
		return nil, errors.New("No API Key found. Provide your api-key!")
	}

	b, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, url, bytes.NewReader(b))
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

	if resp.StatusCode != http.StatusOK {
		return ioutil.ReadAll(resp.Body)
	}

	return ioutil.ReadAll(resp.Body)
}

func (h *httpClient) Do(req *http.Request) (*http.Response, error) {
	return h.httpcl.Do(req)
}

func (h *httpClient) SetTimeout(d time.Duration) {
	h.httpcl.Timeout = d
}

func (h *httpClient) SetAPIKey(t string) {
	h.apiKey = t
}

func NewHttpClient(ak string) *httpClient {
	cl := http.Client{
		Timeout: defaultTimeout,
	}

	return &httpClient{httpcl: &cl, apiKey: ak}
}
