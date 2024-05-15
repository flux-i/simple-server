package handlers

import (
	"log"
	"net/http"

	request "simple-server/src/requests"
	usecase "simple-server/src/usecases"
)

func GetCompanyEmployees(res http.ResponseWriter, req *http.Request) {
	log.Println("API: GetCompanyEmployees")

	appReq := request.NewPathIDAppRequest(res, req)

	uc := usecase.NewGetCompanyEmployeesUseCase(appReq)
	uc.GetCompanyEmployees()
	if appReq.HasErrors() {
		appReq.WriteErrorResponse()
		return
	}
	appReq.WriteSuccessResponse()
}
