package design

import (
	//. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// Output Type (MediaType). It defines types and views
// DepositMedia is the deposits resource media type.

var DepositMedia = MediaType("application/vnd.DepositMedia+json", func() {
	TypeName("deposit")
	Reference(DepositPayload)

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

// WithdrawalMedia is the withdrawal resource media type.

var WithdrawalMedia = MediaType("application/vnd.WithdrawalMedia+json", func() {
	TypeName("withdrawal")
	Reference(WithdrawalPayload)

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
