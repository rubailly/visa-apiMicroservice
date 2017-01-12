//************************************************************************//
// API "mVisa": Application User Types
//
// Generated with goagen v1.0.0, command line:
// $ goagen
// --design=ChamaconektVisa/design
// --out=$(GOPATH)/src/ChamaconektVisa
// --version=v1.0.0
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package client

import (
	"github.com/goadesign/goa"
	"time"
	"unicode/utf8"
)

// DepositsPayload is the type used to create deposits.
type depositsPayload struct {
	// Country of the originator BIN.
	AcquirerCountryCode *int `form:"acquirerCountryCode,omitempty" json:"acquirerCountryCode,omitempty" xml:"acquirerCountryCode,omitempty"`
	// BIN number identifies the originator of cash in transaction.
	AcquiringBin *int `form:"acquiringBin,omitempty" json:"acquiringBin,omitempty" xml:"acquiringBin,omitempty"`
	// Transaction amount in agent currency
	Amount *int `form:"amount,omitempty" json:"amount,omitempty" xml:"amount,omitempty"`
	// This field is populated with business application identifier for the transaction.
	BusinessApplicationID *string `form:"businessApplicationId,omitempty" json:"businessApplicationId,omitempty" xml:"businessApplicationId,omitempty"`
	// The date and time the transaction takes place,
	LocalTransactionDateTime *time.Time `form:"localTransactionDateTime,omitempty" json:"localTransactionDateTime,omitempty" xml:"localTransactionDateTime,omitempty"`
	// Originators should populate 6012 for mVisa transaction.
	MerchantCategoryCode *int `form:"merchantCategoryCode,omitempty" json:"merchantCategoryCode,omitempty" xml:"merchantCategoryCode,omitempty"`
	// Consumer PAN.16-digit PAN as provided by the consumer to agent.
	RecipientPrimaryAccountNumber *string `form:"recipientPrimaryAccountNumber,omitempty" json:"recipientPrimaryAccountNumber,omitempty" xml:"recipientPrimaryAccountNumber,omitempty"`
	// Matches message to others within a given transaction
	RetrievalReferenceNumber *string `form:"retrievalReferenceNumber,omitempty" json:"retrievalReferenceNumber,omitempty" xml:"retrievalReferenceNumber,omitempty"`
	// mVisa cash-in transactions
	SenderAccountNumber *string `form:"senderAccountNumber,omitempty" json:"senderAccountNumber,omitempty" xml:"senderAccountNumber,omitempty"`
	// Name of agents business name
	SenderName *string `form:"senderName,omitempty" json:"senderName,omitempty" xml:"senderName,omitempty"`
	// A reference number unique to the agent
	SenderReference *string `form:"senderReference,omitempty" json:"senderReference,omitempty" xml:"senderReference,omitempty"`
	// Key data element
	SystemsTraceAuditNumber *int `form:"systemsTraceAuditNumber,omitempty" json:"systemsTraceAuditNumber,omitempty" xml:"systemsTraceAuditNumber,omitempty"`
	// Currency code
	TransactionCurrencyCode *string `form:"transactionCurrencyCode,omitempty" json:"transactionCurrencyCode,omitempty" xml:"transactionCurrencyCode,omitempty"`
}

