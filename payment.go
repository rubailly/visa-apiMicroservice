package main

import (
	"ChamaconektVisa/app"
	"github.com/goadesign/goa"
)

// PaymentController implements the payment resource.
type PaymentController struct {
	*goa.Controller
}

// NewPaymentController creates a payment controller.
func NewPaymentController(service *goa.Service) *PaymentController {
	return &PaymentController{Controller: service.NewController("PaymentController")}
}

// Create runs the create action.
func (c *PaymentController) Create(ctx *app.CreatePaymentContext) error {
	// PaymentController_Create: start_implement

	// Put your logic here

	// PaymentController_Create: end_implement
	return nil
}

// Show runs the show action.
func (c *PaymentController) Show(ctx *app.ShowPaymentContext) error {
	// PaymentController_Show: start_implement

	// Put your logic here

	// PaymentController_Show: end_implement
	res := &app.Payment{}
	return ctx.OK(res)
}
