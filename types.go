package iplicit

var (
	IncludeCustomer           Include = "customer"
	IncludeSupplier           Include = "supplier"
	IncludeResource           Include = "resource"
	IncludeContact            Include = "contact"
	IncludeContacts           Include = "contacts"
	IncludeDefaultBankDetails Include = "defaultBankDetails"
	IncludeAllBankDetails     Include = "allBankDetails"
)

type Include string

func (i Include) String() string {
	return string(i)
}

var (
	IncludeExt IncludeForProduct = "ext"
)

type IncludeForProduct string

func (i IncludeForProduct) String() string {
	return string(i)
}

// TokenResponse

type TokenResponsePost struct {
	TokenDue     string `json:"tokenDue"`
	SessionToken string `json:"sessionToken"`
	Domain       string `json:"domain"`
	ApiVersion   string `json:"apiVer"`
}

// ContactAccount

type ContactAccountGet struct {
	ID                      string                `json:"id"`
	Code                    string                `json:"code"`
	IntRef                  string                `json:"intRef"`
	Description             string                `json:"description"`
	CompanyNo               string                `json:"companyNo"`
	TaxNo                   string                `json:"taxNo"`
	LegacyRef               string                `json:"legacyRef"`
	TheirRef                string                `json:"theirRef"`
	CountryCode             string                `json:"countryCode"`
	ContactClassificationID string                `json:"contactClassificationId"`
	IsParent                bool                  `json:"isParent"`
	ParentContactAccountID  string                `json:"parentContactAccountId"`
	Customer                CustomerGet           `json:"customer,omitzero"`
	Supplier                SupplierGet           `json:"supplier,omitzero"`
	Resource                ResourceGet           `json:"resource,omitzero"`
	Contact                 ContactGet            `json:"contact,omitzero"`
	Contacts                []ContactsGet         `json:"contacts"`
	DefaultBankDetails      DefaultBankDetailsGet `json:"defaultBankDetails,omitzero"`
	BankDetails             []BankDetailsGet      `json:"bankDetails"`
	LastModified            DateTime              `json:"lastModified"`
	LastModifiedBy          string                `json:"lastModifiedBy"`
	HasNotes                bool                  `json:"hasNotes"`
	HasAttachments          bool                  `json:"hasAttachments"`
	CreatedDate             DateTime              `json:"createdDate"`
	CreatedBy               string                `json:"createdBy"`
}
type ExtGet struct {
	AdditionalProp1 string `json:"additionalProp1"`
	AdditionalProp2 string `json:"additionalProp2"`
	AdditionalProp3 string `json:"additionalProp3"`
}
type CustomerGet struct {
	IsActive                   bool     `json:"isActive"`
	ContactGroupCustomerID     string   `json:"contactGroupCustomerId"`
	Currency                   string   `json:"currency"`
	DiscountRate               int      `json:"discountRate"`
	PaymentMethodID            string   `json:"paymentMethodId"`
	PayTermID                  string   `json:"payTermId"`
	ProductPriceBandID         string   `json:"productPriceBandId"`
	ReminderGroupID            string   `json:"reminderGroupId"`
	ReceiveInvoices            bool     `json:"receiveInvoices"`
	ReceiveReminders           bool     `json:"receiveReminders"`
	ReceiveStatements          bool     `json:"receiveStatements"`
	TaxBandID                  string   `json:"taxBandId"`
	TaxBandFixed               bool     `json:"taxBandFixed"`
	TaxAuthorityID             string   `json:"taxAuthorityId"`
	IsMultiTaxAuthority        bool     `json:"isMultiTaxAuthority"`
	MultiTaxAuthorityIds       []string `json:"multiTaxAuthorityIds"`
	DefaultDeliveryAddressID   string   `json:"defaultDeliveryAddressId"`
	DefaultInvoiceAddressID    string   `json:"defaultInvoiceAddressId"`
	StockLocationID            string   `json:"stockLocationId"`
	LastStatementDate          DateTime `json:"lastStatementDate"`
	CreditLimit                float64  `json:"creditLimit"`
	CreditScore                float64  `json:"creditScore"`
	CreditReportDate           DateTime `json:"creditReportDate"`
	RiskLevel                  string   `json:"riskLevel"`
	IgnoreDeferredIncome       bool     `json:"ignoreDeferredIncome"`
	IgnoreDeposit              bool     `json:"ignoreDeposit"`
	IgnoreSchedule             bool     `json:"ignoreSchedule"`
	IsHold                     bool     `json:"isHold"`
	IsStop                     bool     `json:"isStop"`
	ProjectID                  string   `json:"projectId"`
	ApplyDomesticReverseCharge bool     `json:"applyDomesticReverseCharge"`
	OrderRequired              bool     `json:"orderRequired"`
	CostCentreID               string   `json:"costCentreId"`
	CostCentreFixed            bool     `json:"costCentreFixed"`
	DepartmentID               string   `json:"departmentId"`
	DepartmentFixed            bool     `json:"departmentFixed"`
	WarningMessage             string   `json:"warningMessage"`
	Ext                        ExtGet   `json:"ext"`
}
type SupplierExtGet struct {
	AdditionalProp1 string `json:"additionalProp1"`
	AdditionalProp2 string `json:"additionalProp2"`
	AdditionalProp3 string `json:"additionalProp3"`
}
type SupplierGet struct {
	IsActive                         bool           `json:"isActive"`
	ContactGroupSupplierID           string         `json:"contactGroupSupplierId"`
	Currency                         string         `json:"currency"`
	CreditLimit                      int            `json:"creditLimit"`
	DiscountRate                     int            `json:"discountRate"`
	PaymentMethodID                  string         `json:"paymentMethodId"`
	PaymentRecipientContactAccountID string         `json:"paymentRecipientContactAccountId"`
	PayTermID                        string         `json:"payTermId"`
	ProductPriceBandID               string         `json:"productPriceBandId"`
	TaxBandID                        string         `json:"taxBandId"`
	TaxBandFixed                     bool           `json:"taxBandFixed"`
	TaxAuthorityID                   string         `json:"taxAuthorityId"`
	IsMultiTaxAuthority              bool           `json:"isMultiTaxAuthority"`
	SupplierMultiTaxAuthorityIds     []string       `json:"multiTaxAuthorityIds"`
	AllowInvoiceDuplication          bool           `json:"allowInvoiceDuplication"`
	AllowSelfBilling                 bool           `json:"allowSelfBilling"`
	SendRemittance                   bool           `json:"sendRemittance"`
	ProjectID                        string         `json:"projectId"`
	ApplyDomesticReverseCharge       bool           `json:"applyDomesticReverseCharge"`
	CostCentreID                     string         `json:"costCentreId"`
	CostCentreFixed                  bool           `json:"costCentreFixed"`
	DepartmentID                     string         `json:"departmentId"`
	DepartmentFixed                  bool           `json:"departmentFixed"`
	ProductID                        string         `json:"productId"`
	AccountID                        string         `json:"accountId"`
	LegalEntityID                    string         `json:"legalEntityId"`
	BankAccountID                    string         `json:"bankAccountId"`
	DefaultDeliveryAddressID         string         `json:"defaultDeliveryAddressId"`
	DefaultInvoiceAddressID          string         `json:"defaultInvoiceAddressId"`
	ResponsibleResourceID            string         `json:"responsibleResourceId"`
	StockLocationID                  string         `json:"stockLocationId"`
	IgnorePrepayment                 bool           `json:"ignorePrepayment"`
	IsHold                           bool           `json:"isHold"`
	IsOcrGross                       bool           `json:"isOcrGross"`
	IsResale                         bool           `json:"isResale"`
	IsStop                           bool           `json:"isStop"`
	OcrAliases                       string         `json:"ocrAliases"`
	OcrCombineLineItems              bool           `json:"ocrCombineLineItems"`
	OcrDisableLineItems              bool           `json:"ocrDisableLineItems"`
	WarningMessage                   string         `json:"warningMessage"`
	SupplierExt                      SupplierExtGet `json:"ext"`
}
type TimesheetProductsGet struct {
	ProductID        string   `json:"productId"`
	ContactAccountID string   `json:"contactAccountId"`
	DateFrom         DateTime `json:"dateFrom"`
	DateTo           DateTime `json:"dateTo"`
	CostRate         int      `json:"costRate"`
	RechargeRate     int      `json:"rechargeRate"`
}
type ExpenseProductsGet struct {
	ProductID string `json:"productId"`
}
type ResourceEmploymentsGet struct {
	DateFrom                      DateTime `json:"dateFrom"`
	DateTo                        DateTime `json:"dateTo"`
	LegalEntityID                 string   `json:"legalEntityId"`
	EmploymentTypeID              string   `json:"employmentTypeId"`
	JobTitle                      string   `json:"jobTitle"`
	JobGrade                      string   `json:"jobGrade"`
	PaymentFrequencyID            string   `json:"paymentFrequencyId"`
	PreviousEmployer              string   `json:"previousEmployer"`
	TerminationNoticeDays         int      `json:"terminationNoticeDays"`
	ExcludeFromIntermediaryReport bool     `json:"excludeFromIntermediaryReport"`
	ExternalReference             string   `json:"externalReference"`
}
type ResourceUmbrellasGet struct {
	UmbrellaContactAccountID string   `json:"umbrellaContactAccountId"`
	DateFrom                 DateTime `json:"dateFrom"`
	DateTo                   DateTime `json:"dateTo"`
}
type ResourceExtGet struct {
	AdditionalProp1 string `json:"additionalProp1"`
	AdditionalProp2 string `json:"additionalProp2"`
	AdditionalProp3 string `json:"additionalProp3"`
}
type ResourceGet struct {
	IsActive                bool                     `json:"isActive"`
	ResourceGroupID         string                   `json:"resourceGroupId"`
	DepartmentID            string                   `json:"departmentId"`
	LocationID              string                   `json:"locationId"`
	NationalInsuranceNo     string                   `json:"nationalInsuranceNo"`
	Currency                string                   `json:"currency"`
	AbsenceGroupID          string                   `json:"absenceGroupId"`
	CostCentreID            string                   `json:"costCentreId"`
	ManagerResourceID       string                   `json:"managerResourceId"`
	UniqueTaxpayerReference string                   `json:"uniqueTaxpayerReference"`
	IsSalesperson           bool                     `json:"isSalesperson"`
	IsBuyer                 bool                     `json:"isBuyer"`
	IsRequestor             bool                     `json:"isRequestor"`
	AllowExpense            bool                     `json:"allowExpense"`
	AllowTask               bool                     `json:"allowTask"`
	AllowTimesheet          bool                     `json:"allowTimesheet"`
	HasEmployment           bool                     `json:"hasEmployment"`
	HasUmbrella             bool                     `json:"hasUmbrella"`
	PaymentMethodID         string                   `json:"paymentMethodId"`
	PaymentTermsID          string                   `json:"paymentTermsId"`
	SendRemittance          bool                     `json:"sendRemittance"`
	TimesheetProducts       []TimesheetProductsGet   `json:"timesheetProducts"`
	ExpenseProducts         []ExpenseProductsGet     `json:"expenseProducts"`
	ResourceEmployments     []ResourceEmploymentsGet `json:"resourceEmployments"`
	ResourceUmbrellas       []ResourceUmbrellasGet   `json:"resourceUmbrellas"`
	ResourceExt             ResourceExtGet           `json:"ext"`
}
type EmailsGet struct {
	Type  string `json:"type"`
	Email string `json:"email"`
	ID    string `json:"id"`
}
type PhonesGet struct {
	Type  string `json:"type"`
	Phone string `json:"phone"`
	ID    string `json:"id"`
}
type AddressesGet struct {
	Type        string `json:"type"`
	Address     string `json:"address"`
	PostCode    string `json:"postCode"`
	City        string `json:"city"`
	County      string `json:"county"`
	CountryCode string `json:"countryCode"`
	Description string `json:"description"`
	ID          string `json:"id"`
}
type ContactGet struct {
	ID              string         `json:"id"`
	IntRef          string         `json:"intRef"`
	ContactType     string         `json:"contactType"`
	Description     string         `json:"description"`
	CompanyName     string         `json:"companyName"`
	Title           string         `json:"title"`
	FirstName       string         `json:"firstName"`
	MiddleName      string         `json:"middleName"`
	LastName        string         `json:"lastName"`
	JobTitle        string         `json:"jobTitle"`
	ParentContactID string         `json:"parentContactId"`
	LegacyRef       string         `json:"legacyRef"`
	IsActive        bool           `json:"isActive"`
	LastModified    DateTime       `json:"lastModified"`
	LastModifiedBy  string         `json:"lastModifiedBy"`
	HasNotes        bool           `json:"hasNotes"`
	HasAttachments  bool           `json:"hasAttachments"`
	Emails          []EmailsGet    `json:"emails"`
	Phones          []PhonesGet    `json:"phones"`
	Addresses       []AddressesGet `json:"addresses"`
}
type ContactsEmailsGet struct {
	Type  string `json:"type"`
	Email string `json:"email"`
	ID    string `json:"id"`
}
type ContactsPhonesGet struct {
	Type  string `json:"type"`
	Phone string `json:"phone"`
	ID    string `json:"id"`
}
type ContactsAddressesGet struct {
	Type        string `json:"type"`
	Address     string `json:"address"`
	PostCode    string `json:"postCode"`
	City        string `json:"city"`
	County      string `json:"county"`
	CountryCode string `json:"countryCode"`
	Description string `json:"description"`
	ID          string `json:"id"`
}
type ContactsGet struct {
	ID                string                 `json:"id"`
	IntRef            string                 `json:"intRef"`
	ContactType       string                 `json:"contactType"`
	Description       string                 `json:"description"`
	CompanyName       string                 `json:"companyName"`
	Title             string                 `json:"title"`
	FirstName         string                 `json:"firstName"`
	MiddleName        string                 `json:"middleName"`
	LastName          string                 `json:"lastName"`
	JobTitle          string                 `json:"jobTitle"`
	ParentContactID   string                 `json:"parentContactId"`
	LegacyRef         string                 `json:"legacyRef"`
	IsActive          bool                   `json:"isActive"`
	LastModified      DateTime               `json:"lastModified"`
	LastModifiedBy    string                 `json:"lastModifiedBy"`
	HasNotes          bool                   `json:"hasNotes"`
	HasAttachments    bool                   `json:"hasAttachments"`
	ContactsEmails    []ContactsEmailsGet    `json:"emails"`
	ContactsPhones    []ContactsPhonesGet    `json:"phones"`
	ContactsAddresses []ContactsAddressesGet `json:"addresses"`
}
type DefaultBankDetailsGet struct {
	Iban           string   `json:"iban"`
	AccountNo      string   `json:"accountNo"`
	SortCode       string   `json:"sortCode"`
	Bic            string   `json:"bic"`
	RoutingCode    string   `json:"routingCode"`
	Currency       string   `json:"currency"`
	LegacyRef      string   `json:"legacyRef"`
	BankID         string   `json:"bankId"`
	BankRef        string   `json:"bankRef"`
	CountryCode    string   `json:"countryCode"`
	BankAddress    string   `json:"bankAddress"`
	AccountName    string   `json:"accountName"`
	ID             string   `json:"id"`
	LastModified   DateTime `json:"lastModified"`
	LastModifiedBy string   `json:"lastModifiedBy"`
}
type BankDetailsGet struct {
	Iban           string   `json:"iban"`
	AccountNo      string   `json:"accountNo"`
	SortCode       string   `json:"sortCode"`
	Bic            string   `json:"bic"`
	RoutingCode    string   `json:"routingCode"`
	Currency       string   `json:"currency"`
	LegacyRef      string   `json:"legacyRef"`
	BankID         string   `json:"bankId"`
	BankRef        string   `json:"bankRef"`
	CountryCode    string   `json:"countryCode"`
	BankAddress    string   `json:"bankAddress"`
	AccountName    string   `json:"accountName"`
	ID             string   `json:"id"`
	LastModified   DateTime `json:"lastModified"`
	LastModifiedBy string   `json:"lastModifiedBy"`
}

