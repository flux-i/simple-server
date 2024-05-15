package handlers

import (
	"log"
	"net/http"

	request "simple-server/src/requests"
	usecase "simple-server/src/usecases"
)

func GetCompanySales(res http.ResponseWriter, req *http.Request) {
	log.Println("API: GetCompanySales")

	appReq := request.NewPathIDAppRequest(res, req)

	uc := usecase.NewGetCompanySalesUseCase(appReq)
	uc.GetCompanySales()
	if appReq.HasErrors() {
		appReq.WriteErrorResponse()
		return
	}
	appReq.WriteSuccessResponse()
}
