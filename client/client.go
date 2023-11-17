package client

import (
	"bytes"
	"crypto/tls"
	"io"
	"net/http"
	"time"
)

const (
	DefaultTimeout = time.Second * 20
)

type Config struct {
	Timeout   time.Duration
	TLSConfig *tls.Config
}

// HTTPClient wraps an http client and provides typical RESTful methods (GET, POST, PUT, DELETE, etc).
type HTTPClient struct {
	client *http.Client
}

// NewHTTPClient creates a new client.
func NewHTTPClient(c ...Config) *HTTPClient {
	// check if config provided
	var config Config
	if len(c) > 0 {
		config = c[0]
	}

	// set configuration defaults if not set
	if config.Timeout == 0 {
		config.Timeout = DefaultTimeout
	}

	// create http client
	client := &http.Client{
		Timeout: config.Timeout,
	}

	if config.TLSConfig != nil {
		client.Transport = &http.Transport{
			TLSClientConfig: config.TLSConfig,
		}
	}

	return &HTTPClient{
		client: client,
	}
}

// Get will retrieve a resource or a collection.
func (c *HTTPClient) Get(url string) (int, []byte, error) {
	return c.Do(http.MethodGet, url, nil)
}

// Post will create a resource.
func (c *HTTPClient) Post(url string, body []byte) (int, []byte, error) {
	return c.Do(http.MethodPost, url, body)
}

// Put will replace a resource or a collection.
func (c *HTTPClient) Put(url string, body []byte) (int, []byte, error) {
	return c.Do(http.MethodPut, url, body)
}

// Delete a resource.
func (c *HTTPClient) Delete(url string) (int, []byte, error) {
	return c.Do(http.MethodDelete, url, nil)
}

// Do sends a RESTful HTTP request. Allows any method supported by net/http. Returns the status code and response body.
func (c *HTTPClient) Do(method, url string, body []byte) (int, []byte, error) {
	return CallHTTP(c.client, method, url, body)
}

// CallHTTP sends a RESTful HTTP request. Allows any method supported by net/http. Returns the status code and response body.
func CallHTTP(client *http.Client, method, url string, body []byte) (int, []byte, error) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return 0, nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return 0, nil, err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, nil, err
	}

	return resp.StatusCode, respBody, nil
}