// Product

type ProductResponseGet struct {
	ID             string        `json:"id"`
	Code           string        `json:"code"`
	Description    string        `json:"description"`
	ProductGroupID string        `json:"productGroupId"`
	ProductType    string        `json:"productType"`
	UomID          string        `json:"uomId"`
	UomGroupID     string        `json:"uomGroupId"`
	LegacyRef      string        `json:"legacyRef"`
	Purchase       Purchase      `json:"purchase,omitzero"`
	Sale           Sale          `json:"sale,omitzero"`
	LastModified   DateTime      `json:"lastModified,omitzero"`
	LastModifiedBy string        `json:"lastModifiedBy,omitzero"`
	IsPurchase     bool          `json:"isPurchase"`
	IsSale         bool          `json:"isSale"`
	Ext            ExtProductGet `json:"ext,omitzero"`
}
type SetupPost struct {
	Attribute        string `json:"attribute"`
	Default          string `json:"default"`
	ReadOnly         bool   `json:"readOnly"`
	Mandatory        bool   `json:"mandatory"`
	DefaultParameter string `json:"defaultParameter"`
}
type Purchase struct {
	IsActive  bool        `json:"isActive"`
	AccountID string      `json:"accountId"`
	SetupPost []SetupPost `json:"setupPost"`
}
type SaleSetupPost struct {
	Attribute        string `json:"attribute"`
	Default          string `json:"default"`
	ReadOnly         bool   `json:"readOnly"`
	Mandatory        bool   `json:"mandatory"`
	DefaultParameter string `json:"defaultParameter"`
}
type Sale struct {
	IsActive      bool            `json:"isActive"`
	AccountID     string          `json:"accountId"`
	SaleSetupPost []SaleSetupPost `json:"setupPost"`
}
type ExtProductGet struct {
	AdditionalProp1 string `json:"additionalProp1"`
	AdditionalProp2 string `json:"additionalProp2"`
	AdditionalProp3 string `json:"additionalProp3"`
}

