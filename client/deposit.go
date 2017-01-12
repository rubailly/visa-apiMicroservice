package client

import (
	"bytes"
	"fmt"
	"golang.org/x/net/context"
	"net/http"
	"net/url"
)

// CreateDepositPath computes a request path to the create action of deposit.
func CreateDepositPath() string {
	return fmt.Sprintf("/deposit")
}

// creates a deposit
func (c *Client) CreateDeposit(ctx context.Context, path string, payload *DepositPayload, contentType string) (*http.Response, error) {
	req, err := c.NewCreateDepositRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewCreateDepositRequest create the request corresponding to the create action endpoint of the deposit resource.
func (c *Client) NewCreateDepositRequest(ctx context.Context, path string, payload *DepositPayload, contentType string) (*http.Request, error) {
	var body bytes.Buffer
	if contentType == "" {
		contentType = "*/*" // Use default encoder
	}
	err := c.Encoder.Encode(payload, &body, contentType)
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("POST", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	if contentType != "*/*" {
		header.Set("Content-Type", contentType)
	}
	return req, nil
}

// ShowDepositPath computes a request path to the show action of deposit.
func ShowDepositPath(id int) string {
	return fmt.Sprintf("/deposit/%v", id)
}

// shows a deposit
func (c *Client) ShowDeposit(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewShowDepositRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewShowDepositRequest create the request corresponding to the show action endpoint of the deposit resource.
func (c *Client) NewShowDepositRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
