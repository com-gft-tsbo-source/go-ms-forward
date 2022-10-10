package msforward

import (
	"github.com/com-gft-tsbo-source/go-common/ms-framework/microservice"
)

// ###########################################################################
// ###########################################################################
// MsForward Response - Device
// ###########################################################################
// ###########################################################################

// ForwardResponse ...
type ForwardResponse struct {
	microservice.Response
	Value int `json:"value"`
}

// ###########################################################################

// InitForwardResponse Constructor of a response of ms-forward
func InitForwardResponse(r *ForwardResponse, code int, status string, ms *MsForward) {
	microservice.InitResponseFromMicroService(&r.Response, ms, code, status)
	r.Value = 0
}

// NewForwardResponse ...
func NewForwardResponse(code int, status string, ms *MsForward) *ForwardResponse {
	var r ForwardResponse
	InitForwardResponse(&r, code, status, ms)
	return &r
}
