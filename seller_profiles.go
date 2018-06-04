package ebay

type SellerProfiles struct {
	SellerPaymentProfile  *SellerPaymentProfile  `xml:",omitempty"`
	SellerReturnProfile   *SellerReturnProfile   `xml:",omitempty"`
	SellerShippingProfile *SellerShippingProfile `xml:",omitempty"`
}

type SellerReturnProfile struct {
	ReturnProfileID   uint
	ReturnProfileName string `xml:",omitempty"`
}

type SellerPaymentProfile struct {
	PaymentProfileID   uint
	PaymentProfileName string `xml:",omitempty"`
}

type SellerShippingProfile struct {
	ShippingProfileID   uint
	ShippingProfileName string `xml:",omitempty"`
}