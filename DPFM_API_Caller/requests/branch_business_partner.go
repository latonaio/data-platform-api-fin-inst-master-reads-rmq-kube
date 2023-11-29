package requests

type BranchBusinessPartner struct {
	FinInstCountry               string `json:"FinInstCountry"`
	FinInstCode                  string `json:"FinInstCode"`
	FinInstBranchCode            string `json:"FinInstBranchCode"`
	FinInstFullCode              string `json:"FinInstFullCode"`
	FinInstBranchBusinessPartner int    `json:"FinInstBranchBusinessPartner"`
	ValidityEndDate              string `json:"ValidityEndDate"`
	ValidityStartDate            string `json:"ValidityStartDate"`
	IsMarkedForDeletion          *bool  `json:"IsMarkedForDeletion"`
}
