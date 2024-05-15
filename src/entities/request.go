package entities

type Request struct {
	CompanyID int64
	API       string
	Response  chan<- ResponseData
}
