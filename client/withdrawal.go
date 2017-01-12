package client

import (
	"bytes"
	"fmt"
	"golang.org/x/net/context"
	"net/http"
	"net/url"
)

// CreateWithdrawalPath computes a request path to the create action of withdrawal.
func CreateWithdrawalPath() string {
	return fmt.Sprintf("/withdrawal")
}

// creates a withdrawal
func (c *Client) CreateWithdrawal(ctx context.Context, path string, payload *WithdrawalPayload, contentType string) (*http.Response, error) {
	req, err := c.NewCreateWithdrawalRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewCreateWithdrawalRequest create the request corresponding to the create action endpoint of the withdrawal resource.
func (c *Client) NewCreateWithdrawalRequest(ctx context.Context, path string, payload *WithdrawalPayload, contentType string) (*http.Request, error) {
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

// ShowWithdrawalPath computes a request path to the show action of withdrawal.
func ShowWithdrawalPath(id int) string {
	return fmt.Sprintf("/withdrawal/%v", id)
}

// shows a withdrawal
func (c *Client) ShowWithdrawal(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewShowWithdrawalRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewShowWithdrawalRequest create the request corresponding to the show action endpoint of the withdrawal resource.
func (c *Client) NewShowWithdrawalRequest(ctx context.Context, path string) (*http.Request, error) {
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
