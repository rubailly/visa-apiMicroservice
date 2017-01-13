//************************************************************************//
// API "ChamaconektVisa": Application Media Types
//
// Generated with goagen v1.0.0, command line:
// $ goagen
// --design=ChamaconektVisa/design
// --out=$(GOPATH)/src/ChamaconektVisa
// --version=v1.0.0
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package app

import (
	"github.com/goadesign/goa"
	"time"
	"unicode/utf8"
)

// deposit media type (default view)
//
// Identifier: application/vnd.depositmedia+json; view=default
type Deposit struct {
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
	// VisaNet reference Number for the transaction
	TransactionIdentifier *int `form:"transactionIdentifier,omitempty" json:"transactionIdentifier,omitempty" xml:"transactionIdentifier,omitempty"`
}

// Validate validates the Deposit media type instance.
func (mt *Deposit) Validate() (err error) {
	if mt.BusinessApplicationID != nil {
		if utf8.RuneCountInString(*mt.BusinessApplicationID) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.businessApplicationId`, *mt.BusinessApplicationID, utf8.RuneCountInString(*mt.BusinessApplicationID), 2, true))
		}
	}
	if mt.BusinessApplicationID != nil {
		if utf8.RuneCountInString(*mt.BusinessApplicationID) > 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.businessApplicationId`, *mt.BusinessApplicationID, utf8.RuneCountInString(*mt.BusinessApplicationID), 2, false))
		}
	}
	if mt.RecipientPrimaryAccountNumber != nil {
		if utf8.RuneCountInString(*mt.RecipientPrimaryAccountNumber) < 13 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.recipientPrimaryAccountNumber`, *mt.RecipientPrimaryAccountNumber, utf8.RuneCountInString(*mt.RecipientPrimaryAccountNumber), 13, true))
		}
	}
	if mt.RecipientPrimaryAccountNumber != nil {
		if utf8.RuneCountInString(*mt.RecipientPrimaryAccountNumber) > 19 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.recipientPrimaryAccountNumber`, *mt.RecipientPrimaryAccountNumber, utf8.RuneCountInString(*mt.RecipientPrimaryAccountNumber), 19, false))
		}
	}
	if mt.RetrievalReferenceNumber != nil {
		if utf8.RuneCountInString(*mt.RetrievalReferenceNumber) < 12 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.retrievalReferenceNumber`, *mt.RetrievalReferenceNumber, utf8.RuneCountInString(*mt.RetrievalReferenceNumber), 12, true))
		}
	}
	if mt.RetrievalReferenceNumber != nil {
		if utf8.RuneCountInString(*mt.RetrievalReferenceNumber) > 12 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.retrievalReferenceNumber`, *mt.RetrievalReferenceNumber, utf8.RuneCountInString(*mt.RetrievalReferenceNumber), 12, false))
		}
	}
	if mt.SenderName != nil {
		if utf8.RuneCountInString(*mt.SenderName) > 30 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.senderName`, *mt.SenderName, utf8.RuneCountInString(*mt.SenderName), 30, false))
		}
	}
	if mt.SenderReference != nil {
		if utf8.RuneCountInString(*mt.SenderReference) > 16 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.senderReference`, *mt.SenderReference, utf8.RuneCountInString(*mt.SenderReference), 16, false))
		}
	}
	if mt.TransactionCurrencyCode != nil {
		if utf8.RuneCountInString(*mt.TransactionCurrencyCode) < 3 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.transactionCurrencyCode`, *mt.TransactionCurrencyCode, utf8.RuneCountInString(*mt.TransactionCurrencyCode), 3, true))
		}
	}
	if mt.TransactionCurrencyCode != nil {
		if utf8.RuneCountInString(*mt.TransactionCurrencyCode) > 3 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.transactionCurrencyCode`, *mt.TransactionCurrencyCode, utf8.RuneCountInString(*mt.TransactionCurrencyCode), 3, false))
		}
	}
	return
}