// ContactAccountPost
type ContactAccountPost struct {
	Description             string            `json:"description,omitempty"`
	Code                    string            `json:"code,omitempty"`
	IntRef                  string            `json:"intRef,omitempty"`
	Customer                CustomerPost      `json:"customer,omitzero"`
	Supplier                SupplierPost      `json:"supplier,omitzero"`
	Resource                ResourcePost      `json:"resource,omitzero"`
	CompanyNo               string            `json:"companyNo,omitempty"`
	TaxNo                   string            `json:"taxNo,omitempty"`
	LegacyRef               string            `json:"legacyRef,omitempty"`
	TheirRef                string            `json:"theirRef,omitempty"`
	CountryCode             string            `json:"countryCode,omitempty"`
	ContactClassificationID string            `json:"contactClassificationId,omitempty"`
	IsParent                bool              `json:"isParent,omitempty"`
	ParentContactAccountID  string            `json:"parentContactAccountId,omitempty"`
	Contact                 ContactPost       `json:"contact,omitzero"`
	Contacts                []ContactsPost    `json:"contacts,omitempty"`
	BankDetails             []BankDetailsPost `json:"bankDetails,omitempty"`
}

type ExtPost struct {
	AdditionalProp1 string `json:"additionalProp1,omitempty"`
	AdditionalProp2 string `json:"additionalProp2,omitempty"`
	AdditionalProp3 string `json:"additionalProp3,omitempty"`
}

