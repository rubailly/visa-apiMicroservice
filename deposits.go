package main

import (
	"ChamaconektVisa/app"
	"github.com/goadesign/goa"
)

// DepositsController implements the deposits resource.
type DepositsController struct {
	*goa.Controller
}

// NewDepositsController creates a deposits controller.
func NewDepositsController(service *goa.Service) *DepositsController {
	return &DepositsController{Controller: service.NewController("DepositsController")}
}

// Create runs the create action.
func (c *DepositsController) Create(ctx *app.CreateDepositsContext) error {
	// DepositsController_Create: start_implement

	// Put your logic here

	// DepositsController_Create: end_implement
	return nil
}

// Show runs the show action.
func (c *DepositsController) Show(ctx *app.ShowDepositsContext) error {
	// DepositsController_Show: start_implement

	// Put your logic here

	// DepositsController_Show: end_implement
	res := &app.Deposits{}
	return ctx.OK(res)
}
