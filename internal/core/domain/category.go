package domain

import "fmt"

const (
	CategoryExpenseType = "expense"
	CategoryIncomeType  = "income"
)

const (
	ExpenseRent                = "rent"
	ExpenseWaterBill           = "water_bill"
	ExpenseElectricityBill     = "electricity_bill"
	ExpenseGroceries           = "groceries"
	ExpenseRestaurantsDelivery = "restaurants_delivery"
	ExpenseTransportation      = "transportation"
	ExpenseHealth              = "health"
	ExpenseEducation           = "education"
	ExpenseEntertainment       = "entertainment"
	ExpenseClothing            = "clothing"
	ExpenseBeauty              = "beauty"
	ExpenseGifts               = "gifts"
	ExpenseVehicle             = "vehicle"
	ExpenseHome                = "home"
	ExpenseInsurance           = "insurance"
	ExpenseTaxes               = "taxes"
	ExpenseInvestments         = "investments"
	ExpenseSavings             = "savings"
	ExpenseDebts               = "debts"
	ExpenseLoans               = "loans"
	ExpenseCreditCard          = "credit_card"
	ExpenseOthersExpense       = "others_expense"
)

const (
	IncomeSalary       = "salary"
	IncomeInvestments  = "investments"
	IncomeFreelancer   = "freelancer"
	IncomeBonus        = "bonus"
	IncomeOthersIncome = "others_income"
)

type Category struct {
	Base
	Type string `json:"type"`
	Name string `json:"name"`
}

func BuildNewCategory(categoryType, name string) (*Category, error) {
	c := &Category{
		Type: categoryType,
		Name: name,
	}

	if err := c.validate(); err != nil {
		return nil, err
	}

	return c, nil
}

func (c *Category) validate() error {
	validCategoryTypes := map[string]bool{
		CategoryExpenseType: true,
		CategoryIncomeType:  true,
	}

	if !validCategoryTypes[c.Type] {
		return fmt.Errorf("invalid category type")
	}

	validExpenseCategories := map[string]bool{
		ExpenseRent:                true,
		ExpenseWaterBill:           true,
		ExpenseElectricityBill:     true,
		ExpenseGroceries:           true,
		ExpenseRestaurantsDelivery: true,
		ExpenseTransportation:      true,
		ExpenseHealth:              true,
		ExpenseEducation:           true,
		ExpenseEntertainment:       true,
		ExpenseClothing:            true,
		ExpenseBeauty:              true,
		ExpenseGifts:               true,
		ExpenseVehicle:             true,
		ExpenseHome:                true,
		ExpenseInsurance:           true,
		ExpenseTaxes:               true,
		ExpenseInvestments:         true,
		ExpenseSavings:             true,
		ExpenseDebts:               true,
		ExpenseLoans:               true,
		ExpenseCreditCard:          true,
		ExpenseOthersExpense:       true,
	}

	validIncomeCategories := map[string]bool{
		IncomeSalary:       true,
		IncomeInvestments:  true,
		IncomeFreelancer:   true,
		IncomeBonus:        true,
		IncomeOthersIncome: true,
	}

	if c.Type == CategoryExpenseType && !validExpenseCategories[c.Name] {
		return fmt.Errorf("invalid expense category")
	}

	if c.Type == CategoryIncomeType && !validIncomeCategories[c.Name] {
		return fmt.Errorf("invalid income category")
	}

	return nil
}
