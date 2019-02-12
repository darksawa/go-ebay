package ebay

type ProductListingDetails struct {
	UPC      string
	EAN      string
	BrandMPN *BrandMPN
	ISBN	 string `xml:",omitempty"`
}
