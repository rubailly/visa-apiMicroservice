//************************************************************************//
// API "ChamaconektVisa": Application Controllers
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
	"net/http"
)

// initService sets up the service encoders, decoders and mux.
func initService(service *goa.Service) {
	// Setup encoders and decoders
	service.Encoder.Register(goa.NewJSONEncoder, "application/json")
	service.Encoder.Register(goa.NewGobEncoder, "application/gob", "application/x-gob")
	service.Encoder.Register(goa.NewXMLEncoder, "application/xml")
	service.Decoder.Register(goa.NewJSONDecoder, "application/json")
	service.Decoder.Register(goa.NewGobDecoder, "application/gob", "application/x-gob")
	service.Decoder.Register(goa.NewXMLDecoder, "application/xml")

	// Setup default encoder and decoder
	service.Encoder.Register(goa.NewJSONEncoder, "*/*")
	service.Decoder.Register(goa.NewJSONDecoder, "*/*")
}

// DepositController is the controller interface for the Deposit actions.
type DepositController interface {
	goa.Muxer
	Create(*CreateDepositContext) error
	Show(*ShowDepositContext) error
}

// MountDepositController "mounts" a Deposit resource controller on the given service.
func MountDepositController(service *goa.Service, ctrl DepositController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewCreateDepositContext(ctx, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*DepositPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Create(rctx)
	}
	service.Mux.Handle("POST", "/deposit", ctrl.MuxHandler("Create", h, unmarshalCreateDepositPayload))
	service.LogInfo("mount", "ctrl", "Deposit", "action", "Create", "route", "POST /deposit")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewShowDepositContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.Show(rctx)
	}
	service.Mux.Handle("GET", "/deposit/:id", ctrl.MuxHandler("Show", h, nil))
	service.LogInfo("mount", "ctrl", "Deposit", "action", "Show", "route", "GET /deposit/:id")
}

// unmarshalCreateDepositPayload unmarshals the request body into the context request data Payload field.
func unmarshalCreateDepositPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &depositPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// PaymentController is the controller interface for the Payment actions.
type PaymentController interface {
	goa.Muxer
	Create(*CreatePaymentContext) error
	Show(*ShowPaymentContext) error
}

// MountPaymentController "mounts" a Payment resource controller on the given service.
func MountPaymentController(service *goa.Service, ctrl PaymentController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewCreatePaymentContext(ctx, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*PaymentPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Create(rctx)
	}
	service.Mux.Handle("POST", "/payment", ctrl.MuxHandler("Create", h, unmarshalCreatePaymentPayload))
	service.LogInfo("mount", "ctrl", "Payment", "action", "Create", "route", "POST /payment")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewShowPaymentContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.Show(rctx)
	}
	service.Mux.Handle("GET", "/payment/:id", ctrl.MuxHandler("Show", h, nil))
	service.LogInfo("mount", "ctrl", "Payment", "action", "Show", "route", "GET /payment/:id")
}

// unmarshalCreatePaymentPayload unmarshals the request body into the context request data Payload field.
func unmarshalCreatePaymentPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &paymentPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// WithdrawalController is the controller interface for the Withdrawal actions.
type WithdrawalController interface {
	goa.Muxer
	Create(*CreateWithdrawalContext) error
	Show(*ShowWithdrawalContext) error
}

// MountWithdrawalController "mounts" a Withdrawal resource controller on the given service.
func MountWithdrawalController(service *goa.Service, ctrl WithdrawalController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewCreateWithdrawalContext(ctx, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*WithdrawalPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Create(rctx)
	}
	service.Mux.Handle("POST", "/withdrawal", ctrl.MuxHandler("Create", h, unmarshalCreateWithdrawalPayload))
	service.LogInfo("mount", "ctrl", "Withdrawal", "action", "Create", "route", "POST /withdrawal")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewShowWithdrawalContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.Show(rctx)
	}
	service.Mux.Handle("GET", "/withdrawal/:id", ctrl.MuxHandler("Show", h, nil))
	service.LogInfo("mount", "ctrl", "Withdrawal", "action", "Show", "route", "GET /withdrawal/:id")
}

// unmarshalCreateWithdrawalPayload unmarshals the request body into the context request data Payload field.
func unmarshalCreateWithdrawalPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &withdrawalPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}
