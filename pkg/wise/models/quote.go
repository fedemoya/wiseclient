package models

const (
	QuotePayoutTypeBankTransfer string = "BANK_TRANSFER"
	QuotePayoutTypeBalance             = "BALANCE"
	QuotePayoutTypeSwift               = "SWIFT"
	QuotePayoutTypeSwiftOur            = "SWIFT_OUR"
	QuotePayoutTypeInterac             = "INTERAC"
)

type CreateQuoteRequestBody struct {
	SourceCurrency *string  `json:"sourceCurrency,omitempty"`
	TargetCurrency *string  `json:"targetCurrency,omitempty"`
	SourceAmount   *float64 `json:"sourceAmount,omitempty"`
	TargetAmount   *float64 `json:"targetAmount,omitempty"`
	PayOut         *string  `json:"payOut,omitempty"`
	PreferredPayIn *string  `json:"preferredPayIn,omitempty"`
}

type CreateQuoteRequestBodyBuilder struct {
	SourceCurrency *string
	TargetCurrency *string
	SourceAmount   *float64
	TargetAmount   *float64
	PayOut         *string
	PreferredPayIn *string
}

func NewCreateQuoteRequestBodyBuilder() *CreateQuoteRequestBodyBuilder {
	return &CreateQuoteRequestBodyBuilder{}
}

func (b *CreateQuoteRequestBodyBuilder) SetSourceCurrency(sourceCurrency string) *CreateQuoteRequestBodyBuilder {
	b.SourceCurrency = &sourceCurrency
	return b
}

func (b *CreateQuoteRequestBodyBuilder) SetTargetCurrency(targetCurrency string) *CreateQuoteRequestBodyBuilder {
	b.TargetCurrency = &targetCurrency
	return b
}

func (b *CreateQuoteRequestBodyBuilder) SetSourceAmount(sourceAmount float64) *CreateQuoteRequestBodyBuilder {
	b.SourceAmount = &sourceAmount
	return b
}

func (b *CreateQuoteRequestBodyBuilder) SetPayOut(payOut string) *CreateQuoteRequestBodyBuilder {
	b.PayOut = &payOut
	return b
}

func (b *CreateQuoteRequestBodyBuilder) SetPreferredIn(preferredPayIn string) *CreateQuoteRequestBodyBuilder {
	b.PreferredPayIn = &preferredPayIn
	return b
}

func (b *CreateQuoteRequestBodyBuilder) Build() CreateQuoteRequestBody {
	return CreateQuoteRequestBody{
		SourceCurrency: b.SourceCurrency,
		TargetCurrency: b.TargetCurrency,
		SourceAmount:   b.SourceAmount,
		TargetAmount:   b.TargetAmount,
		PayOut:         b.PayOut,
		PreferredPayIn: b.PreferredPayIn,
	}
}
