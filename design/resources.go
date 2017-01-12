package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("deposits", func() {
	Description("Cash deposit to client account at mVisa agent ")
	BasePath("/deposits")

	Action("create", func() {
		Description("creates a deposit")
		Routing(POST("/"))
		Payload(DepositsPayload)
		Response(Created)
	})

	Action("show", func() {
		Description("shows a deposit")
		Routing(GET("/:id"))
		Params(func() {
			Param("id", Integer)
		})
		Response(OK, DepositsMedia)
	})
})
