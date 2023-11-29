package requests

type GeneralBusinessPartner struct {
	FinInstCountry         string `json:"FinInstCountry"`
	FinInstCode            string `json:"FinInstCode"`
	FinInstBusinessPartner int    `json:"FinInstBusinessPartner"`
	ValidityEndDate        string `json:"ValidityEndDate"`
	ValidityStartDate      string `json:"ValidityStartDate"`
	IsMarkedForDeletion    *bool  `json:"IsMarkedForDeletion"`
}
