package usecases

import (
	"log"

	request "simple-server/src/requests"
)

type GetCompanySalesUseCase struct {
	req *request.PathIDAppRequest
}

func NewGetCompanySalesUseCase(req *request.PathIDAppRequest) *GetCompanySalesUseCase {
	return &GetCompanySalesUseCase{req: req}
}

func (uc *GetCompanySalesUseCase) GetCompanySales() {
	if !uc.req.Valid() {
		log.Println("Get company sales request is not valid")
		return
	}

	// Some processing logic here

	// Send response
	// uc.req.AddResponsePayload("company_sales", data)
}
