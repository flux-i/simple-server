package usecases

import (
	"log"

	request "simple-server/src/requests"
)

type GetCompanyFinancialsUseCase struct {
	req *request.PathIDAppRequest
}

func NewGetCompanyFinancialsUseCase(req *request.PathIDAppRequest) *GetCompanyFinancialsUseCase {
	return &GetCompanyFinancialsUseCase{req: req}
}

func (uc *GetCompanyFinancialsUseCase) GetCompanyFinancials() {
	if !uc.req.Valid() {
		log.Println("Get company financials is not valid")
		return
	}

	// Some processing logic here

	// Send response
	// uc.req.AddResponsePayload("company_financials", data)
}
