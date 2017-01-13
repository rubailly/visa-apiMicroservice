package client

import (
	"bytes"
	"fmt"
	"golang.org/x/net/context"
	"net/http"
	"net/url"
)

// CreatePaymentPath computes a request path to the create action of payment.
func CreatePaymentPath() string {
	return fmt.Sprintf("/payment")
}

// creates a payment
func (c *Client) CreatePayment(ctx context.Context, path string, payload *PaymentPayload, contentType string) (*http.Response, error) {
	req, err := c.NewCreatePaymentRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewCreatePaymentRequest create the request corresponding to the create action endpoint of the payment resource.
func (c *Client) NewCreatePaymentRequest(ctx context.Context, path string, payload *PaymentPayload, contentType string) (*http.Request, error) {
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

// ShowPaymentPath computes a request path to the show action of payment.
func ShowPaymentPath(id int) string {
	return fmt.Sprintf("/payment/%v", id)
}

// shows a payment
func (c *Client) ShowPayment(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewShowPaymentRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewShowPaymentRequest create the request corresponding to the show action endpoint of the payment resource.
func (c *Client) NewShowPaymentRequest(ctx context.Context, path string) (*http.Request, error) {
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
