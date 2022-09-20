package responses

type PurchasingOrganization struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			Supplier                       string `json:"Supplier"`
			PurchasingOrganization         string `json:"PurchasingOrganization"`
			IncotermsClassification        string `json:"IncotermsClassification"`
			InvoiceIsGoodsReceiptBased     bool   `json:"InvoiceIsGoodsReceiptBased"`
			PaymentTerms                   string `json:"PaymentTerms"`
			PurOrdAutoGenerationIsAllowed  bool   `json:"PurOrdAutoGenerationIsAllowed"`
			PurchaseOrderCurrency          string `json:"PurchaseOrderCurrency"`
			PurchasingGroup                string `json:"PurchasingGroup"`
			ShippingCondition              string `json:"ShippingCondition"`
			SupplierPhoneNumber            string `json:"SupplierPhoneNumber"`
			SupplierRespSalesPersonName    string `json:"SupplierRespSalesPersonName"`
			PurchasingIsBlockedForSupplier bool   `json:"PurchasingIsBlockedForSupplier"`
			DeletionIndicator              bool   `json:"DeletionIndicator"`
			ToPartnerFunction              struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"to_PartnerFunction"`
		} `json:"results"`
	} `json:"d"`
}
