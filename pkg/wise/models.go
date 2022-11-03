package wise

import "time"

type PaymentOption struct {
	Disabled                   bool          `json:"disabled"`
	EstimatedDelivery          time.Time     `json:"estimatedDelivery"`
	FormattedEstimatedDelivery string        `json:"formattedEstimatedDelivery"`
	EstimatedDeliveryDelays    []interface{} `json:"estimatedDeliveryDelays"`
	Fee                        struct {
		Transferwise float64 `json:"transferwise"`
		PayIn        float64 `json:"payIn"`
		Discount     float64 `json:"discount"`
		Partner      float64 `json:"partner"`
		Total        float64 `json:"total"`
	} `json:"fee"`
	Price struct {
		PriceSetId int `json:"priceSetId"`
		Total      struct {
			Type  string `json:"type"`
			Label string `json:"label"`
			Value struct {
				Amount   float64 `json:"amount"`
				Currency string  `json:"currency"`
				Label    string  `json:"label:"`
			} `json:"value"`
		} `json:"total"`
		Items []struct {
			Type  string `json:"type"`
			Label string `json:"label"`
			Value struct {
				Amount   float64 `json:"amount"`
				Currency string  `json:"currency"`
				Label    string  `json:"label"`
			} `json:"value"`
			Id          int `json:"id,omitempty"`
			Explanation struct {
				PlainText string `json:"plainText"`
			} `json:"explanation,omitempty"`
		} `json:"items"`
	} `json:"price"`
	SourceAmount        float64  `json:"sourceAmount"`
	TargetAmount        float64  `json:"targetAmount"`
	SourceCurrency      string   `json:"sourceCurrency"`
	TargetCurrency      string   `json:"targetCurrency"`
	PayIn               string   `json:"payIn"`
	PayOut              string   `json:"payOut"`
	AllowedProfileTypes []string `json:"allowedProfileTypes"`
	PayInProduct        string   `json:"payInProduct"`
	FeePercentage       float64  `json:"feePercentage"`
}

type TransferFlowConfig struct {
	HighAmount struct {
		ShowFeePercentage           bool `json:"showFeePercentage"`
		TrackAsHighAmountSender     bool `json:"trackAsHighAmountSender"`
		ShowEducationStep           bool `json:"showEducationStep"`
		OfferPrefundingOption       bool `json:"offerPrefundingOption"`
		OverLimitThroughCs          bool `json:"overLimitThroughCs"`
		OverLimitThroughWiseAccount bool `json:"overLimitThroughWiseAccount"`
	} `json:"highAmount"`
}

type Notice struct {
	Text string `json:"text"`
	Link string `json:"link"`
	Type string `json:"type"`
}

type CreateQuoteResponse struct {
	SourceAmount                  float64            `json:"sourceAmount"`
	GuaranteedTargetAmountAllowed bool               `json:"guaranteedTargetAmountAllowed"`
	TargetAmountAllowed           bool               `json:"targetAmountAllowed"`
	PaymentOptions                []PaymentOption    `json:"paymentOptions"`
	Notices                       []Notice           `json:"notices"`
	TransferFlowConfig            TransferFlowConfig `json:"transferFlowConfig"`
	RateTimestamp                 time.Time          `json:"rateTimestamp"`
	ClientId                      string             `json:"clientId"`
	Id                            string             `json:"id"`
	Type                          string             `json:"type"`
	ExpirationTime                time.Time          `json:"expirationTime"`
	Status                        string             `json:"status"`
	Profile                       int                `json:"profile"`
	Rate                          float64            `json:"rate"`
	SourceCurrency                string             `json:"sourceCurrency"`
	TargetCurrency                string             `json:"targetCurrency"`
	CreatedTime                   time.Time          `json:"createdTime"`
	User                          int                `json:"user"`
	RateType                      string             `json:"rateType"`
	RateExpirationTime            time.Time          `json:"rateExpirationTime"`
	PayOut                        string             `json:"payOut"`
	GuaranteedTargetAmount        bool               `json:"guaranteedTargetAmount"`
	ProvidedAmountType            string             `json:"providedAmountType"`
	PayInCountry                  string             `json:"payInCountry"`
	Funding                       string             `json:"funding"`
}
