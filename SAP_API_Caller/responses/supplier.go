package responses

type Supplier struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			Supplier                       string      `json:"Supplier"`
			AuthorizationGroup             string      `json:"AuthorizationGroup"`
			CreationDate                   string      `json:"CreationDate"`
			Customer                       string      `json:"Customer"`
			PaymentIsBlockedForSupplier    bool        `json:"PaymentIsBlockedForSupplier"`
			PostingIsBlocked               bool        `json:"PostingIsBlocked"`
			PurchasingIsBlocked            bool        `json:"PurchasingIsBlocked"`
			SupplierAccountGroup           string      `json:"SupplierAccountGroup"`
			SupplierFullName               string      `json:"SupplierFullName"`
			SupplierName                   string      `json:"SupplierName"`
			BirthDate                      string      `json:"BirthDate"`
			DeletionIndicator              bool        `json:"DeletionIndicator"`
			Industry                       string      `json:"Industry"`
			IsNaturalPerson                string      `json:"IsNaturalPerson"`
			SupplierCorporateGroup         string      `json:"SupplierCorporateGroup"`
			SupplierProcurementBlock       string      `json:"SupplierProcurementBlock"`
			ToPurchasingOrganization struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"to_SupplierPurchasingOrg"`
			ToCompany              struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"to_SupplierCompany"`
		} `json:"results"`
	} `json:"d"`
}
