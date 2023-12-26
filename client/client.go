package client

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"time"

	"github.com/Rican7/retry"
	"github.com/Rican7/retry/backoff"
	"github.com/Rican7/retry/jitter"
	"github.com/Rican7/retry/strategy"
)

const (
	DefaultTimeout  = time.Second * 20
	MaxRetries      = 3
	BackoffDuration = 200 * time.Millisecond
)

var generator = rand.New(rand.NewSource(time.Now().UnixNano()))

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
func (c *HTTPClient) Get(url string, queryParams ...string) (int, []byte, error) {
	return c.Do(http.MethodGet, url, nil, queryParams...)
}

// Post will create a resource.
func (c *HTTPClient) Post(url string, body []byte, queryParams ...string) (int, []byte, error) {
	return c.Do(http.MethodPost, url, body, queryParams...)
}

// Put will replace a resource or a collection.
func (c *HTTPClient) Put(url string, body []byte, queryParams ...string) (int, []byte, error) {
	return c.Do(http.MethodPut, url, body, queryParams...)
}

// Delete a resource.
func (c *HTTPClient) Delete(url string, queryParams ...string) (int, []byte, error) {
	return c.Do(http.MethodDelete, url, nil, queryParams...)
}

// Do sends a RESTful HTTP request. Allows any method supported by net/http. Returns the status code and response body.
func (c *HTTPClient) Do(method, url string, body []byte, queryParams ...string) (int, []byte, error) {
	return CallHTTP(c.client, method, url, body, queryParams...)
}

// CallHTTP sends a RESTful HTTP request. Allows any method supported by net/http. Returns the status code and response body.
func CallHTTP(client *http.Client, method, url string, body []byte, queryParams ...string) (int, []byte, error) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return 0, nil, err
	}

	if len(queryParams) > 0 {
		if len(queryParams)%2 != 0 {
			return 0, nil, errors.New(`queryParams must have an even number of elements. Ex: "Name1", "Value1", "Name2", "Value2"`)
		}
		query := req.URL.Query()
		for i := 0; i < len(queryParams)-1; i += 2 {
			query.Add(queryParams[i], queryParams[i+1])
		}
		req.URL.RawQuery = query.Encode()
	}

	// print request
	if body != nil {
		var prettyReq bytes.Buffer
		json.Indent(&prettyReq, body, "", "    ")
		fmt.Printf("%s %s\n%s\n\n", req.Method, req.URL, prettyReq.String())
	} else {
		fmt.Printf("%s %s\n\n", req.Method, req.URL)
	}

	var resp *http.Response
	var respBody []byte
	retryErr := retry.Retry(
		func(attempt uint) error {
			if attempt > 1 {
				fmt.Printf("Retrying operation, attempt %v of %v\n", attempt-1, MaxRetries)
			}

			resp, err = client.Do(req)
			if err != nil {
				return nil
			}
			defer resp.Body.Close()

			respBody, err = io.ReadAll(resp.Body)
			if err != nil {
				return nil
			}

			// determine if operation should be retried
			if (resp.StatusCode >= 500 && resp.StatusCode <= 599) || resp.StatusCode == http.StatusRequestTimeout {
				if attempt <= MaxRetries {
					fmt.Printf("Retrying http operation after receiving retryable status code: %v\n", resp.StatusCode)
				} else {
					fmt.Printf("Max retries reached - operation failed with status code: %v\n\n", resp.StatusCode)
				}
				return errors.New("retrying http operation after receiving retryable status code")
			}

			return nil
		},
		strategy.Limit(MaxRetries+1), // strategy.Limit is the total number of attempts, so original attempt + max retries
		strategy.BackoffWithJitter(
			backoff.BinaryExponential(BackoffDuration),
			jitter.Deviation(generator, 0.5),
		),
	)

	// check for any errors that may have occurred during retry
	if err != nil && retryErr == nil {
		return 0, respBody, err
	}

	// print response
	if len(respBody) > 0 {
		var prettyResp bytes.Buffer
		json.Indent(&prettyResp, respBody, "", "    ")
		fmt.Printf("%s\n%s\n", resp.Status, prettyResp.String())
	} else {
		fmt.Printf("%s\n", resp.Status)
	}

	return resp.StatusCode, respBody, nil
}
