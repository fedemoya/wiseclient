package models

import (
	"github.com/google/uuid"
)

type TransferDetails struct {
	Reference       string `json:"reference"`
	TransferPurpose string `json:"transferPurpose"`
	SourceOfFunds   string `json:"sourceOfFunds"`
}

type CreateTransferRequestBody struct {
	TargetAccount         *int64           `json:"targetAccount"`
	QuoteUuid             *uuid.UUID       `json:"quoteUuid"`
	CustomerTransactionId *uuid.UUID       `json:"customerTransactionId"`
	Details               *TransferDetails `json:"details"`
}

type CreateTransferResponse struct {
	Id                    int64           `json:"id"`
	User                  int64           `json:"user"`
	TargetAccount         int64           `json:"targetAccount"`
	SourceAccount         int64           `json:"sourceAccount"`
	Quote                 interface{}     `json:"quote"`
	QuoteUuid             string          `json:"quoteUuid"`
	Status                string          `json:"status"`
	Reference             string          `json:"reference"`
	Rate                  float64         `json:"rate"`
	Created               string          `json:"created"`
	Business              int64           `json:"business"`
	TransferRequest       interface{}     `json:"transferRequest"`
	Details               TransferDetails `json:"details"`
	HasActiveIssues       bool            `json:"hasActiveIssues"`
	SourceCurrency        string          `json:"sourceCurrency"`
	SourceValue           float64         `json:"sourceValue"`
	TargetCurrency        string          `json:"targetCurrency"`
	TargetValue           float64         `json:"targetValue"`
	CustomerTransactionId string          `json:"customerTransactionId"`
}
