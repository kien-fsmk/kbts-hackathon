package payment

type Payment struct {
	PaymentID   string    `json:"payment_id"`
	Amount      int       `json:"amount"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   string    `json:"created_at"`
	Customer    Customer  `json:"customer"`
	Recipient   Recipient `json:"recipient"`
	Category    string    `json:"category"`
}

type PaymentCategoryPercentage struct {
	CategoryName string
	Percentage   float64
}

type Customer struct {
	CustomerID string `json:"customer_id"`
	Email      string `json:"email"`
}

type Recipient struct {
	RecipientID             string      `json:"recipient_id"`
	Name                    string      `json:"name"`
	Email                   string      `json:"email"`
	BusinessRegistrationNum string      `json:"business_registration_number"`
	BankAccount             BankAccount `json:"bank_account"`
}

type BankAccount struct {
	BankName    string `json:"bank_name"`
	AccountNo   string `json:"account_no"`
	AccountName string `json:"account_name"`
	Country     string `json:"country"`
}

type PaymentCategory int

const (
	PaymentCategoryAdvertising PaymentCategory = iota + 1
	PaymentCategorySubscription
	PaymentCategoryEmployeeBenefitProgram
	PaymentCategoryInsurance
	PaymentCategoryLegalAndProfessionalExpenses
	PaymentCategoryOfficeExpensesAndSupplies
	PaymentCategoryTelecommunication
	PaymentCategoryUtilities
	PaymentCategoryRentalOrPayment
	PaymentCategoryPostageAndShipping
	PaymentCategoryMedicalExpenses
	PaymentCategoryEntertainment
	PaymentCategoryLicensesAndPermits
	PaymentCategoryWages
	PaymentCategoryOther
	PaymentCategoryInvestment
)

var CategoryStr = map[PaymentCategory]string{
	PaymentCategoryAdvertising:                  "Advertising",
	PaymentCategorySubscription:                 "Subscription",
	PaymentCategoryEmployeeBenefitProgram:       "Employee Benefit Program",
	PaymentCategoryInsurance:                    "Insurance",
	PaymentCategoryLegalAndProfessionalExpenses: "Legal and Professional Expenses",
	PaymentCategoryOfficeExpensesAndSupplies:    "Office Expenses and Supplies",
	PaymentCategoryTelecommunication:            "Telecommunication",
	PaymentCategoryUtilities:                    "Utilities",
	PaymentCategoryRentalOrPayment:              "Rental or Payment",
	PaymentCategoryPostageAndShipping:           "Postage and Shipping",
	PaymentCategoryMedicalExpenses:              "Medical Expenses",
	PaymentCategoryEntertainment:                "Entertainment",
	PaymentCategoryLicensesAndPermits:           "Licenses and Permits",
	PaymentCategoryWages:                        "Wages",
	PaymentCategoryOther:                        "Other",
	PaymentCategoryInvestment:                   "Investment",
}