// Validate validates the depositsPayload type instance.
func (ut *depositsPayload) Validate() (err error) {
	if ut.BusinessApplicationID != nil {
		if utf8.RuneCountInString(*ut.BusinessApplicationID) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.businessApplicationId`, *ut.BusinessApplicationID, utf8.RuneCountInString(*ut.BusinessApplicationID), 2, true))
		}
	}
	if ut.BusinessApplicationID != nil {
		if utf8.RuneCountInString(*ut.BusinessApplicationID) > 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.businessApplicationId`, *ut.BusinessApplicationID, utf8.RuneCountInString(*ut.BusinessApplicationID), 2, false))
		}
	}
	if ut.RecipientPrimaryAccountNumber != nil {
		if utf8.RuneCountInString(*ut.RecipientPrimaryAccountNumber) < 13 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.recipientPrimaryAccountNumber`, *ut.RecipientPrimaryAccountNumber, utf8.RuneCountInString(*ut.RecipientPrimaryAccountNumber), 13, true))
		}
	}
	if ut.RecipientPrimaryAccountNumber != nil {
		if utf8.RuneCountInString(*ut.RecipientPrimaryAccountNumber) > 19 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.recipientPrimaryAccountNumber`, *ut.RecipientPrimaryAccountNumber, utf8.RuneCountInString(*ut.RecipientPrimaryAccountNumber), 19, false))
		}
	}
	if ut.RetrievalReferenceNumber != nil {
		if utf8.RuneCountInString(*ut.RetrievalReferenceNumber) < 12 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.retrievalReferenceNumber`, *ut.RetrievalReferenceNumber, utf8.RuneCountInString(*ut.RetrievalReferenceNumber), 12, true))
		}
	}
	if ut.RetrievalReferenceNumber != nil {
		if utf8.RuneCountInString(*ut.RetrievalReferenceNumber) > 12 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.retrievalReferenceNumber`, *ut.RetrievalReferenceNumber, utf8.RuneCountInString(*ut.RetrievalReferenceNumber), 12, false))
		}
	}
	if ut.SenderName != nil {
		if utf8.RuneCountInString(*ut.SenderName) > 30 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.senderName`, *ut.SenderName, utf8.RuneCountInString(*ut.SenderName), 30, false))
		}
	}
	if ut.SenderReference != nil {
		if utf8.RuneCountInString(*ut.SenderReference) > 16 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.senderReference`, *ut.SenderReference, utf8.RuneCountInString(*ut.SenderReference), 16, false))
		}
	}
	if ut.TransactionCurrencyCode != nil {
		if utf8.RuneCountInString(*ut.TransactionCurrencyCode) < 3 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.transactionCurrencyCode`, *ut.TransactionCurrencyCode, utf8.RuneCountInString(*ut.TransactionCurrencyCode), 3, true))
		}
	}
	if ut.TransactionCurrencyCode != nil {
		if utf8.RuneCountInString(*ut.TransactionCurrencyCode) > 3 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.transactionCurrencyCode`, *ut.TransactionCurrencyCode, utf8.RuneCountInString(*ut.TransactionCurrencyCode), 3, false))
		}
	}
	return
}

// Publicize creates DepositsPayload from depositsPayload
func (ut *depositsPayload) Publicize() *DepositsPayload {
	var pub DepositsPayload
	if ut.AcquirerCountryCode != nil {
		pub.AcquirerCountryCode = ut.AcquirerCountryCode
	}
	if ut.AcquiringBin != nil {
		pub.AcquiringBin = ut.AcquiringBin
	}
	if ut.Amount != nil {
		pub.Amount = ut.Amount
	}
	if ut.BusinessApplicationID != nil {
		pub.BusinessApplicationID = ut.BusinessApplicationID
	}
	if ut.LocalTransactionDateTime != nil {
		pub.LocalTransactionDateTime = ut.LocalTransactionDateTime
	}
	if ut.MerchantCategoryCode != nil {
		pub.MerchantCategoryCode = ut.MerchantCategoryCode
	}
	if ut.RecipientPrimaryAccountNumber != nil {
		pub.RecipientPrimaryAccountNumber = ut.RecipientPrimaryAccountNumber
	}
	if ut.RetrievalReferenceNumber != nil {
		pub.RetrievalReferenceNumber = ut.RetrievalReferenceNumber
	}
	if ut.SenderAccountNumber != nil {
		pub.SenderAccountNumber = ut.SenderAccountNumber
	}
	if ut.SenderName != nil {
		pub.SenderName = ut.SenderName
	}
	if ut.SenderReference != nil {
		pub.SenderReference = ut.SenderReference
	}
	if ut.SystemsTraceAuditNumber != nil {
		pub.SystemsTraceAuditNumber = ut.SystemsTraceAuditNumber
	}
	if ut.TransactionCurrencyCode != nil {
		pub.TransactionCurrencyCode = ut.TransactionCurrencyCode
	}
	return &pub
}

