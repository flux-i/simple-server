package routes

import (
	"github.com/gorilla/mux"

	handler "simple-server/src/handlers"
)

func AddRoutes(router *mux.Router) {
	router.HandleFunc("/api/company/financials", handler.GetCompanyFinancials).Methods("GET").Name(
		"get_company_financials")
	router.HandleFunc("/api/sales/data", handler.GetCompanySales).Methods("GET").Name(
		"get_company_sales")
	router.HandleFunc("/api/employee/stats", handler.GetCompanyEmployees).Methods("GET").Name(
		"get_company_employees")
}
