package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("ChamaconektVisa", func() { // Global variable is used.We dont care about the return value hence the _ .
	Title("The Visa Payment Service")
	Description("A Chamaconekt API service that interacts with the Visa API")
	Contact(func() {
		Name("William Ondenge")
		Email("ondengew@gmail.com")
	})
	License(func() {
		Name("Apache 2.0")
		URL("")
	})
	Scheme("http")
	Host("localhost:8080")
})
