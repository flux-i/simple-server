package routes

import (
	"github.com/gorilla/mux"

	handler "simple-server/src/handlers"
	manager "simple-server/src/managers"
)

func AddRoutes(rm *manager.RequestManager, router *mux.Router) {
	router.HandleFunc("/api/company/financials", handler.GetCompanyFinancials(rm)).Methods("GET").Name(
		"get_company_financials")
	router.HandleFunc("/api/sales/data", handler.GetCompanySales(rm)).Methods("GET").Name(
		"get_company_sales")
	router.HandleFunc("/api/employee/stats", handler.GetCompanyEmployees(rm)).Methods("GET").Name(
		"get_company_employees")
}
