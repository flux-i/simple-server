package usecases

import (
	"log"

	et "simple-server/src/entities"
	manager "simple-server/src/managers"
	request "simple-server/src/requests"
)

type GetCompanyEmployeesUseCase struct {
	rm  *manager.RequestManager
	req *request.PathIDAppRequest
}

func NewGetCompanyEmployeesUseCase(
	rm *manager.RequestManager, req *request.PathIDAppRequest) *GetCompanyEmployeesUseCase {
	return &GetCompanyEmployeesUseCase{rm: rm, req: req}
}

func (uc *GetCompanyEmployeesUseCase) GetCompanyEmployees() {
	if !uc.req.Valid() {
		log.Println("Get company employees request is not valid")
		return
	}

	response := make(chan et.ResponseData)
	request := et.Request{
		CompanyID: uc.req.ID,
		API:       "employee",
		Response:  response,
	}
	go uc.rm.ProcessRequest(request)
	data := <-response
	uc.req.AddResponsePayload("employee_data", data)
}
