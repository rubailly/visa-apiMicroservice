package client

import (
	"bytes"
	"fmt"
	"golang.org/x/net/context"
	"net/http"
	"net/url"
)

// CreateDepositsPath computes a request path to the create action of deposits.
func CreateDepositsPath() string {
	return fmt.Sprintf("/deposits")
}

// creates a deposit
func (c *Client) CreateDeposits(ctx context.Context, path string, payload *DepositsPayload, contentType string) (*http.Response, error) {
	req, err := c.NewCreateDepositsRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewCreateDepositsRequest create the request corresponding to the create action endpoint of the deposits resource.
func (c *Client) NewCreateDepositsRequest(ctx context.Context, path string, payload *DepositsPayload, contentType string) (*http.Request, error) {
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

// ShowDepositsPath computes a request path to the show action of deposits.
func ShowDepositsPath(id int) string {
	return fmt.Sprintf("/deposits/%v", id)
}

// shows a deposit
func (c *Client) ShowDeposits(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewShowDepositsRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewShowDepositsRequest create the request corresponding to the show action endpoint of the deposits resource.
func (c *Client) NewShowDepositsRequest(ctx context.Context, path string) (*http.Request, error) {
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
