package main

import (
	"ChamaconektVisa/app"
	"github.com/goadesign/goa"
)

// DepositController implements the deposit resource.
type DepositController struct {
	*goa.Controller
}

// NewDepositController creates a deposit controller.
func NewDepositController(service *goa.Service) *DepositController {
	return &DepositController{Controller: service.NewController("DepositController")}
}

// Create runs the create action.
func (c *DepositController) Create(ctx *app.CreateDepositContext) error {
	// DepositController_Create: start_implement

	// Put your logic here

	// DepositController_Create: end_implement
	return nil
}

// Show runs the show action.
func (c *DepositController) Show(ctx *app.ShowDepositContext) error {
	// DepositController_Show: start_implement

	// Put your logic here

	// DepositController_Show: end_implement
	res := &app.Deposit{}
	return ctx.OK(res)
}
