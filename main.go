//go:generate goagen bootstrap -d ChamaconektVisa/design

package main

import (
	"ChamaconektVisa/app"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
)

func main() {
	// Create service
	service := goa.New("ChamaconektVisa")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "deposit" controller
	c := NewDepositController(service)
	app.MountDepositController(service, c)
	// Mount "payment" controller
	c2 := NewPaymentController(service)
	app.MountPaymentController(service, c2)
	// Mount "withdrawal" controller
	c3 := NewWithdrawalController(service)
	app.MountWithdrawalController(service, c3)

	// Start service
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}
}
