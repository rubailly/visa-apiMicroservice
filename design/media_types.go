package design

import (
	//. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// Output Type (MediaType). It defines types and views
// DepositsMedia is the deposits resource media type.

var DepositsMedia = MediaType("application/vnd.DepositsMedia+json", func() {
	TypeName("deposits")
	Reference(DepositsPayload)

	Attributes(func() {
		Attribute("ID")
		Attribute("acquirerCountryCode")
		Attribute("acquiringBin")
		Attribute("amount")
		Attribute("businessApplicationId")
		Attribute("localTransactionDateTime")
		Attribute("merchantCategoryCode")
		Attribute("recipientPrimaryAccountNumber")
		Attribute("retrievalReferenceNumber")
		Attribute("senderAccountNumber")
		Attribute("senderName")
		Attribute("senderReference")
		Attribute("systemsTraceAuditNumber")
		Attribute("transactionCurrencyCode")
		Attribute("transactionIdentifier")

	})

	View("default", func() {
		Attribute("ID")
		Attribute("acquirerCountryCode")
		Attribute("acquiringBin")
		Attribute("amount")
		Attribute("businessApplicationId")
		Attribute("localTransactionDateTime")
		Attribute("merchantCategoryCode")
		Attribute("recipientPrimaryAccountNumber")
		Attribute("retrievalReferenceNumber")
		Attribute("senderAccountNumber")
		Attribute("senderName")
		Attribute("senderReference")
		Attribute("systemsTraceAuditNumber")
		Attribute("transactionCurrencyCode")
		Attribute("transactionIdentifier")
	})
})
