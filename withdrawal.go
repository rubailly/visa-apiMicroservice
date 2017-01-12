package main

import (
	"ChamaconektVisa/app"
	"github.com/goadesign/goa"
)

// WithdrawalController implements the withdrawal resource.
type WithdrawalController struct {
	*goa.Controller
}

// NewWithdrawalController creates a withdrawal controller.
func NewWithdrawalController(service *goa.Service) *WithdrawalController {
	return &WithdrawalController{Controller: service.NewController("WithdrawalController")}
}

// Create runs the create action.
func (c *WithdrawalController) Create(ctx *app.CreateWithdrawalContext) error {
	// WithdrawalController_Create: start_implement

	// Put your logic here

	// WithdrawalController_Create: end_implement
	return nil
}

// Show runs the show action.
func (c *WithdrawalController) Show(ctx *app.ShowWithdrawalContext) error {
	// WithdrawalController_Show: start_implement

	// Put your logic here

	// WithdrawalController_Show: end_implement
	res := &app.Withdrawal{}
	return ctx.OK(res)
}
