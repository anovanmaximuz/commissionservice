package representation

type CommissionRepresentation struct {
	MerchantId          string    `json:"merchantId"`
	Sku                 string    `json:"sku"`
	BrandId             string    `json:"brandId"`
	Rate                float64   `json:"rate"`
}

