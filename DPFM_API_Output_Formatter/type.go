package dpfm_api_output_formatter

type SDC struct {
	ConnectionKey       string      `json:"connection_key"`
	Result              bool        `json:"result"`
	RedisKey            string      `json:"redis_key"`
	Filepath            string      `json:"filepath"`
	APIStatusCode       int         `json:"api_status_code"`
	RuntimeSessionID    string      `json:"runtime_session_id"`
	BusinessPartnerID   *int        `json:"business_partner"`
	ServiceLabel        string      `json:"service_label"`
	APIType             string      `json:"api_type"`
	Message             interface{} `json:"message"`
	APISchema           string      `json:"api_schema"`
	Accepter            []string    `json:"accepter"`
	Deleted             bool        `json:"deleted"`
	SQLUpdateResult     *bool       `json:"sql_update_result"`
	SQLUpdateError      string      `json:"sql_update_error"`
	SubfuncResult       *bool       `json:"subfunc_result"`
	SubfuncError        string      `json:"subfunc_error"`
	ExconfResult        *bool       `json:"exconf_result"`
	ExconfError         string      `json:"exconf_error"`
	APIProcessingResult *bool       `json:"api_processing_result"`
	APIProcessingError  string      `json:"api_processing_error"`
}

type Message struct {
	General                *[]General                `json:"General"`
	GeneralBusinessPartner *[]GeneralBusinessPartner `json:"GeneralBusinessPartner"`
	Branch                 *[]Branch                 `json:"Branch"`
	BranchBusinessPartner  *[]BranchBusinessPartner  `json:"BranchBusinessPartner"`
}

type General struct {
	FinInstCountry      string  `json:"FinInstCountry"`
	FinInstCode         string  `json:"FinInstCode"`
	FinInstName         *string `json:"FinInstName"`
	FinInstFullName     *string `json:"FinInstFullName"`
	AddressID           *int    `json:"AddressID"`
	SWIFTCode           *string `json:"SWIFTCode"`
	IsMarkedForDeletion *bool   `json:"IsMarkedForDeletion"`
}

type GeneralBusinessPartner struct {
	FinInstCountry         string `json:"FinInstCountry"`
	FinInstCode            string `json:"FinInstCode"`
	FinInstBusinessPartner int    `json:"FinInstBusinessPartner"`
	ValidityEndDate        string `json:"ValidityEndDate"`
	ValidityStartDate      string `json:"ValidityStartDate"`
	IsMarkedForDeletion    *bool  `json:"IsMarkedForDeletion"`
}

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
