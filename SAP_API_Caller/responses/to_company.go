package responses

type ToCompany struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			Supplier                    string `json:"Supplier"`
			CompanyCode                 string `json:"CompanyCode"`
			PaymentBlockingReason       string `json:"PaymentBlockingReason"`
			PaymentMethodsList          string `json:"PaymentMethodsList"`
			PaymentTerms                string `json:"PaymentTerms"`
			ClearCustomerSupplier       bool   `json:"ClearCustomerSupplier"`
			HouseBank                   string `json:"HouseBank"`
			ReconciliationAccount       string `json:"ReconciliationAccount"`
			SupplierIsBlockedForPosting bool   `json:"SupplierIsBlockedForPosting"`
			DeletionIndicator           bool   `json:"DeletionIndicator"`
		} `json:"results"`
	} `json:"d"`
}
