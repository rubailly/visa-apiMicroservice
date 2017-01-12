//************************************************************************//
// API "mVisa": Application Contexts
//
// Generated with goagen v1.0.0, command line:
// $ goagen
// --design=ChamaconektVisa/design
// --out=$(GOPATH)/src/ChamaconektVisa
// --version=v1.0.0
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package app

import (
	"github.com/goadesign/goa"
	"golang.org/x/net/context"
	"strconv"
)

// CreateDepositContext provides the deposit create action context.
type CreateDepositContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Payload *DepositPayload
}

// NewCreateDepositContext parses the incoming request URL and body, performs validations and creates the
// context used by the deposit controller create action.
func NewCreateDepositContext(ctx context.Context, service *goa.Service) (*CreateDepositContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := CreateDepositContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// Created sends a HTTP response with status code 201.
func (ctx *CreateDepositContext) Created() error {
	ctx.ResponseData.WriteHeader(201)
	return nil
}

// ShowDepositContext provides the deposit show action context.
type ShowDepositContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	ID int
}

// NewShowDepositContext parses the incoming request URL and body, performs validations and creates the
// context used by the deposit controller show action.
func NewShowDepositContext(ctx context.Context, service *goa.Service) (*ShowDepositContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := ShowDepositContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramID := req.Params["id"]
	if len(paramID) > 0 {
		rawID := paramID[0]
		if id, err2 := strconv.Atoi(rawID); err2 == nil {
			rctx.ID = id
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("id", rawID, "integer"))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowDepositContext) OK(r *Deposit) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.depositmedia+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// CreateWithdrawalContext provides the withdrawal create action context.
type CreateWithdrawalContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Payload *WithdrawalPayload
}

// NewCreateWithdrawalContext parses the incoming request URL and body, performs validations and creates the
// context used by the withdrawal controller create action.
func NewCreateWithdrawalContext(ctx context.Context, service *goa.Service) (*CreateWithdrawalContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := CreateWithdrawalContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// Created sends a HTTP response with status code 201.
func (ctx *CreateWithdrawalContext) Created() error {
	ctx.ResponseData.WriteHeader(201)
	return nil
}

// ShowWithdrawalContext provides the withdrawal show action context.
type ShowWithdrawalContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	ID int
}

// NewShowWithdrawalContext parses the incoming request URL and body, performs validations and creates the
// context used by the withdrawal controller show action.
func NewShowWithdrawalContext(ctx context.Context, service *goa.Service) (*ShowWithdrawalContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := ShowWithdrawalContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramID := req.Params["id"]
	if len(paramID) > 0 {
		rawID := paramID[0]
		if id, err2 := strconv.Atoi(rawID); err2 == nil {
			rctx.ID = id
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("id", rawID, "integer"))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowWithdrawalContext) OK(r *Withdrawal) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.withdrawalmedia+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}
