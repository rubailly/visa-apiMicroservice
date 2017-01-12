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

// CreateDepositsContext provides the deposits create action context.
type CreateDepositsContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Payload *DepositsPayload
}

// NewCreateDepositsContext parses the incoming request URL and body, performs validations and creates the
// context used by the deposits controller create action.
func NewCreateDepositsContext(ctx context.Context, service *goa.Service) (*CreateDepositsContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := CreateDepositsContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// Created sends a HTTP response with status code 201.
func (ctx *CreateDepositsContext) Created() error {
	ctx.ResponseData.WriteHeader(201)
	return nil
}

// ShowDepositsContext provides the deposits show action context.
type ShowDepositsContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	ID int
}

// NewShowDepositsContext parses the incoming request URL and body, performs validations and creates the
// context used by the deposits controller show action.
func NewShowDepositsContext(ctx context.Context, service *goa.Service) (*ShowDepositsContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := ShowDepositsContext{Context: ctx, ResponseData: resp, RequestData: req}
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
func (ctx *ShowDepositsContext) OK(r *Deposits) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.depositsmedia+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}
