package usecases

import (
	"log"

	et "simple-server/src/entities"
	manager "simple-server/src/managers"
	request "simple-server/src/requests"
)

type GetCompanySalesUseCase struct {
	rm  *manager.RequestManager
	req *request.PathIDAppRequest
}

func NewGetCompanySalesUseCase(
	rm *manager.RequestManager, req *request.PathIDAppRequest) *GetCompanySalesUseCase {
	return &GetCompanySalesUseCase{rm: rm, req: req}
}

func (uc *GetCompanySalesUseCase) GetCompanySales() {
	if !uc.req.Valid() {
		log.Println("Get company sales request is not valid")
		return
	}

	response := make(chan et.ResponseData)
	request := et.Request{
		CompanyID: uc.req.ID,
		API:       "sale",
		Response:  response,
	}
	go uc.rm.ProcessRequest(request)
	data := <-response
	uc.req.AddResponsePayload("sale_data", data)
}
