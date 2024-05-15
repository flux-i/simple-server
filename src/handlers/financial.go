package handlers

import (
	"log"
	"net/http"

	request "simple-server/src/requests"
	usecase "simple-server/src/usecases"
)

func GetCompanyFinancials(res http.ResponseWriter, req *http.Request) {
	log.Println("API: GetCompanyFinancials")

	appReq := request.NewPathIDAppRequest(res, req)

	uc := usecase.NewGetCompanyFinancialsUseCase(appReq)
	uc.GetCompanyFinancials()
	if appReq.HasErrors() {
		appReq.WriteErrorResponse()
		return
	}
	appReq.WriteSuccessResponse()
}
