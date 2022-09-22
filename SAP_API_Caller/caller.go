package sap_api_caller

import (
	"fmt"
	"io/ioutil"
	sap_api_output_formatter "sap-api-integrations-business-partner-reads-supplier/SAP_API_Output_Formatter"
	"strings"
	"sync"

	sap_api_request_client_header_setup "github.com/latonaio/sap-api-request-client-header-setup"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
)

type SAPAPICaller struct {
	baseURL         string
	sapClientNumber string
	requestClient   *sap_api_request_client_header_setup.SAPRequestClient
	log             *logger.Logger
}

func NewSAPAPICaller(baseUrl, sapClientNumber string, requestClient *sap_api_request_client_header_setup.SAPRequestClient, l *logger.Logger) *SAPAPICaller {
	return &SAPAPICaller{
		baseURL:         baseUrl,
		requestClient:   requestClient,
		sapClientNumber: sapClientNumber,
		log:             l,
	}
}

func (c *SAPAPICaller) AsyncGetBPSupplier(businessPartner, businessPartnerRole, addressID, bankCountryKey, bankNumber, bPName, supplier, purchasingOrganization, companyCode string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "General":
			func() {
				c.General(businessPartner)
				wg.Done()
			}()
		case "Role":
			func() {
				c.Role(businessPartner, businessPartnerRole)
				wg.Done()
			}()
		case "Address":
			func() {
				c.Address(businessPartner, addressID)
				wg.Done()
			}()
		case "Bank":
			func() {
				c.Bank(businessPartner, bankCountryKey, bankNumber)
				wg.Done()
			}()
		case "BPName":
			func() {
				c.BPName(bPName)
				wg.Done()
			}()
		case "Supplier":
			func() {
				c.Supplier(supplier)
				wg.Done()
			}()
		case "PurchasingOrganization":
			func() {
				c.PurchasingOrganization(supplier, purchasingOrganization)
				wg.Done()
			}()
		case "Company":
			func() {
				c.Company(supplier, companyCode)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}

func (c *SAPAPICaller) General(businessPartner string) {
	generalData, err := c.callBPSupplierSrvAPIRequirementGeneral("A_BusinessPartner", businessPartner)
	if err != nil {
		c.log.Error(err)
	} else {
		c.log.Info(generalData)
	}

	roleData, err := c.callToRole(generalData[0].ToRole)
	if err != nil {
		c.log.Error(err)
	} else {
		c.log.Info(roleData)
	}

	addressData, err := c.callToAddress(generalData[0].ToAddress)
	if err != nil {
		c.log.Error(err)
	} else {
		c.log.Info(addressData)
	}

	bankData, err := c.callToBank(generalData[0].ToBank)
	if err != nil {
		c.log.Error(err)
	} else {
		c.log.Info(bankData)
	}

	supplierData, err := c.callToSupplier(generalData[0].ToSupplier)
	if err != nil {
		c.log.Error(err)
	} else {
		c.log.Info(supplierData)
	}

	purchasingOrganizationData, err := c.callToPurchasingOrganization(supplierData.ToPurchasingOrganization)
	if err != nil {
		c.log.Error(err)
	} else {
		c.log.Info(purchasingOrganizationData)
	}

	partnerFunctionData, err := c.callToPartnerFunction(purchasingOrganizationData[0].ToPartnerFunction)
	if err != nil {
		c.log.Error(err)
	} else {
		c.log.Info(partnerFunctionData)
	}

	companyData, err := c.callToCompany(supplierData.ToCompany)
	if err != nil {
		c.log.Error(err)
	} else {
		c.log.Info(companyData)
	}
}

func (c *SAPAPICaller) callBPSupplierSrvAPIRequirementGeneral(api, businessPartner string) ([]sap_api_output_formatter.General, error) {
	url := strings.Join([]string{c.baseURL, "API_BUSINESS_PARTNER", api}, "/")
	param := c.getQueryWithGeneral(map[string]string{}, businessPartner)

	resp, err := c.requestClient.Request("GET", url, param, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToGeneral(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToRole(url string) ([]sap_api_output_formatter.ToRole, error) {
	resp, err := c.requestClient.Request("GET", url, map[string]string{}, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToRole(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToAddress(url string) ([]sap_api_output_formatter.ToAddress, error) {
	resp, err := c.requestClient.Request("GET", url, map[string]string{}, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToAddress(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToBank(url string) ([]sap_api_output_formatter.ToBank, error) {
	resp, err := c.requestClient.Request("GET", url, map[string]string{}, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToBank(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToSupplier(url string) (*sap_api_output_formatter.ToSupplier, error) {
	resp, err := c.requestClient.Request("GET", url, map[string]string{}, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToSupplier(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToPurchasingOrganization(url string) ([]sap_api_output_formatter.ToPurchasingOrganization, error) {
	resp, err := c.requestClient.Request("GET", url, map[string]string{}, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToPurchasingOrganization(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToPartnerFunction(url string) ([]sap_api_output_formatter.ToPartnerFunction, error) {
	resp, err := c.requestClient.Request("GET", url, map[string]string{}, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToPartnerFunction(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToCompany(url string) ([]sap_api_output_formatter.ToCompany, error) {
	resp, err := c.requestClient.Request("GET", url, map[string]string{}, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToCompany(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) Role(businessPartner, businessPartnerRole string) {
	data, err := c.callBPSupplierSrvAPIRequirementRole("A_BusinessPartnerRole", businessPartner, businessPartnerRole)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)
}

func (c *SAPAPICaller) callBPSupplierSrvAPIRequirementRole(api, businessPartner, businessPartnerRole string) ([]sap_api_output_formatter.Role, error) {
	url := strings.Join([]string{c.baseURL, "API_BUSINESS_PARTNER", api}, "/")

	param := c.getQueryWithRole(map[string]string{}, businessPartner, businessPartnerRole)

	resp, err := c.requestClient.Request("GET", url, param, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToRole(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) Address(businessPartner, addressID string) {
	data, err := c.callBPSupplierSrvAPIRequirementAddress("A_BusinessPartnerAddress", businessPartner, addressID)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)
}

func (c *SAPAPICaller) callBPSupplierSrvAPIRequirementAddress(api, businessPartner, addressID string) ([]sap_api_output_formatter.Address, error) {
	url := strings.Join([]string{c.baseURL, "API_BUSINESS_PARTNER", api}, "/")

	param := c.getQueryWithAddress(map[string]string{}, businessPartner, addressID)

	resp, err := c.requestClient.Request("GET", url, param, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToAddress(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) Bank(businessPartner, bankCountryKey, bankNumber string) {
	data, err := c.callBPSupplierSrvAPIRequirementBank("A_BusinessPartnerBank", businessPartner, bankCountryKey, bankNumber)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)
}

func (c *SAPAPICaller) callBPSupplierSrvAPIRequirementBank(api, businessPartner, bankCountryKey, bankNumber string) ([]sap_api_output_formatter.Bank, error) {
	url := strings.Join([]string{c.baseURL, "API_BUSINESS_PARTNER", api}, "/")

	param := c.getQueryWithBank(map[string]string{}, businessPartner, bankCountryKey, bankNumber)

	resp, err := c.requestClient.Request("GET", url, param, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToBank(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) BPName(bPName string) {
	data, err := c.callBPSupplierSrvAPIRequirementBPName("A_BusinessPartner", bPName)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)
}

func (c *SAPAPICaller) callBPSupplierSrvAPIRequirementBPName(api, bPName string) ([]sap_api_output_formatter.General, error) {
	url := strings.Join([]string{c.baseURL, "API_BUSINESS_PARTNER", api}, "/")

	param := c.getQueryWithBPName(map[string]string{}, bPName)

	resp, err := c.requestClient.Request("GET", url, param, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToGeneral(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) Supplier(supplier string) {
	supplierData, err := c.callBPSupplierSrvAPIRequirementSupplier("A_Supplier", supplier)

	if err != nil {
		c.log.Error(err)
	} else {
		c.log.Info(supplierData)
	}

	purchasingOrganizationData, err := c.callToPurchasingOrganization(supplierData[0].ToPurchasingOrganization)
	if err != nil {
		c.log.Error(err)
	} else {
		c.log.Info(purchasingOrganizationData)
	}

	partnerFunctionData, err := c.callToPartnerFunction(purchasingOrganizationData[0].ToPartnerFunction)
	if err != nil {
		c.log.Error(err)
	} else {
		c.log.Info(partnerFunctionData)
	}

	companyData, err := c.callToCompany(supplierData[0].ToCompany)
	if err != nil {
		c.log.Error(err)
	} else {
		c.log.Info(companyData)
	}

}

func (c *SAPAPICaller) callBPSupplierSrvAPIRequirementSupplier(api, supplier string) ([]sap_api_output_formatter.Supplier, error) {
	url := strings.Join([]string{c.baseURL, "API_BUSINESS_PARTNER", api}, "/")

	param := c.getQueryWithSupplier(map[string]string{}, supplier)

	resp, err := c.requestClient.Request("GET", url, param, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToSupplier(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) PurchasingOrganization(supplier, purchasingOrganization string) {
	purchasingOrganizationData, err := c.callBPSupplierSrvAPIRequirementPurchasingOrganization("A_SupplierPurchasingOrg", supplier, purchasingOrganization)
	if err != nil {
		c.log.Error(err)
	} else {
		c.log.Info(purchasingOrganizationData)
	}

	partnerFunctionData, err := c.callToPartnerFunction(purchasingOrganizationData[0].ToPartnerFunction)
	if err != nil {
		c.log.Error(err)
	} else {
		c.log.Info(partnerFunctionData)
	}
}

func (c *SAPAPICaller) callBPSupplierSrvAPIRequirementPurchasingOrganization(api, supplier, purchasingOrganization string) ([]sap_api_output_formatter.PurchasingOrganization, error) {
	url := strings.Join([]string{c.baseURL, "API_BUSINESS_PARTNER", api}, "/")

	param := c.getQueryWithPurchasingOrganization(map[string]string{}, supplier, purchasingOrganization)

	resp, err := c.requestClient.Request("GET", url, param, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToPurchasingOrganization(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) Company(supplier, companyCode string) {
	data, err := c.callBPSupplierSrvAPIRequirementCompany("A_SupplierCompany", supplier, companyCode)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)
}

func (c *SAPAPICaller) callBPSupplierSrvAPIRequirementCompany(api, supplier, companyCode string) ([]sap_api_output_formatter.Company, error) {
	url := strings.Join([]string{c.baseURL, "API_BUSINESS_PARTNER", api}, "/")

	param := c.getQueryWithCompany(map[string]string{}, supplier, companyCode)

	resp, err := c.requestClient.Request("GET", url, param, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToCompany(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) getQueryWithGeneral(params map[string]string, businessPartner string) map[string]string {
	if len(params) == 0 {
		params = make(map[string]string, 1)
	}
	params["$filter"] = fmt.Sprintf("BusinessPartner eq '%s'", businessPartner)
	return params
}

func (c *SAPAPICaller) getQueryWithRole(params map[string]string, businessPartner, businessPartnerRole string) map[string]string {
	if len(params) == 0 {
		params = make(map[string]string, 1)
	}
	params["$filter"] = fmt.Sprintf("BusinessPartner eq '%s' and BusinessPartnerRole eq '%s'", businessPartner, businessPartnerRole)
	return params
}

func (c *SAPAPICaller) getQueryWithAddress(params map[string]string, businessPartner, addressID string) map[string]string {
	if len(params) == 0 {
		params = make(map[string]string, 1)
	}
	params["$filter"] = fmt.Sprintf("BusinessPartner eq '%s' and AddressID eq '%s'", businessPartner, addressID)
	return params
}

func (c *SAPAPICaller) getQueryWithBank(params map[string]string, businessPartner, bankCountryKey, bankNumber string) map[string]string {
	if len(params) == 0 {
		params = make(map[string]string, 1)
	}
	params["$filter"] = fmt.Sprintf("BusinessPartner eq '%s' and BankCountryKey eq '%s' and BankNumber eq '%s'", businessPartner, bankCountryKey, bankNumber)
	return params
}

func (c *SAPAPICaller) getQueryWithBPName(params map[string]string, bPName string) map[string]string {
	if len(params) == 0 {
		params = make(map[string]string, 1)
	}
	params["$filter"] = fmt.Sprintf("BPName eq '%s'", bPName)
	return params
}

func (c *SAPAPICaller) getQueryWithSupplier(params map[string]string, supplier string) map[string]string {
	if len(params) == 0 {
		params = make(map[string]string, 1)
	}
	params["$filter"] = fmt.Sprintf("Supplier eq '%s'", supplier)
	return params
}

func (c *SAPAPICaller) getQueryWithPurchasingOrganization(params map[string]string, supplier, purchasingOrganization string) map[string]string {
	if len(params) == 0 {
		params = make(map[string]string, 1)
	}
	params["$filter"] = fmt.Sprintf("Supplier eq '%s' and PurchasingOrganization eq '%s'", supplier, purchasingOrganization)
	return params
}

func (c *SAPAPICaller) getQueryWithCompany(params map[string]string, supplier, companyCode string) map[string]string {
	if len(params) == 0 {
		params = make(map[string]string, 1)
	}
	params["$filter"] = fmt.Sprintf("Supplier eq '%s' and CompanyCode eq '%s'", supplier, companyCode)
	return params
}
