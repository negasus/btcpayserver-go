package models

type PaymentStatus string

const (
	PaymentStatusInvalid    PaymentStatus = "Invalid"
	PaymentStatusProcessing PaymentStatus = "Processing"
	PaymentStatusSettled    PaymentStatus = "Settled"
)

type Payment struct {
	ID           string        `json:"id"`
	ReceivedDate Time          `json:"receivedDate"`
	Value        Float         `json:"value"`
	Fee          Float         `json:"fee"`
	Status       PaymentStatus `json:"status"`
	Destination  string        `json:"destination"`
}

type AdditionalData struct {
	ProvidedComment          *string `json:"providedComment"`
	ConsumedLightningAddress *string `json:"consumedLightningAddress"`
	PaymentHash              *string `json:"paymentHash"` // not documented, but returned
}

type PaymentMethod struct {
	PaymentMethod     string         `json:"paymentMethod"`
	CryptoCode        string         `json:"cryptoCode"`
	Destination       string         `json:"destination"`
	PaymentLink       *string        `json:"paymentLink"`
	Rate              Float          `json:"rate"`
	PaymentMethodPaid Float          `json:"paymentMethodPaid"`
	TotalPaid         Float          `json:"totalPaid"`
	Due               Float          `json:"due"`
	Amount            Float          `json:"amount"`
	NetworkFee        Float          `json:"networkFee"`
	Payments          []*Payment     `json:"payments"`
	Activated         bool           `json:"activated"`
	AdditionalData    AdditionalData `json:"additionalData"`
}
