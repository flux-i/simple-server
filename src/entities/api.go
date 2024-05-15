package entities

// Define structure for company data
// Include fields for financials, sales, employee statistics, etc.

type ResponseData struct {
	InitData interface{}
	ApiData  interface{}
}
type CompanyData struct {
	CompanyID  int64
	OtherValue int
	ThirdValue bool
}

type FinancialData struct {
	Revenue          float64 `json:"revenue"`            // Total revenue of the company
	Expenses         float64 `json:"expenses"`           // Total expenses of the company
	Profit           float64 `json:"profit"`             // Profit (Revenue - Expenses) of the company
	NetIncome        float64 `json:"net_income"`         // Net income of the company
	EarningsPerShare float64 `json:"earnings_per_share"` // Earnings per share of the company
}

// SalesData represents sales information for a company.
type SalesData struct {
	TotalSales     float64 `json:"total_sales"`     // Total sales revenue of the company
	UnitsSold      int     `json:"units_sold"`      // Total number of units sold by the company
	AveragePrice   float64 `json:"average_price"`   // Average price per unit sold
	TotalCustomers int     `json:"total_customers"` // Total number of customers who made purchases
}

// EmployeeStats represents statistics on employees for a company.
type EmployeeStats struct {
	TotalEmployees int `json:"total_employees"` // Total number of employees in the company
	AvgSalary      int `json:"avg_salary"`      // Average salary of employees
}