// DepositsPayload is the type used to create deposits.
type DepositsPayload struct {
	// Country of the originator BIN.
	AcquirerCountryCode *int `form:"acquirerCountryCode,omitempty" json:"acquirerCountryCode,omitempty" xml:"acquirerCountryCode,omitempty"`
	// BIN number identifies the originator of cash in transaction.
	AcquiringBin *int `form:"acquiringBin,omitempty" json:"acquiringBin,omitempty" xml:"acquiringBin,omitempty"`
	// Transaction amount in agent currency
	Amount *int `form:"amount,omitempty" json:"amount,omitempty" xml:"amount,omitempty"`
	// This field is populated with business application identifier for the transaction.
	BusinessApplicationID *string `form:"businessApplicationId,omitempty" json:"businessApplicationId,omitempty" xml:"businessApplicationId,omitempty"`
	// The date and time the transaction takes place,
	LocalTransactionDateTime *time.Time `form:"localTransactionDateTime,omitempty" json:"localTransactionDateTime,omitempty" xml:"localTransactionDateTime,omitempty"`
	// Originators should populate 6012 for mVisa transaction.
	MerchantCategoryCode *int `form:"merchantCategoryCode,omitempty" json:"merchantCategoryCode,omitempty" xml:"merchantCategoryCode,omitempty"`
	// Consumer PAN.16-digit PAN as provided by the consumer to agent.
	RecipientPrimaryAccountNumber *string `form:"recipientPrimaryAccountNumber,omitempty" json:"recipientPrimaryAccountNumber,omitempty" xml:"recipientPrimaryAccountNumber,omitempty"`
	// Matches message to others within a given transaction
	RetrievalReferenceNumber *string `form:"retrievalReferenceNumber,omitempty" json:"retrievalReferenceNumber,omitempty" xml:"retrievalReferenceNumber,omitempty"`
	// mVisa cash-in transactions
	SenderAccountNumber *string `form:"senderAccountNumber,omitempty" json:"senderAccountNumber,omitempty" xml:"senderAccountNumber,omitempty"`
	// Name of agents business name
	SenderName *string `form:"senderName,omitempty" json:"senderName,omitempty" xml:"senderName,omitempty"`
	// A reference number unique to the agent
	SenderReference *string `form:"senderReference,omitempty" json:"senderReference,omitempty" xml:"senderReference,omitempty"`
	// Key data element
	SystemsTraceAuditNumber *int `form:"systemsTraceAuditNumber,omitempty" json:"systemsTraceAuditNumber,omitempty" xml:"systemsTraceAuditNumber,omitempty"`
	// Currency code
	TransactionCurrencyCode *string `form:"transactionCurrencyCode,omitempty" json:"transactionCurrencyCode,omitempty" xml:"transactionCurrencyCode,omitempty"`
}

// Validate validates the DepositsPayload type instance.
func (ut *DepositsPayload) Validate() (err error) {
	if ut.BusinessApplicationID != nil {
		if utf8.RuneCountInString(*ut.BusinessApplicationID) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.businessApplicationId`, *ut.BusinessApplicationID, utf8.RuneCountInString(*ut.BusinessApplicationID), 2, true))
		}
	}
	if ut.BusinessApplicationID != nil {
		if utf8.RuneCountInString(*ut.BusinessApplicationID) > 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.businessApplicationId`, *ut.BusinessApplicationID, utf8.RuneCountInString(*ut.BusinessApplicationID), 2, false))
		}
	}
	if ut.RecipientPrimaryAccountNumber != nil {
		if utf8.RuneCountInString(*ut.RecipientPrimaryAccountNumber) < 13 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.recipientPrimaryAccountNumber`, *ut.RecipientPrimaryAccountNumber, utf8.RuneCountInString(*ut.RecipientPrimaryAccountNumber), 13, true))
		}
	}
	if ut.RecipientPrimaryAccountNumber != nil {
		if utf8.RuneCountInString(*ut.RecipientPrimaryAccountNumber) > 19 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.recipientPrimaryAccountNumber`, *ut.RecipientPrimaryAccountNumber, utf8.RuneCountInString(*ut.RecipientPrimaryAccountNumber), 19, false))
		}
	}
	if ut.RetrievalReferenceNumber != nil {
		if utf8.RuneCountInString(*ut.RetrievalReferenceNumber) < 12 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.retrievalReferenceNumber`, *ut.RetrievalReferenceNumber, utf8.RuneCountInString(*ut.RetrievalReferenceNumber), 12, true))
		}
	}
	if ut.RetrievalReferenceNumber != nil {
		if utf8.RuneCountInString(*ut.RetrievalReferenceNumber) > 12 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.retrievalReferenceNumber`, *ut.RetrievalReferenceNumber, utf8.RuneCountInString(*ut.RetrievalReferenceNumber), 12, false))
		}
	}
	if ut.SenderName != nil {
		if utf8.RuneCountInString(*ut.SenderName) > 30 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.senderName`, *ut.SenderName, utf8.RuneCountInString(*ut.SenderName), 30, false))
		}
	}
	if ut.SenderReference != nil {
		if utf8.RuneCountInString(*ut.SenderReference) > 16 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.senderReference`, *ut.SenderReference, utf8.RuneCountInString(*ut.SenderReference), 16, false))
		}
	}
	if ut.TransactionCurrencyCode != nil {
		if utf8.RuneCountInString(*ut.TransactionCurrencyCode) < 3 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.transactionCurrencyCode`, *ut.TransactionCurrencyCode, utf8.RuneCountInString(*ut.TransactionCurrencyCode), 3, true))
		}
	}
	if ut.TransactionCurrencyCode != nil {
		if utf8.RuneCountInString(*ut.TransactionCurrencyCode) > 3 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.transactionCurrencyCode`, *ut.TransactionCurrencyCode, utf8.RuneCountInString(*ut.TransactionCurrencyCode), 3, false))
		}
	}
	return
}
