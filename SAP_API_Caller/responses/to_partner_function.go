package responses

type ToPartnerFunction struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			Supplier               string `json:"Supplier"`
			PurchasingOrganization string `json:"PurchasingOrganization"`
			Plant                  string `json:"Plant"`
			PartnerFunction        string `json:"PartnerFunction"`
			PartnerCounter         string `json:"PartnerCounter"`
			DefaultPartner         bool   `json:"DefaultPartner"`
			CreationDate           string `json:"CreationDate"`
			ReferenceSupplier      string `json:"ReferenceSupplier"`
			AuthorizationGroup     string `json:"AuthorizationGroup"`
		} `json:"results"`
	} `json:"d"`
}
