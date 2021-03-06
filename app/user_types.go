//************************************************************************//
// API "ChamaconektVisa": Application User Types
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

// DepositsPayload is the type used to create deposits.
type depositPayload struct {
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

// Validate validates the depositPayload type instance.
func (ut *depositPayload) Validate() (err error) {
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

// Publicize creates DepositPayload from depositPayload
func (ut *depositPayload) Publicize() *DepositPayload {
	var pub DepositPayload
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
	if ut.TransactionIdentifier != nil {
		pub.TransactionIdentifier = ut.TransactionIdentifier
	}
	return &pub
}

// DepositsPayload is the type used to create deposits.
type DepositPayload struct {
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

// Validate validates the DepositPayload type instance.
func (ut *DepositPayload) Validate() (err error) {
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

// PaymentPayload is the type used to create a payment.
type paymentPayload struct {
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

// Validate validates the paymentPayload type instance.
func (ut *paymentPayload) Validate() (err error) {
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
	if ut.FeeProgramIndicator != nil {
		if utf8.RuneCountInString(*ut.FeeProgramIndicator) < 3 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.feeProgramIndicator`, *ut.FeeProgramIndicator, utf8.RuneCountInString(*ut.FeeProgramIndicator), 3, true))
		}
	}
	if ut.FeeProgramIndicator != nil {
		if utf8.RuneCountInString(*ut.FeeProgramIndicator) > 3 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.feeProgramIndicator`, *ut.FeeProgramIndicator, utf8.RuneCountInString(*ut.FeeProgramIndicator), 3, false))
		}
	}
	if ut.RecipientName != nil {
		if utf8.RuneCountInString(*ut.RecipientName) > 30 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.recipientName`, *ut.RecipientName, utf8.RuneCountInString(*ut.RecipientName), 30, false))
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
	if ut.SecondaryID != nil {
		if utf8.RuneCountInString(*ut.SecondaryID) > 28 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.secondaryId`, *ut.SecondaryID, utf8.RuneCountInString(*ut.SecondaryID), 28, false))
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

// Publicize creates PaymentPayload from paymentPayload
func (ut *paymentPayload) Publicize() *PaymentPayload {
	var pub PaymentPayload
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
	if ut.FeeProgramIndicator != nil {
		pub.FeeProgramIndicator = ut.FeeProgramIndicator
	}
	if ut.LocalTransactionDateTime != nil {
		pub.LocalTransactionDateTime = ut.LocalTransactionDateTime
	}
	if ut.RecipientName != nil {
		pub.RecipientName = ut.RecipientName
	}
	if ut.RecipientPrimaryAccountNumber != nil {
		pub.RecipientPrimaryAccountNumber = ut.RecipientPrimaryAccountNumber
	}
	if ut.RetrievalReferenceNumber != nil {
		pub.RetrievalReferenceNumber = ut.RetrievalReferenceNumber
	}
	if ut.SecondaryID != nil {
		pub.SecondaryID = ut.SecondaryID
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
	if ut.TransactionIdentifier != nil {
		pub.TransactionIdentifier = ut.TransactionIdentifier
	}
	return &pub
}

// PaymentPayload is the type used to create a payment.
type PaymentPayload struct {
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

// Validate validates the PaymentPayload type instance.
func (ut *PaymentPayload) Validate() (err error) {
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
	if ut.FeeProgramIndicator != nil {
		if utf8.RuneCountInString(*ut.FeeProgramIndicator) < 3 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.feeProgramIndicator`, *ut.FeeProgramIndicator, utf8.RuneCountInString(*ut.FeeProgramIndicator), 3, true))
		}
	}
	if ut.FeeProgramIndicator != nil {
		if utf8.RuneCountInString(*ut.FeeProgramIndicator) > 3 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.feeProgramIndicator`, *ut.FeeProgramIndicator, utf8.RuneCountInString(*ut.FeeProgramIndicator), 3, false))
		}
	}
	if ut.RecipientName != nil {
		if utf8.RuneCountInString(*ut.RecipientName) > 30 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.recipientName`, *ut.RecipientName, utf8.RuneCountInString(*ut.RecipientName), 30, false))
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
	if ut.SecondaryID != nil {
		if utf8.RuneCountInString(*ut.SecondaryID) > 28 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.secondaryId`, *ut.SecondaryID, utf8.RuneCountInString(*ut.SecondaryID), 28, false))
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

// WithdrawalPayload is the type used to create a withdrawal.
type withdrawalPayload struct {
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

// Validate validates the withdrawalPayload type instance.
func (ut *withdrawalPayload) Validate() (err error) {
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

// Publicize creates WithdrawalPayload from withdrawalPayload
func (ut *withdrawalPayload) Publicize() *WithdrawalPayload {
	var pub WithdrawalPayload
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
	if ut.TransactionIdentifier != nil {
		pub.TransactionIdentifier = ut.TransactionIdentifier
	}
	return &pub
}

// WithdrawalPayload is the type used to create a withdrawal.
type WithdrawalPayload struct {
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

// Validate validates the WithdrawalPayload type instance.
func (ut *WithdrawalPayload) Validate() (err error) {
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
