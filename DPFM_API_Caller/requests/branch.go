package requests

type Branch struct {
	FinInstCountry        string  `json:"FinInstCountry"`
	FinInstCode           string  `json:"FinInstCode"`
	FinInstBranchCode     *string `json:"FinInstBranchCode"`
	FinInstFullCode       *string `json:"FinInstFullCode"`
	FinInstBranchName     *string `json:"FinInstBranchName"`
	FinInstBranchFullName *string `json:"FinInstBranchFullName"`
	AddressID             *int    `json:"AddressID"`
	SWIFTCode             *string `json:"SWIFTCode"`
	IsMarkedForDeletion   *bool   `json:"IsMarkedForDeletion"`
}
