package gladia

import (
	"net/http"
	"time"
)

const defaultBaseURL = "https://api.gladia.io/"
const defaultTimeout = 30 * time.Second
const gladiaHeaderKey = "x-gladia-key"

// HTTPDoer represents the minimal interface needed for HTTP operations
type HTTPDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client is the client for interacting with the Gladia API
type Client struct {
	APIKey     string
	BaseURL    string
	httpClient HTTPDoer
}

// NewClient creates a new Gladia API client
func NewClient(apiKey string, opts ...ClientOption) *Client {
	client := &Client{
		APIKey:  apiKey,
		BaseURL: defaultBaseURL,
		httpClient: &http.Client{
			Timeout: defaultTimeout,
		},
	}

	// Apply options
	for _, opt := range opts {
		opt(client)
	}

	return client
}

// ClientOption is a function that configures a Client
type ClientOption func(*Client)

// WithBaseURL sets the base URL for the client
func WithBaseURL(url string) ClientOption {
	return func(c *Client) {
		c.BaseURL = url
	}
}

// WithHTTPClient sets the HTTP client for the client
func WithHTTPClient(httpClient HTTPDoer) ClientOption {
	return func(c *Client) {
		c.httpClient = httpClient
	}
}

// WithTimeout sets the timeout duration for the HTTP client
func WithTimeout(timeout time.Duration) ClientOption {
	return func(c *Client) {
		if httpClient, ok := c.httpClient.(*http.Client); ok {
			httpClient.Timeout = timeout
		}
	}
}