// payment media type (default view)
//
// Identifier: application/vnd.paymentmedia+json; view=default
type Payment struct {
	// Country of the originator BIN.
	AcquirerCountryCode *int `form:"acquirerCountryCode,omitempty" json:"acquirerCountryCode,omitempty" xml:"acquirerCountryCode,omitempty"`
	// BIN number identifies the originator of cash in transaction.
	AcquiringBin *int `form:"acquiringBin,omitempty" json:"acquiringBin,omitempty" xml:"acquiringBin,omitempty"`
	// Transaction amount in agent currency
	Amount *int `form:"amount,omitempty" json:"amount,omitempty" xml:"amount,omitempty"`
	// This field is populated with business application identifier for the transaction.
	BusinessApplicationID *string `form:"businessApplicationId,omitempty" json:"businessApplicationId,omitempty" xml:"businessApplicationId,omitempty"`
	// Optional field populated by recipient
	FeeProgramIndicator *string `form:"feeProgramIndicator,omitempty" json:"feeProgramIndicator,omitempty" xml:"feeProgramIndicator,omitempty"`
	// The date and time the transaction takes place,
	LocalTransactionDateTime *time.Time `form:"localTransactionDateTime,omitempty" json:"localTransactionDateTime,omitempty" xml:"localTransactionDateTime,omitempty"`
	MerchantCategoryCode     *string    `form:"merchantCategoryCode,omitempty" json:"merchantCategoryCode,omitempty" xml:"merchantCategoryCode,omitempty"`
	// Recipients Name
	RecipientName *string `form:"recipientName,omitempty" json:"recipientName,omitempty" xml:"recipientName,omitempty"`
	// Consumer PAN.16-digit PAN as provided by the consumer to agent.
	RecipientPrimaryAccountNumber *string `form:"recipientPrimaryAccountNumber,omitempty" json:"recipientPrimaryAccountNumber,omitempty" xml:"recipientPrimaryAccountNumber,omitempty"`
	// Matches message to others within a given transaction
	RetrievalReferenceNumber *string `form:"retrievalReferenceNumber,omitempty" json:"retrievalReferenceNumber,omitempty" xml:"retrievalReferenceNumber,omitempty"`
	// Obtains additional data along with the payment instruction
	SecondaryID *string `form:"secondaryId,omitempty" json:"secondaryId,omitempty" xml:"secondaryId,omitempty"`
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
	// VisaNet reference Number for the transaction
	TransactionIdentifier *int `form:"transactionIdentifier,omitempty" json:"transactionIdentifier,omitempty" xml:"transactionIdentifier,omitempty"`
}

// Validate validates the Payment media type instance.
func (mt *Payment) Validate() (err error) {
	if mt.BusinessApplicationID != nil {
		if utf8.RuneCountInString(*mt.BusinessApplicationID) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.businessApplicationId`, *mt.BusinessApplicationID, utf8.RuneCountInString(*mt.BusinessApplicationID), 2, true))
		}
	}
	if mt.BusinessApplicationID != nil {
		if utf8.RuneCountInString(*mt.BusinessApplicationID) > 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.businessApplicationId`, *mt.BusinessApplicationID, utf8.RuneCountInString(*mt.BusinessApplicationID), 2, false))
		}
	}
	if mt.FeeProgramIndicator != nil {
		if utf8.RuneCountInString(*mt.FeeProgramIndicator) < 3 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.feeProgramIndicator`, *mt.FeeProgramIndicator, utf8.RuneCountInString(*mt.FeeProgramIndicator), 3, true))
		}
	}
	if mt.FeeProgramIndicator != nil {
		if utf8.RuneCountInString(*mt.FeeProgramIndicator) > 3 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.feeProgramIndicator`, *mt.FeeProgramIndicator, utf8.RuneCountInString(*mt.FeeProgramIndicator), 3, false))
		}
	}
	if mt.RecipientName != nil {
		if utf8.RuneCountInString(*mt.RecipientName) > 30 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.recipientName`, *mt.RecipientName, utf8.RuneCountInString(*mt.RecipientName), 30, false))
		}
	}
	if mt.RecipientPrimaryAccountNumber != nil {
		if utf8.RuneCountInString(*mt.RecipientPrimaryAccountNumber) < 13 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.recipientPrimaryAccountNumber`, *mt.RecipientPrimaryAccountNumber, utf8.RuneCountInString(*mt.RecipientPrimaryAccountNumber), 13, true))
		}
	}
	if mt.RecipientPrimaryAccountNumber != nil {
		if utf8.RuneCountInString(*mt.RecipientPrimaryAccountNumber) > 19 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.recipientPrimaryAccountNumber`, *mt.RecipientPrimaryAccountNumber, utf8.RuneCountInString(*mt.RecipientPrimaryAccountNumber), 19, false))
		}
	}
	if mt.RetrievalReferenceNumber != nil {
		if utf8.RuneCountInString(*mt.RetrievalReferenceNumber) < 12 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.retrievalReferenceNumber`, *mt.RetrievalReferenceNumber, utf8.RuneCountInString(*mt.RetrievalReferenceNumber), 12, true))
		}
	}
	if mt.RetrievalReferenceNumber != nil {
		if utf8.RuneCountInString(*mt.RetrievalReferenceNumber) > 12 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.retrievalReferenceNumber`, *mt.RetrievalReferenceNumber, utf8.RuneCountInString(*mt.RetrievalReferenceNumber), 12, false))
		}
	}
	if mt.SecondaryID != nil {
		if utf8.RuneCountInString(*mt.SecondaryID) > 28 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.secondaryId`, *mt.SecondaryID, utf8.RuneCountInString(*mt.SecondaryID), 28, false))
		}
	}
	if mt.SenderName != nil {
		if utf8.RuneCountInString(*mt.SenderName) > 30 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.senderName`, *mt.SenderName, utf8.RuneCountInString(*mt.SenderName), 30, false))
		}
	}
	if mt.SenderReference != nil {
		if utf8.RuneCountInString(*mt.SenderReference) > 16 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.senderReference`, *mt.SenderReference, utf8.RuneCountInString(*mt.SenderReference), 16, false))
		}
	}
	if mt.TransactionCurrencyCode != nil {
		if utf8.RuneCountInString(*mt.TransactionCurrencyCode) < 3 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.transactionCurrencyCode`, *mt.TransactionCurrencyCode, utf8.RuneCountInString(*mt.TransactionCurrencyCode), 3, true))
		}
	}
	if mt.TransactionCurrencyCode != nil {
		if utf8.RuneCountInString(*mt.TransactionCurrencyCode) > 3 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.transactionCurrencyCode`, *mt.TransactionCurrencyCode, utf8.RuneCountInString(*mt.TransactionCurrencyCode), 3, false))
		}
	}
	return
}