type CustomerPost struct {
	ContactGroupCustomerID     string   `json:"contactGroupCustomerId,omitempty"`
	Currency                   string   `json:"currency,omitempty"`
	CreditLimit                int      `json:"creditLimit,omitempty"`
	DiscountRate               int      `json:"discountRate,omitempty"`
	PaymentMethodID            string   `json:"paymentMethodId,omitempty"`
	PayTermID                  string   `json:"payTermId,omitempty"`
	ProductPriceBandID         string   `json:"productPriceBandId,omitempty"`
	ReceiveInvoices            bool     `json:"receiveInvoices,omitempty"`
	ReminderGroupID            string   `json:"reminderGroupId,omitempty"`
	ReceiveReminders           bool     `json:"receiveReminders,omitempty"`
	ReceiveStatements          bool     `json:"receiveStatements,omitempty"`
	TaxBandID                  string   `json:"taxBandId,omitempty"`
	TaxBandFixed               bool     `json:"taxBandFixed,omitempty"`
	TaxAuthorityID             string   `json:"taxAuthorityId,omitempty"`
	IsMultiTaxAuthority        bool     `json:"isMultiTaxAuthority,omitempty"`
	MultiTaxAuthorityIds       []string `json:"multiTaxAuthorityIds,omitempty"`
	StockLocationID            string   `json:"stockLocationId,omitempty"`
	IsActive                   bool     `json:"isActive,omitempty"`
	CreditReportDate           DateTime `json:"creditReportDate,omitempty"`
	CreditScore                int      `json:"creditScore,omitempty"`
	RiskLevel                  string   `json:"riskLevel,omitempty"`
	IgnoreDeferredIncome       bool     `json:"ignoreDeferredIncome,omitempty"`
	IgnoreDeposit              bool     `json:"ignoreDeposit,omitempty"`
	IgnoreSchedule             bool     `json:"ignoreSchedule,omitempty"`
	IsHold                     bool     `json:"isHold,omitempty"`
	IsStop                     bool     `json:"isStop,omitempty"`
	ProjectID                  string   `json:"projectId,omitempty"`
	ApplyDomesticReverseCharge bool     `json:"applyDomesticReverseCharge,omitempty"`
	OrderRequired              bool     `json:"orderRequired,omitempty"`
	CostCentreID               string   `json:"costCentreId,omitempty"`
	CostCentreFixed            bool     `json:"costCentreFixed,omitempty"`
	DepartmentID               string   `json:"departmentId,omitempty"`
	DepartmentFixed            bool     `json:"departmentFixed,omitempty"`
	WarningMessage             string   `json:"warningMessage,omitempty"`
	Ext                        ExtPost  `json:"ext,omitzero"`
}

type SupplierExtPost struct {
	AdditionalProp1 string `json:"additionalProp1,omitempty"`
	AdditionalProp2 string `json:"additionalProp2,omitempty"`
	AdditionalProp3 string `json:"additionalProp3,omitempty"`
}

type SupplierPost struct {
	ContactGroupSupplierID           string          `json:"contactGroupSupplierId,omitempty"`
	Currency                         string          `json:"currency,omitempty"`
	CreditLimit                      int             `json:"creditLimit,omitempty"`
	DiscountRate                     int             `json:"discountRate,omitempty"`
	PaymentMethodID                  string          `json:"paymentMethodId,omitempty"`
	PaymentRecipientContactAccountID string          `json:"paymentRecipientContactAccountId,omitempty"`
	PayTermID                        string          `json:"payTermId,omitempty"`
	ProductPriceBandID               string          `json:"productPriceBandId,omitempty"`
	TaxBandID                        string          `json:"taxBandId,omitempty"`
	TaxBandFixed                     bool            `json:"taxBandFixed,omitempty"`
	TaxAuthorityID                   string          `json:"taxAuthorityId,omitempty"`
	IsMultiTaxAuthority              bool            `json:"isMultiTaxAuthority,omitempty"`
	SupplierMultiTaxAuthorityIds     []string        `json:"multiTaxAuthorityIds,omitempty"`
	AllowInvoiceDuplication          bool            `json:"allowInvoiceDuplication,omitempty"`
	AllowSelfBilling                 bool            `json:"allowSelfBilling,omitempty"`
	SendRemittance                   bool            `json:"sendRemittance,omitempty"`
	IsActive                         bool            `json:"isActive,omitempty"`
	ProjectID                        string          `json:"projectId,omitempty"`
	ApplyDomesticReverseCharge       bool            `json:"applyDomesticReverseCharge,omitempty"`
	CostCentreID                     string          `json:"costCentreId,omitempty"`
	CostCentreFixed                  bool            `json:"costCentreFixed,omitempty"`
	DepartmentID                     string          `json:"departmentId,omitempty"`
	DepartmentFixed                  bool            `json:"departmentFixed,omitempty"`
	ProductID                        string          `json:"productId,omitempty"`
	AccountID                        string          `json:"accountId,omitempty"`
	LegalEntityID                    string          `json:"legalEntityId,omitempty"`
	BankAccountID                    string          `json:"bankAccountId,omitempty"`
	ResponsibleResourceID            string          `json:"responsibleResourceId,omitempty"`
	StockLocationID                  string          `json:"stockLocationId,omitempty"`
	IgnorePrepayment                 bool            `json:"ignorePrepayment,omitempty"`
	IsHold                           bool            `json:"isHold,omitempty"`
	IsOcrGross                       bool            `json:"isOcrGross,omitempty"`
	IsResale                         bool            `json:"isResale,omitempty"`
	IsStop                           bool            `json:"isStop,omitempty"`
	OcrAliases                       string          `json:"ocrAliases,omitempty"`
	OcrCombineLineItems              bool            `json:"ocrCombineLineItems,omitempty"`
	OcrDisableLineItems              bool            `json:"ocrDisableLineItems,omitempty"`
	WarningMessage                   string          `json:"warningMessage,omitempty"`
	SupplierExt                      SupplierExtPost `json:"ext,omitzero"`
}

