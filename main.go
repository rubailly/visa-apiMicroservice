//go:generate goagen bootstrap -d ChamaconektVisa/design

package main

import (
	"ChamaconektVisa/app"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
)

func main() {
	// Create service
	service := goa.New("mVisa")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "deposit" controller
	c := NewDepositController(service)
	app.MountDepositController(service, c)
	// Mount "withdrawal" controller
	c2 := NewWithdrawalController(service)
	app.MountWithdrawalController(service, c2)

	// Start service
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}
}
