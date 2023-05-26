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

type PaymentCategory string

const (
	PaymentCategoryAdvertising                  PaymentCategory = "Advertising"
	PaymentCategoryEmployeeBenefitProgram       PaymentCategory = "Employee Benefit Program"
	PaymentCategoryInsurance                    PaymentCategory = "Insurance"
	PaymentCategoryLegalAndProfessionalExpenses PaymentCategory = "Legal and professional expenses"
	PaymentCategoryOfficeExpensesAndSupplies    PaymentCategory = "Office expenses and supplies"
	PaymentCategoryTelecommunication            PaymentCategory = "Telecommunication"
	PaymentCategoryUtilities                    PaymentCategory = "Utilities"
	PaymentCategoryRentalOrPayment              PaymentCategory = "Rental / Payment"
	PaymentCategoryPostageAndShipping           PaymentCategory = "Postage and Shipping"
	PaymentCategoryMedicalExpenses              PaymentCategory = "Medical expenses"
	PaymentCategoryEntertainment                PaymentCategory = "Entertainment"
	PaymentCategoryLicensesAndPermits           PaymentCategory = "Licenses and Permits"
	PaymentCategoryWages                        PaymentCategory = "Wages"
	PaymentCategoryOther                        PaymentCategory = "Other"
	PaymentCategoryInvestment                   PaymentCategory = "Investment"
)
