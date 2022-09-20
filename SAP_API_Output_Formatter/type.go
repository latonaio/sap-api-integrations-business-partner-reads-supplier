package sap_api_output_formatter

type BusinessPartner struct {
	ConnectionKey       string `json:"connection_key"`
	Result              bool   `json:"result"`
	RedisKey            string `json:"redis_key"`
	Filepath            string `json:"filepath"`
	Product             string `json:"Product"`
	APISchema           string `json:"api_schema"`
	BusinessPartnerCode string `json:"business_partner_code"`
	Deleted             string `json:"deleted"`
}

type General struct {
	BusinessPartner               string `json:"BusinessPartner"`
	Customer                      string `json:"Customer"`
	Supplier                      string `json:"Supplier"`
	AcademicTitle                 string `json:"AcademicTitle"`
	AuthorizationGroup            string `json:"AuthorizationGroup"`
	BusinessPartnerCategory       string `json:"BusinessPartnerCategory"`
	BusinessPartnerFullName       string `json:"BusinessPartnerFullName"`
	BusinessPartnerGrouping       string `json:"BusinessPartnerGrouping"`
	BusinessPartnerName           string `json:"BusinessPartnerName"`
	CorrespondenceLanguage        string `json:"CorrespondenceLanguage"`
	CreationDate                  string `json:"CreationDate"`
	CreationTime                  string `json:"CreationTime"`
	FirstName                     string `json:"FirstName"`
	Industry                      string `json:"Industry"`
	IsFemale                      bool   `json:"IsFemale"`
	IsMale                        bool   `json:"IsMale"`
	IsNaturalPerson               string `json:"IsNaturalPerson"`
	IsSexUnknown                  bool   `json:"IsSexUnknown"`
	GenderCodeName                string `json:"GenderCodeName"`
	Language                      string `json:"Language"`
	LastChangeDate                string `json:"LastChangeDate"`
	LastChangeTime                string `json:"LastChangeTime"`
	LastName                      string `json:"LastName"`
	OrganizationBPName1           string `json:"OrganizationBPName1"`
	OrganizationBPName2           string `json:"OrganizationBPName2"`
	OrganizationBPName3           string `json:"OrganizationBPName3"`
	OrganizationBPName4           string `json:"OrganizationBPName4"`
	OrganizationFoundationDate    string `json:"OrganizationFoundationDate"`
	OrganizationLiquidationDate   string `json:"OrganizationLiquidationDate"`
	SearchTerm1                   string `json:"SearchTerm1"`
	SearchTerm2                   string `json:"SearchTerm2"`
	AdditionalLastName            string `json:"AdditionalLastName"`
	BirthDate                     string `json:"BirthDate"`
	BusinessPartnerBirthplaceName string `json:"BusinessPartnerBirthplaceName"`
	BusinessPartnerDeathDate      string `json:"BusinessPartnerDeathDate"`
	BusinessPartnerIsBlocked      bool   `json:"BusinessPartnerIsBlocked"`
	BusinessPartnerType           string `json:"BusinessPartnerType"`
	GroupBusinessPartnerName1     string `json:"GroupBusinessPartnerName1"`
	GroupBusinessPartnerName2     string `json:"GroupBusinessPartnerName2"`
	IndependentAddressID          string `json:"IndependentAddressID"`
	MiddleName                    string `json:"MiddleName"`
	NameCountry                   string `json:"NameCountry"`
	PersonFullName                string `json:"PersonFullName"`
	PersonNumber                  string `json:"PersonNumber"`
	IsMarkedForArchiving          bool   `json:"IsMarkedForArchiving"`
	BusinessPartnerIDByExtSystem  string `json:"BusinessPartnerIDByExtSystem"`
	TradingPartner                string `json:"TradingPartner"`
	ToRole                        string `json:"to_BusinessPartnerRole"`
	ToAddress                     string `json:"to_BusinessPartnerAddress"`
	ToBank                        string `json:"to_BusinessPartnerBank"`
	ToSupplier                    string `json:"to_Supplier"`
}

type Role struct {
	BusinessPartner     string `json:"BusinessPartner"`
	BusinessPartnerRole string `json:"BusinessPartnerRole"`
	ValidFrom           string `json:"ValidFrom"`
	ValidTo             string `json:"ValidTo"`
}

type Address struct {
	BusinessPartner   string `json:"BusinessPartner"`
	AddressID         string `json:"AddressID"`
	ValidityStartDate string `json:"ValidityStartDate"`
	ValidityEndDate   string `json:"ValidityEndDate"`
	Country           string `json:"Country"`
	Region            string `json:"Region"`
	StreetName        string `json:"StreetName"`
	CityName          string `json:"CityName"`
	PostalCode        string `json:"PostalCode"`
	Language          string `json:"Language"`
}

type Bank struct {
	BusinessPartner          string `json:"BusinessPartner"`
	BankIdentification       string `json:"BankIdentification"`
	BankCountryKey           string `json:"BankCountryKey"`
	BankName                 string `json:"BankName"`
	BankNumber               string `json:"BankNumber"`
	SWIFTCode                string `json:"SWIFTCode"`
	BankControlKey           string `json:"BankControlKey"`
	BankAccountHolderName    string `json:"BankAccountHolderName"`
	BankAccountName          string `json:"BankAccountName"`
	ValidityStartDate        string `json:"ValidityStartDate"`
	ValidityEndDate          string `json:"ValidityEndDate"`
	IBAN                     string `json:"IBAN"`
	IBANValidityStartDate    string `json:"IBANValidityStartDate"`
	BankAccount              string `json:"BankAccount"`
	BankAccountReferenceText string `json:"BankAccountReferenceText"`
	CollectionAuthInd        bool   `json:"CollectionAuthInd"`
	CityName                 string `json:"CityName"`
	AuthorizationGroup       string `json:"AuthorizationGroup"`
}

