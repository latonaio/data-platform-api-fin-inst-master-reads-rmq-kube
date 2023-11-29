package requests

type General struct {
	FinInstCountry      string  `json:"FinInstCountry"`
	FinInstCode         string  `json:"FinInstCode"`
	FinInstName         *string `json:"FinInstName"`
	FinInstFullName     *string `json:"FinInstFullName"`
	AddressID           *int    `json:"AddressID"`
	SWIFTCode           *string `json:"SWIFTCode"`
	IsMarkedForDeletion *bool   `json:"IsMarkedForDeletion"`
}
