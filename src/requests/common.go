package requests

import (
	"log"
	"net/http"
	"strconv"
)

// PathIDAppRequest handles requests that contain ID parameter in
// request path
type PathIDAppRequest struct {
	Request
	ID int64
}

// NewPathIDAppRequest creates a new PathIDAppRequest
func NewPathIDAppRequest(res http.ResponseWriter,
	req *http.Request) *PathIDAppRequest {
	appReq := new(PathIDAppRequest)
	appReq.Request = Request{request: req, response: res}
	return appReq
}

// Valid extracts the id from path query or sets the errors appropriately.
// We do not do this in mux. In mux we should directly fetch the ID from the path
func (ar *PathIDAppRequest) Valid() bool {
	//vars := mux.Vars(ar.request)
	id := ar.request.URL.Query().Get("companyId")

	if len(id) == 0 {
		ar.Request.AddError(http.StatusBadRequest,
			"path", "companyId", "companyId field is missing or empty")
		return false
	}

	var err error
	ar.ID, err = strconv.ParseInt(id, 10, 64)
	if err != nil {
		ar.Request.AddError(http.StatusBadRequest,
			"path", "companyId", "ID not a valid integer")
		return false
	}
	if ar.ID <= 0 {
		ar.Request.AddError(http.StatusBadRequest,
			"path", "companyId", "ID value should be positive non-zero value")
		return false
	}
	log.Println(ar.ID)
	return true
}
