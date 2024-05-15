package usecases

import (
	"log"

	request "simple-server/src/requests"
)

type GetCompanyEmployeesUseCase struct {
	req *request.PathIDAppRequest
}

func NewGetCompanyEmployeesUseCase(req *request.PathIDAppRequest) *GetCompanyEmployeesUseCase {
	return &GetCompanyEmployeesUseCase{req: req}
}

func (uc *GetCompanyEmployeesUseCase) GetCompanyEmployees() {
	if !uc.req.Valid() {
		log.Println("Get company employees request is not valid")
		return
	}

	// Some processing logic here

	// Send response
	// uc.req.AddResponsePayload("company_employees", data)
}