type Supplier struct {
	Supplier                    string `json:"Supplier"`
	AuthorizationGroup          string `json:"AuthorizationGroup"`
	CreationDate                string `json:"CreationDate"`
	Customer                    string `json:"Customer"`
	PaymentIsBlockedForSupplier bool   `json:"PaymentIsBlockedForSupplier"`
	PostingIsBlocked            bool   `json:"PostingIsBlocked"`
	PurchasingIsBlocked         bool   `json:"PurchasingIsBlocked"`
	SupplierAccountGroup        string `json:"SupplierAccountGroup"`
	SupplierFullName            string `json:"SupplierFullName"`
	SupplierName                string `json:"SupplierName"`
	BirthDate                   string `json:"BirthDate"`
	DeletionIndicator           bool   `json:"DeletionIndicator"`
	Industry                    string `json:"Industry"`
	IsNaturalPerson             string `json:"IsNaturalPerson"`
	SupplierCorporateGroup      string `json:"SupplierCorporateGroup"`
	SupplierProcurementBlock    string `json:"SupplierProcurementBlock"`
	ToPurchasingOrganization    string `json:"to_SupplierPurchasingOrg"`
	ToCompany                   string `json:"to_SupplierCompany"`
}

type PurchasingOrganization struct {
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
	ToPartnerFunction              string `json:"to_PartnerFunction"`
}

type Company struct {
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
}

type ToRole struct {
	BusinessPartner     string `json:"BusinessPartner"`
	BusinessPartnerRole string `json:"BusinessPartnerRole"`
	ValidFrom           string `json:"ValidFrom"`
	ValidTo             string `json:"ValidTo"`
}

type ToAddress struct {
	BusinessPartner   string `json:"BusinessPartner"`
	AddressID         string `json:"AddressID"`
	ValidityStartDate string `json:"ValidityStartDate"`
	ValidityEndDate   string `json:"ValidityEndDate"`
	Country           string `json:"Country"`
	Region            string `json:"Region"`
	StreetName        string `json:"StreetName"`
	CityName          string `json:"CityName"`
	PostalCode        string `json:"PostalCode"`
	Language          string `json:"Language"`
}

type ToBank struct {
	BusinessPartner          string `json:"BusinessPartner"`
	BankIdentification       string `json:"BankIdentification"`
	BankCountryKey           string `json:"BankCountryKey"`
	BankName                 string `json:"BankName"`
	BankNumber               string `json:"BankNumber"`
	SWIFTCode                string `json:"SWIFTCode"`
	BankControlKey           string `json:"BankControlKey"`
	BankAccountHolderName    string `json:"BankAccountHolderName"`
	BankAccountName          string `json:"BankAccountName"`
	ValidityStartDate        string `json:"ValidityStartDate"`
	ValidityEndDate          string `json:"ValidityEndDate"`
	IBAN                     string `json:"IBAN"`
	IBANValidityStartDate    string `json:"IBANValidityStartDate"`
	BankAccount              string `json:"BankAccount"`
	BankAccountReferenceText string `json:"BankAccountReferenceText"`
	CollectionAuthInd        bool   `json:"CollectionAuthInd"`
	CityName                 string `json:"CityName"`
	AuthorizationGroup       string `json:"AuthorizationGroup"`
}

type ToSupplier struct {
	Supplier                    string `json:"Supplier"`
	AuthorizationGroup          string `json:"AuthorizationGroup"`
	CreationDate                string `json:"CreationDate"`
	Customer                    string `json:"Customer"`
	PaymentIsBlockedForSupplier bool   `json:"PaymentIsBlockedForSupplier"`
	PostingIsBlocked            bool   `json:"PostingIsBlocked"`
	PurchasingIsBlocked         bool   `json:"PurchasingIsBlocked"`
	SupplierAccountGroup        string `json:"SupplierAccountGroup"`
	SupplierFullName            string `json:"SupplierFullName"`
	SupplierName                string `json:"SupplierName"`
	BirthDate                   string `json:"BirthDate"`
	DeletionIndicator           bool   `json:"DeletionIndicator"`
	Industry                    string `json:"Industry"`
	IsNaturalPerson             string `json:"IsNaturalPerson"`
	SupplierCorporateGroup      string `json:"SupplierCorporateGroup"`
	SupplierProcurementBlock    string `json:"SupplierProcurementBlock"`
	ToPurchasingOrganization    string `json:"to_SupplierPurchasingOrg"`
	ToCompany                   string `json:"to_SupplierCompany"`
}

type ToPurchasingOrganization struct {
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
	ToPartnerFunction              string `json:"to_PartnerFunction"`
}

type ToPartnerFunction struct {
	Supplier               string `json:"Supplier"`
	PurchasingOrganization string `json:"PurchasingOrganization"`
	Plant                  string `json:"Plant"`
	PartnerFunction        string `json:"PartnerFunction"`
	PartnerCounter         string `json:"PartnerCounter"`
	DefaultPartner         bool   `json:"DefaultPartner"`
	CreationDate           string `json:"CreationDate"`
	ReferenceSupplier      string `json:"ReferenceSupplier"`
	AuthorizationGroup     string `json:"AuthorizationGroup"`
}

type ToCompany struct {
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
}