type ResourceEmploymentsPost struct {
	DateFrom                      DateTime `json:"dateFrom,omitempty"`
	DateTo                        DateTime `json:"dateTo,omitempty"`
	LegalEntityID                 string   `json:"legalEntityId,omitempty"`
	EmploymentTypeID              string   `json:"employmentTypeId,omitempty"`
	JobTitle                      string   `json:"jobTitle,omitempty"`
	JobGrade                      string   `json:"jobGrade,omitempty"`
	PaymentFrequencyID            string   `json:"paymentFrequencyId,omitempty"`
	PreviousEmployer              string   `json:"previousEmployer,omitempty"`
	TerminationNoticeDays         int      `json:"terminationNoticeDays,omitempty"`
	ExcludeFromIntermediaryReport bool     `json:"excludeFromIntermediaryReport,omitempty"`
	ExternalReference             string   `json:"externalReference,omitempty"`
}

type ResourceUmbrellasPost struct {
	UmbrellaContactAccountID string   `json:"umbrellaContactAccountId,omitempty"`
	DateFrom                 DateTime `json:"dateFrom,omitempty"`
	DateTo                   DateTime `json:"dateTo,omitempty"`
}

type ResourceExtPost struct {
	AdditionalProp1 string `json:"additionalProp1,omitempty"`
	AdditionalProp2 string `json:"additionalProp2,omitempty"`
	AdditionalProp3 string `json:"additionalProp3,omitempty"`
}

type ResourcePost struct {
	IsActive                bool                      `json:"isActive,omitempty"`
	ResourceGroupID         *string                   `json:"resourceGroupId,omitempty"`
	DepartmentID            string                    `json:"departmentId,omitempty"`
	LocationID              string                    `json:"locationId,omitempty"`
	NationalInsuranceNo     string                    `json:"nationalInsuranceNo,omitempty"`
	Currency                string                    `json:"currency,omitempty"`
	AbsenceGroupID          string                    `json:"absenceGroupId,omitempty"`
	CostCentreID            string                    `json:"costCentreId,omitempty"`
	ManagerResourceID       string                    `json:"managerResourceId,omitempty"`
	UniqueTaxpayerReference string                    `json:"uniqueTaxpayerReference,omitempty"`
	IsSalesperson           bool                      `json:"isSalesperson,omitempty"`
	IsBuyer                 bool                      `json:"isBuyer,omitempty"`
	IsRequestor             bool                      `json:"isRequestor,omitempty"`
	AllowExpense            bool                      `json:"allowExpense,omitempty"`
	AllowTask               bool                      `json:"allowTask,omitempty"`
	AllowTimesheet          bool                      `json:"allowTimesheet,omitempty"`
	HasEmployment           bool                      `json:"hasEmployment,omitempty"`
	HasUmbrella             bool                      `json:"hasUmbrella,omitempty"`
	PaymentMethodID         string                    `json:"paymentMethodId,omitempty"`
	PaymentTermsID          string                    `json:"paymentTermsId,omitempty"`
	SendRemittance          bool                      `json:"sendRemittance,omitempty"`
	ResourceEmployments     []ResourceEmploymentsPost `json:"resourceEmployments,omitempty"`
	ResourceUmbrellas       []ResourceUmbrellasPost   `json:"resourceUmbrellas,omitempty"`
	ResourceExt             ResourceExtPost           `json:"ext,omitzero"`
}

type EmailsPost struct {
	Type  string `json:"type,omitempty"`
	Email string `json:"email,omitempty"`
}

type PhonesPost struct {
	Type  string `json:"type,omitempty"`
	Phone string `json:"phone,omitempty"`
}

type AddressesPost struct {
	Type        string `json:"type,omitempty"`
	Address     string `json:"address,omitempty"`
	PostCode    string `json:"postCode,omitempty"`
	City        string `json:"city,omitempty"`
	County      string `json:"county,omitempty"`
	CountryCode string `json:"countryCode,omitempty"`
	Description string `json:"description,omitempty"`
}

type ContactPost struct {
	IntRef      string          `json:"intRef,omitempty"`
	CompanyName string          `json:"companyName,omitempty"`
	Title       string          `json:"title,omitempty"`
	FirstName   string          `json:"firstName,omitempty"`
	MiddleName  string          `json:"middleName,omitempty"`
	LastName    string          `json:"lastName,omitempty"`
	JobTitle    string          `json:"jobTitle,omitempty"`
	LegacyRef   string          `json:"legacyRef,omitempty"`
	Emails      []EmailsPost    `json:"emails,omitempty"`
	Phones      []PhonesPost    `json:"phones,omitempty"`
	Addresses   []AddressesPost `json:"addresses,omitempty"`
}

type ContactsEmailsPost struct {
	Type  string `json:"type,omitempty"`
	Email string `json:"email,omitempty"`
}

type ContactsPhonesPost struct {
	Type  string `json:"type,omitempty"`
	Phone string `json:"phone,omitempty"`
}

type ContactsAddressesPost struct {
	Type        string `json:"type,omitempty"`
	Address     string `json:"address,omitempty"`
	PostCode    string `json:"postCode,omitempty"`
	City        string `json:"city,omitempty"`
	County      string `json:"county,omitempty"`
	CountryCode string `json:"countryCode,omitempty"`
	Description string `json:"description,omitempty"`
}