// withdrawal media type (default view)
//
// Identifier: application/vnd.withdrawalmedia+json; view=default
type Withdrawal struct {
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
	// VisaNet reference Number for the transaction
	TransactionIdentifier *int `form:"transactionIdentifier,omitempty" json:"transactionIdentifier,omitempty" xml:"transactionIdentifier,omitempty"`
}

// Validate validates the Withdrawal media type instance.
func (mt *Withdrawal) Validate() (err error) {
	if mt.BusinessApplicationID != nil {
		if utf8.RuneCountInString(*mt.BusinessApplicationID) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.businessApplicationId`, *mt.BusinessApplicationID, utf8.RuneCountInString(*mt.BusinessApplicationID), 2, true))
		}
	}
	if mt.BusinessApplicationID != nil {
		if utf8.RuneCountInString(*mt.BusinessApplicationID) > 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.businessApplicationId`, *mt.BusinessApplicationID, utf8.RuneCountInString(*mt.BusinessApplicationID), 2, false))
		}
	}
	if mt.RecipientPrimaryAccountNumber != nil {
		if utf8.RuneCountInString(*mt.RecipientPrimaryAccountNumber) < 13 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.recipientPrimaryAccountNumber`, *mt.RecipientPrimaryAccountNumber, utf8.RuneCountInString(*mt.RecipientPrimaryAccountNumber), 13, true))
		}
	}
	if mt.RecipientPrimaryAccountNumber != nil {
		if utf8.RuneCountInString(*mt.RecipientPrimaryAccountNumber) > 19 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.recipientPrimaryAccountNumber`, *mt.RecipientPrimaryAccountNumber, utf8.RuneCountInString(*mt.RecipientPrimaryAccountNumber), 19, false))
		}
	}
	if mt.RetrievalReferenceNumber != nil {
		if utf8.RuneCountInString(*mt.RetrievalReferenceNumber) < 12 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.retrievalReferenceNumber`, *mt.RetrievalReferenceNumber, utf8.RuneCountInString(*mt.RetrievalReferenceNumber), 12, true))
		}
	}
	if mt.RetrievalReferenceNumber != nil {
		if utf8.RuneCountInString(*mt.RetrievalReferenceNumber) > 12 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.retrievalReferenceNumber`, *mt.RetrievalReferenceNumber, utf8.RuneCountInString(*mt.RetrievalReferenceNumber), 12, false))
		}
	}
	if mt.SenderName != nil {
		if utf8.RuneCountInString(*mt.SenderName) > 30 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.senderName`, *mt.SenderName, utf8.RuneCountInString(*mt.SenderName), 30, false))
		}
	}
	if mt.SenderReference != nil {
		if utf8.RuneCountInString(*mt.SenderReference) > 16 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.senderReference`, *mt.SenderReference, utf8.RuneCountInString(*mt.SenderReference), 16, false))
		}
	}
	if mt.TransactionCurrencyCode != nil {
		if utf8.RuneCountInString(*mt.TransactionCurrencyCode) < 3 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.transactionCurrencyCode`, *mt.TransactionCurrencyCode, utf8.RuneCountInString(*mt.TransactionCurrencyCode), 3, true))
		}
	}
	if mt.TransactionCurrencyCode != nil {
		if utf8.RuneCountInString(*mt.TransactionCurrencyCode) > 3 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.transactionCurrencyCode`, *mt.TransactionCurrencyCode, utf8.RuneCountInString(*mt.TransactionCurrencyCode), 3, false))
		}
	}
	return
}
