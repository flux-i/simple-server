package handlers

import (
	"log"
	"net/http"

	manager "simple-server/src/managers"
	request "simple-server/src/requests"
	usecase "simple-server/src/usecases"
)

func GetCompanyEmployees(rm *manager.RequestManager) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		log.Println("API: GetCompanyEmployees")

		appReq := request.NewPathIDAppRequest(res, req)

		uc := usecase.NewGetCompanyEmployeesUseCase(rm, appReq)
		uc.GetCompanyEmployees()
		if appReq.HasErrors() {
			appReq.WriteErrorResponse()
			return
		}
		appReq.WriteSuccessResponse()
	}
}