type ContactsPost struct {
	IntRef            string                  `json:"intRef,omitempty"`
	CompanyName       string                  `json:"companyName,omitempty"`
	Title             string                  `json:"title,omitempty"`
	FirstName         string                  `json:"firstName,omitempty"`
	MiddleName        string                  `json:"middleName,omitempty"`
	LastName          string                  `json:"lastName,omitempty"`
	JobTitle          string                  `json:"jobTitle,omitempty"`
	LegacyRef         string                  `json:"legacyRef,omitempty"`
	ContactsEmails    []ContactsEmailsPost    `json:"emails,omitempty"`
	ContactsPhones    []ContactsPhonesPost    `json:"phones,omitempty"`
	ContactsAddresses []ContactsAddressesPost `json:"addresses,omitempty"`
}

type BankDetailsPost struct {
	Iban        string `json:"iban,omitempty"`
	AccountNo   string `json:"accountNo,omitempty"`
	SortCode    string `json:"sortCode,omitempty"`
	Bic         string `json:"bic,omitempty"`
	RoutingCode string `json:"routingCode,omitempty"`
	Currency    string `json:"currency,omitempty"`
	LegacyRef   string `json:"legacyRef,omitempty"`
	BankID      string `json:"bankId,omitempty"`
	BankRef     string `json:"bankRef,omitempty"`
	CountryCode string `json:"countryCode,omitempty"`
	BankAddress string `json:"bankAddress,omitempty"`
	AccountName string `json:"accountName,omitempty"`
}

// SaleInvoicePost

type SaleInvoicePost struct {
	DocNo                 string                     `json:"docNo,omitempty"`
	TransNo               int                        `json:"transNo,omitempty"`
	DocSerieID            string                     `json:"docSerieId,omitempty"`
	DocTypeID             string                     `json:"docTypeId,omitempty"`
	LegalEntityID         string                     `json:"legalEntityId,omitempty"`
	CdacID                string                     `json:"cdacId,omitempty"`
	ContactAccountID      string                     `json:"contactAccountId,omitempty"`
	PeriodID              string                     `json:"periodId,omitempty"`
	AccountID             string                     `json:"accountId,omitempty"`
	Post                  PostSaleRequest            `json:"post,omitzero"`
	TaxAuthorityID        string                     `json:"taxAuthorityId,omitempty"`
	PaymentMethodID       string                     `json:"paymentMethodId,omitempty"`
	ResponsibleResourceID string                     `json:"responsibleResourceId,omitempty"`
	ProjectID             string                     `json:"projectId,omitempty"`
	DocDate               DateTime                   `json:"docDate,omitempty"`
	Description           string                     `json:"description,omitempty"`
	TaxDate               DateTime                   `json:"taxDate,omitempty"`
	DueDate               DateTime                   `json:"dueDate,omitempty"`
	Currency              string                     `json:"currency,omitempty"`
	CurrencyRate          int                        `json:"currencyRate,omitempty"`
	IsNetEntry            bool                       `json:"isNetEntry,omitempty"`
	ExplicitAmounts       bool                       `json:"explicitAmounts,omitempty"`
	TheirDocNo            string                     `json:"theirDocNo,omitempty"`
	TheirRef              string                     `json:"theirRef,omitempty"`
	IntRef                string                     `json:"intRef,omitempty"`
	LegacyRef             string                     `json:"legacyRef,omitempty"`
	Details               []DetailsSaleRequest       `json:"details,omitempty"`
	BankAccountID         string                     `json:"bankAccountId,omitempty"`
	BankCurrencyRate      int                        `json:"bankCurrencyRate,omitempty"`
	BankCurrencyAmount    int                        `json:"bankCurrencyAmount,omitempty"`
	BankRef               string                     `json:"bankRef,omitempty"`
	CurrencyAmount        float64                    `json:"currencyAmount,omitempty"`
	AllocateAuto          bool                       `json:"allocateAuto,omitempty"`
	Allocations           []AllocationsSaleRequest   `json:"allocations,omitempty"`
	DeliveryAddress       DeliveryAddressSaleRequest `json:"deliveryAddress,omitzero"`
	BillingAddress        BillingAddressSaleRequest  `json:"billingAddress,omitzero"`
	NetCurrencyAmount     int                        `json:"netCurrencyAmount,omitempty"`
	TaxCurrencyAmount     int                        `json:"taxCurrencyAmount,omitempty"`
	GrossCurrencyAmount   int                        `json:"grossCurrencyAmount,omitempty"`
	NetAmount             int                        `json:"netAmount,omitempty"`
	TaxAmount             int                        `json:"taxAmount,omitempty"`
	GrossAmount           int                        `json:"grossAmount,omitempty"`
	FromStockLocationID   string                     `json:"fromStockLocationId,omitempty"`
	ToStockLocationID     string                     `json:"toStockLocationId,omitempty"`
	StockDate             DateTime                   `json:"stockDate,omitempty"`
	PaymentTermsID        string                     `json:"paymentTermsId,omitempty"`
	Ext                   ExtSaleRequest             `json:"ext,omitzero"`
	TextHeader            string                     `json:"textHeader,omitempty"`
	TextFooter            string                     `json:"textFooter,omitempty"`
	OrderNo               string                     `json:"orderNo,omitempty"`
	DeliveryDate          DateTime                   `json:"deliveryDate,omitempty"`
	AdditionalProp1       string                     `json:"additionalProp1,omitempty"`
	AdditionalProp2       string                     `json:"additionalProp2,omitempty"`
	AdditionalProp3       string                     `json:"additionalProp3,omitempty"`
}
type PostSaleRequest struct {
	AdditionalProp1 string `json:"additionalProp1,omitempty"`
	AdditionalProp2 string `json:"additionalProp2,omitempty"`
	AdditionalProp3 string `json:"additionalProp3,omitempty"`
}
type DetailsPostSaleRequest struct {
	CostCentre string `json:"CostCentre,omitempty"`
	Department string `json:"Department,omitempty"`
	RateCode   string `json:"RateCode,omitempty"`
}
type DeferredDistributionSaleRequest struct {
	ProfileID string `json:"profileId,omitempty"`
	DateFrom  string `json:"dateFrom,omitempty"`
	DateTo    string `json:"dateTo,omitempty"`
	Periods   int    `json:"periods,omitempty"`
}
type DetailsSaleRequest struct {
	ProductID              string                          `json:"productId,omitempty"`
	ProductSkuID           string                          `json:"productSkuId,omitempty"`
	ProjectID              string                          `json:"projectId,omitempty"`
	TaxCodeID              string                          `json:"taxCodeId,omitempty"`
	UomID                  string                          `json:"uomId,omitempty"`
	AccountID              string                          `json:"accountId,omitempty"`
	DetailsPost            DetailsPostSaleRequest          `json:"post,omitzero"`
	TaxBandID              string                          `json:"taxBandId,omitempty"`
	LineDate               DateTime                        `json:"lineDate,omitempty"`
	Description            string                          `json:"description,omitempty"`
	IsNetEntry             bool                            `json:"isNetEntry,omitempty"`
	CurrencyUnitPrice      float64                         `json:"currencyUnitPrice,omitempty"`
	NetCurrencyUnitPrice   float64                         `json:"netCurrencyUnitPrice,omitempty"`
	GrossCurrencyUnitPrice float64                         `json:"grossCurrencyUnitPrice,omitempty"`
	Quantity               float64                         `json:"quantity,omitempty"`
	SkuQuantity            float64                         `json:"skuQuantity,omitempty"`
	NetCurrencyAmount      float64                         `json:"netCurrencyAmount,omitempty"`
	TaxCurrencyAmount      float64                         `json:"taxCurrencyAmount,omitempty"`
	GrossCurrencyAmount    float64                         `json:"grossCurrencyAmount,omitempty"`
	NetAmount              float64                         `json:"netAmount,omitempty"`
	TaxAmount              float64                         `json:"taxAmount,omitempty"`
	GrossAmount            float64                         `json:"grossAmount,omitempty"`
	FromStockLocationID    string                          `json:"fromStockLocationId,omitempty"`
	ToStockLocationID      string                          `json:"toStockLocationId,omitempty"`
	StockDate              DateTime                        `json:"stockDate,omitempty"`
	IsResale               bool                            `json:"isResale,omitempty"`
	TextPre                string                          `json:"textPre,omitempty"`
	TextPost               string                          `json:"textPost,omitempty"`
	IntRef                 string                          `json:"intRef,omitempty"`
	LegacyRef              string                          `json:"legacyRef,omitempty"`
	ProductType            string                          `json:"productType,omitempty"`
	IsDeferred             bool                            `json:"isDeferred,omitempty"`
	DeferredDistribution   DeferredDistributionSaleRequest `json:"deferredDistribution,omitzero"`
	AdditionalProp1        string                          `json:"additionalProp1,omitempty"`
	AdditionalProp2        string                          `json:"additionalProp2,omitempty"`
	AdditionalProp3        string                          `json:"additionalProp3,omitempty"`
}
type AllocationsSaleRequest struct {
	DocID               string `json:"docId,omitempty"`
	AllocAmountCurrency int    `json:"allocAmountCurrency,omitempty"`
	AllocAmountBase     int    `json:"allocAmountBase,omitempty"`
}
type DeliveryAddressSaleRequest struct {
	Type        string `json:"type,omitempty"`
	Address     string `json:"address,omitempty"`
	PostCode    string `json:"postCode,omitempty"`
	City        string `json:"city,omitempty"`
	County      string `json:"county,omitempty"`
	CountryCode string `json:"countryCode,omitempty"`
	Description string `json:"description,omitempty"`
}
type BillingAddressSaleRequest struct {
	Type        string `json:"type,omitempty"`
	Address     string `json:"address,omitempty"`
	PostCode    string `json:"postCode,omitempty"`
	City        string `json:"city,omitempty"`
	County      string `json:"county,omitempty"`
	CountryCode string `json:"countryCode,omitempty"`
	Description string `json:"description,omitempty"`
}
type ExtSaleRequest struct {
	AdditionalProp1 string `json:"additionalProp1,omitempty"`
	AdditionalProp2 string `json:"additionalProp2,omitempty"`
	AdditionalProp3 string `json:"additionalProp3,omitempty"`
}

