package handlers

import (
	"log"
	"net/http"

	manager "simple-server/src/managers"
	request "simple-server/src/requests"
	usecase "simple-server/src/usecases"
)

func GetCompanyFinancials(rm *manager.RequestManager) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		log.Println("API: GetCompanyFinancials")

		appReq := request.NewPathIDAppRequest(res, req)

		uc := usecase.NewGetCompanyFinancialsUseCase(rm, appReq)
		uc.GetCompanyFinancials()
		if appReq.HasErrors() {
			appReq.WriteErrorResponse()
			return
		}
		appReq.WriteSuccessResponse()
	}
}
