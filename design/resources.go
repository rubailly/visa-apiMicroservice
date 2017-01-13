package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("deposit", func() {
	Description("Cash deposit to client account at mVisa agent ")
	BasePath("/deposit")

	Action("create", func() {
		Description("creates a deposit")
		Routing(POST("/"))
		Payload(DepositPayload)
		Response(Created)
	})

	Action("show", func() {
		Description("shows a deposit")
		Routing(GET("/:id"))
		Params(func() {
			Param("id", Integer)
		})
		Response(OK, DepositMedia)
	})
})

var _ = Resource("withdrawal", func() {
	Description("Cash withdrawal by client from account at mVisa agent ")
	BasePath("/withdrawal")

	Action("create", func() {
		Description("creates a withdrawal")
		Routing(POST("/"))
		Payload(WithdrawalPayload)
		Response(Created)
	})

	Action("show", func() {
		Description("shows a withdrawal")
		Routing(GET("/:id"))
		Params(func() {
			Param("id", Integer)
		})
		Response(OK, WithdrawalMedia)
	})
})

var _ = Resource("payment", func() {
	Description("Payment to a merchant for goods or services purchased")
	BasePath("/payment")

	Action("create", func() {
		Description("creates a payment")
		Routing(POST("/"))
		Payload(PaymentPayload)
		Response(Created)
	})

	Action("show", func() {
		Description("shows a payment")
		Routing(GET("/:id"))
		Params(func() {
			Param("id", Integer)
		})
		Response(OK, PaymentMedia)
	})
})

var _ = Resource("swagger", func() {
	Description("The API Swagger specification")

	Files("/swagger.json", "swagger/swagger.json")
	Files("/swagger-ui/*filepath", "swagger-ui/")
})
