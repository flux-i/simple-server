package usecases

import (
	"log"

	et "simple-server/src/entities"
	manager "simple-server/src/managers"
	request "simple-server/src/requests"
)

type GetCompanyFinancialsUseCase struct {
	rm  *manager.RequestManager
	req *request.PathIDAppRequest
}

func NewGetCompanyFinancialsUseCase(
	rm *manager.RequestManager, req *request.PathIDAppRequest) *GetCompanyFinancialsUseCase {
	return &GetCompanyFinancialsUseCase{rm: rm, req: req}
}

func (uc *GetCompanyFinancialsUseCase) GetCompanyFinancials() {
	if !uc.req.Valid() {
		log.Println("Get company financials is not valid")
		return
	}

	response := make(chan et.ResponseData)
	request := et.Request{
		CompanyID: uc.req.ID,
		API:       "financial",
		Response:  response,
	}
	go uc.rm.ProcessRequest(request)
	data := <-response
	log.Println(data)
	uc.req.AddResponsePayload("financial_data", data)
}
