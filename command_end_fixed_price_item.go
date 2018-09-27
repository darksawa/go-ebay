package ebay

import "encoding/xml"

type EndFixedPriceItem struct {
	EndingReason string
	ItemID       string `xml:",omitempty"`
	SKU          string `xml:",omitempty"`
}

func (c EndFixedPriceItem) CallName() string {
	return "EndFixedPriceItem"
}

func (c EndFixedPriceItem) ParseResponse(r []byte) (EbayResponse, error) {
	var xmlResponse EndFixedPriceItemResponse
	err := xml.Unmarshal(r, &xmlResponse)

	return xmlResponse, err
}

func (c EndFixedPriceItem) Body() interface{} {
	type Item struct {
		EndFixedPriceItem
	}

	return Item{c}
}

type EndFixedPriceItemResponse struct {
	ebayResponse

	EndTime string
}

func (r EndFixedPriceItemResponse) ResponseErrors() ebayErrors {
	return r.ebayResponse.Errors
}
