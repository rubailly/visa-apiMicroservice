package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// Input Type.This defines the shape of the request.

// DepositPayload defines the data structure used in the create deposit request body
// It is also the base type for the deposit media type used to render deposits.

var DepositPayload = Type("DepositPayload", func() {
	Description("DepositsPayload is the type used to create deposits.")

	Attribute("acquirerCountryCode", Integer, "Country of the originator BIN.", func() {
		Example(643)
	})
	Attribute("acquiringBin", Integer, "BIN number identifies the originator of cash in transaction.", func() {
		Example(400171)
	})
	Attribute("amount", Integer, "Transaction amount in agent currency", func() {
		Example(12400)
	})
	Attribute("businessApplicationId", String, "This field is populated with business application identifier for the transaction.", func() {
		MinLength(2)
		MaxLength(2)
		Example("CI")
	})
	Attribute("localTransactionDateTime", DateTime, "The date and time the transaction takes place, ", func() {

	})
	Attribute("merchantCategoryCode", Integer, "Originators should populate 6012 for mVisa transaction.", func() {
		Example(4829)
	})

	Attribute("recipientPrimaryAccountNumber", String, "Consumer PAN.16-digit PAN as provided by the consumer to agent.", func() {
		MinLength(13)
		MaxLength(19)
		Example("4123640062698797")
	})
	Attribute("retrievalReferenceNumber", String, "Matches message to others within a given transaction", func() {
		MinLength(12)
		MaxLength(12)
		Example("430000367618")
	})

	Attribute("senderAccountNumber", String, "mVisa cash-in transactions", func() {
		Example("4541237895236")
	})
	Attribute("senderName", String, "Name of agents business name", func() {
		MaxLength(30)
		Example("Mohammed Qasim")
	})
	Attribute("senderReference", String, "A reference number unique to the agent", func() {
		MaxLength(16)
		Example("1234")
	})

	Attribute("systemsTraceAuditNumber", Integer, "Key data element", func() {
		Example(313042)
	})
	Attribute("transactionCurrencyCode", String, "Currency code", func() {
		MinLength(3)
		MaxLength(3)
		Example("USD")
	})
	Attribute("transactionIdentifier", Integer, "VisaNet reference Number for the transaction", func() {
		// transactionIdentifier is a Visa generated field that client receives in the response message.
	})

})

// WithdrawalPayload defines the data structure used in the create withdrawal request body
// It is also the base type for the withdrawal media type used to render withdrawals.

var WithdrawalPayload = Type("WithdrawalPayload", func() {
	Description("WithdrawalPayload is the type used to create a withdrawal.")

	Attribute("acquirerCountryCode", Integer, "Country of the originator BIN.", func() {
		Example(643)
	})
	Attribute("acquiringBin", Integer, "BIN number identifies the originator of cash in transaction.", func() {
		Example(400171)
	})
	Attribute("amount", Integer, "Transaction amount in agent currency", func() {
		Example(12400)
	})
	Attribute("businessApplicationId", String, "This field is populated with business application identifier for the transaction.", func() {
		MinLength(2)
		MaxLength(2)
		Example("CI")
	})
	Attribute("localTransactionDateTime", DateTime, "The date and time the transaction takes place, ", func() {

	})
	Attribute("merchantCategoryCode", Integer, "Originators should populate 6012 for mVisa transaction.", func() {
		Example(4829)
	})
	Attribute("recipientPrimaryAccountNumber", String, "Consumer PAN.16-digit PAN as provided by the consumer to agent.", func() {
		MinLength(13)
		MaxLength(19)
		Example("4123640062698797")
	})
	Attribute("retrievalReferenceNumber", String, "Matches message to others within a given transaction", func() {
		MinLength(12)
		MaxLength(12)
		Example("430000367618")
	})
	Attribute("senderAccountNumber", String, "mVisa cash-in transactions", func() {
		Example("4541237895236")
	})
	Attribute("senderName", String, "Name of agents business name", func() {
		MaxLength(30)
		Example("Mohammed Qasim")
	})

	Attribute("senderReference", String, "A reference number unique to the agent", func() {
		MaxLength(16)
		Example("1234")
	})
	Attribute("systemsTraceAuditNumber", Integer, "Key data element", func() {
		Example(313042)
	})
	Attribute("transactionCurrencyCode", String, "Currency code", func() {
		MinLength(3)
		MaxLength(3)
		Example("USD")
	})
	Attribute("transactionIdentifier", Integer, "VisaNet reference Number for the transaction", func() {
		// transactionIdentifier is a Visa generated field that client receives in the response message.
	})

})

// PaymentPayload defines the data structure used in the create payment request body
// It is also the base type for the payment media type used to render payments.

var PaymentPayload = Type("PaymentPayload", func() {
	Description("PaymentPayload is the type used to create a payment.")

	Attribute("acquirerCountryCode", Integer, "Country of the originator BIN.", func() {
		Example(643)
	})
	Attribute("acquiringBin", Integer, "BIN number identifies the originator of cash in transaction.", func() {
		Example(400171)
	})
	Attribute("amount", Integer, "Transaction amount in agent currency", func() {
		Example(12400)
	})
	Attribute("businessApplicationId", String, "This field is populated with business application identifier for the transaction.", func() {
		MinLength(2)
		MaxLength(2)
		Example("CI")
	})

	//Attribute(cardAcceptor)
	//address
	//city
	//country
	//idCode
	//name

	Attribute("feeProgramIndicator", String, "Optional field populated by recipient", func() {
		MinLength(3)
		MaxLength(3)
	})
	Attribute("localTransactionDateTime", DateTime, "The date and time the transaction takes place, ", func() {

	})

	// Attribute(purchaseIdentifier)
	// referenceNumber
	// type

	Attribute("recipientName", String, "Recipients Name", func() {
		MaxLength(30)
	})
	Attribute("recipientPrimaryAccountNumber", String, "Consumer PAN.16-digit PAN as provided by the consumer to agent.", func() {
		MinLength(13)
		MaxLength(19)
		Example("4123640062698797")
	})
	Attribute("retrievalReferenceNumber", String, "Matches message to others within a given transaction", func() {
		MinLength(12)
		MaxLength(12)
		Example("430000367618")
	})
	Attribute("secondaryId", String, "Obtains additional data along with the payment instruction", func() {
		MaxLength(28)
	})
	Attribute("senderAccountNumber", String, "mVisa cash-in transactions", func() {
		Example("4541237895236")
	})
	Attribute("senderName", String, "Name of agents business name", func() {
		MaxLength(30)
		Example("Mohammed Qasim")
	})

	Attribute("senderReference", String, "A reference number unique to the agent", func() {
		MaxLength(16)
		Example("1234")
	})
	Attribute("systemsTraceAuditNumber", Integer, "Key data element", func() {
		Example(313042)
	})
	Attribute("transactionCurrencyCode", String, "Currency code", func() {
		MinLength(3)
		MaxLength(3)
		Example("USD")
	})
	Attribute("transactionIdentifier", Integer, "VisaNet reference Number for the transaction", func() {
		// transactionIdentifier is a Visa generated field that client receives in the response message.
	})

})