// ReceiptPost

type ReceiptPost struct {
	DocNo                 string                    `json:"docNo,omitempty"`
	TransNo               int                       `json:"transNo,omitempty"`
	DocSerieID            string                    `json:"docSerieId,omitempty"`
	DocTypeID             string                    `json:"docTypeId,omitempty"`
	LegalEntityID         string                    `json:"legalEntityId,omitempty"`
	CdacID                string                    `json:"cdacId,omitempty"`
	ContactAccountID      string                    `json:"contactAccountId,omitempty"`
	PeriodID              string                    `json:"periodId,omitempty"`
	AccountID             string                    `json:"accountId,omitempty"`
	Post                  PostReceiptReq            `json:"post,omitzero"`
	TaxAuthorityID        string                    `json:"taxAuthorityId,omitempty"`
	PaymentMethodID       string                    `json:"paymentMethodId,omitempty"`
	ResponsibleResourceID string                    `json:"responsibleResourceId,omitempty"`
	ProjectID             string                    `json:"projectId,omitempty"`
	DocDate               DateTime                  `json:"docDate,omitempty"`
	Description           string                    `json:"description,omitempty"`
	TaxDate               DateTime                  `json:"taxDate,omitempty"`
	DueDate               DateTime                  `json:"dueDate,omitempty"`
	Currency              string                    `json:"currency,omitempty"`
	CurrencyRate          int                       `json:"currencyRate,omitempty"`
	IsNetEntry            bool                      `json:"isNetEntry,omitempty"`
	ExplicitAmounts       bool                      `json:"explicitAmounts,omitempty"`
	TheirDocNo            string                    `json:"theirDocNo,omitempty"`
	TheirRef              string                    `json:"theirRef,omitempty"`
	IntRef                string                    `json:"intRef,omitempty"`
	LegacyRef             string                    `json:"legacyRef,omitempty"`
	Details               []DetailsReceiptReq       `json:"details,omitempty"`
	BankAccountID         string                    `json:"bankAccountId,omitempty"`
	BankCurrencyRate      int                       `json:"bankCurrencyRate,omitempty"`
	BankCurrencyAmount    float64                   `json:"bankCurrencyAmount,omitempty"`
	BankCurrency          string                    `json:"bankCurrency,omitempty"`
	BankRef               string                    `json:"bankRef,omitempty"`
	CurrencyAmount        float64                   `json:"currencyAmount,omitempty"`
	AllocateAuto          bool                      `json:"allocateAuto,omitempty"`
	Allocations           []AllocationsReceiptReq   `json:"allocations,omitempty"`
	DeliveryAddress       DeliveryAddressReceiptReq `json:"deliveryAddress,omitzero"`
	BillingAddress        BillingAddressReceiptReq  `json:"billingAddress,omitzero"`
	NetCurrencyAmount     int                       `json:"netCurrencyAmount,omitempty"`
	TaxCurrencyAmount     int                       `json:"taxCurrencyAmount,omitempty"`
	GrossCurrencyAmount   int                       `json:"grossCurrencyAmount,omitempty"`
	NetAmount             int                       `json:"netAmount,omitempty"`
	TaxAmount             int                       `json:"taxAmount,omitempty"`
	GrossAmount           int                       `json:"grossAmount,omitempty"`
	FromStockLocationID   string                    `json:"fromStockLocationId,omitempty"`
	ToStockLocationID     string                    `json:"toStockLocationId,omitempty"`
	StockDate             DateTime                  `json:"stockDate,omitempty"`
	PaymentTermsID        string                    `json:"paymentTermsId,omitempty"`
	Ext                   ExtReceiptReq             `json:"ext,omitzero"`
	TextHeader            string                    `json:"textHeader,omitempty"`
	TextFooter            string                    `json:"textFooter,omitempty"`
	OrderNo               string                    `json:"orderNo,omitempty"`
	DeliveryDate          DateTime                  `json:"deliveryDate,omitempty"`
	AdditionalProp1       string                    `json:"additionalProp1,omitempty"`
	AdditionalProp2       string                    `json:"additionalProp2,omitempty"`
	AdditionalProp3       string                    `json:"additionalProp3,omitempty"`
}
type PostReceiptReq struct {
	CostCentre string `json:"CostCentre,omitempty"`
	Department string `json:"Department,omitempty"`
	RateCode   string `json:"RateCode,omitempty"`
}
type DetailsPostReceiptReq struct {
	AdditionalProp1 string `json:"additionalProp1,omitempty"`
	AdditionalProp2 string `json:"additionalProp2,omitempty"`
	AdditionalProp3 string `json:"additionalProp3,omitempty"`
}
type DeferredDistributionReceiptReq struct {
	ProfileID string `json:"profileId,omitempty"`
	DateFrom  string `json:"dateFrom,omitempty"`
	DateTo    string `json:"dateTo,omitempty"`
	Periods   int    `json:"periods,omitempty"`
}
type DetailsReceiptReq struct {
	ProductID              string                         `json:"productId,omitempty"`
	ProductSkuID           string                         `json:"productSkuId,omitempty"`
	ProjectID              string                         `json:"projectId,omitempty"`
	TaxCodeID              string                         `json:"taxCodeId,omitempty"`
	UomID                  string                         `json:"uomId,omitempty"`
	AccountID              string                         `json:"accountId,omitempty"`
	DetailsPost            DetailsPostReceiptReq          `json:"post,omitzero"`
	TaxBandID              string                         `json:"taxBandId,omitempty"`
	LineDate               DateTime                       `json:"lineDate,omitempty"`
	Description            string                         `json:"description,omitempty"`
	IsNetEntry             bool                           `json:"isNetEntry,omitempty"`
	CurrencyUnitPrice      int                            `json:"currencyUnitPrice,omitempty"`
	NetCurrencyUnitPrice   int                            `json:"netCurrencyUnitPrice,omitempty"`
	GrossCurrencyUnitPrice int                            `json:"grossCurrencyUnitPrice,omitempty"`
	Quantity               int                            `json:"quantity,omitempty"`
	SkuQuantity            int                            `json:"skuQuantity,omitempty"`
	NetCurrencyAmount      int                            `json:"netCurrencyAmount,omitempty"`
	TaxCurrencyAmount      int                            `json:"taxCurrencyAmount,omitempty"`
	GrossCurrencyAmount    int                            `json:"grossCurrencyAmount,omitempty"`
	NetAmount              int                            `json:"netAmount,omitempty"`
	TaxAmount              int                            `json:"taxAmount,omitempty"`
	GrossAmount            int                            `json:"grossAmount,omitempty"`
	FromStockLocationID    string                         `json:"fromStockLocationId,omitempty"`
	ToStockLocationID      string                         `json:"toStockLocationId,omitempty"`
	StockDate              DateTime                       `json:"stockDate,omitempty"`
	IsResale               bool                           `json:"isResale,omitempty"`
	TextPre                string                         `json:"textPre,omitempty"`
	TextPost               string                         `json:"textPost,omitempty"`
	IntRef                 string                         `json:"intRef,omitempty"`
	LegacyRef              string                         `json:"legacyRef,omitempty"`
	ProductType            string                         `json:"productType,omitempty"`
	IsDeferred             bool                           `json:"isDeferred,omitempty"`
	DeferredDistribution   DeferredDistributionReceiptReq `json:"deferredDistribution,omitzero"`
	AdditionalProp1        string                         `json:"additionalProp1,omitempty"`
	AdditionalProp2        string                         `json:"additionalProp2,omitempty"`
	AdditionalProp3        string                         `json:"additionalProp3,omitempty"`
}
type AllocationsReceiptReq struct {
	DocID               string `json:"docId,omitempty"`
	AllocAmountCurrency int    `json:"allocAmountCurrency,omitempty"`
	AllocAmountBase     int    `json:"allocAmountBase,omitempty"`
}
type DeliveryAddressReceiptReq struct {
	Type        string `json:"type,omitempty"`
	Address     string `json:"address,omitempty"`
	PostCode    string `json:"postCode,omitempty"`
	City        string `json:"city,omitempty"`
	County      string `json:"county,omitempty"`
	CountryCode string `json:"countryCode,omitempty"`
	Description string `json:"description,omitempty"`
}
type BillingAddressReceiptReq struct {
	Type        string `json:"type,omitempty"`
	Address     string `json:"address,omitempty"`
	PostCode    string `json:"postCode,omitempty"`
	City        string `json:"city,omitempty"`
	County      string `json:"county,omitempty"`
	CountryCode string `json:"countryCode,omitempty"`
	Description string `json:"description,omitempty"`
}
type ExtReceiptReq struct {
	AdditionalProp1 string `json:"additionalProp1,omitempty"`
	AdditionalProp2 string `json:"additionalProp2,omitempty"`
	AdditionalProp3 string `json:"additionalProp3,omitempty"`
}
