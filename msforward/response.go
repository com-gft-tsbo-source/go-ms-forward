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
func InitForwardResponse(r *ForwardResponse, status string, ms *MsForward) {
	microservice.InitResponseFromMicroService(&r.Response, ms, status)
	r.Value = 0
}

// NewForwardResponse ...
func NewForwardResponse(status string, ms *MsForward) *ForwardResponse {
	var r ForwardResponse
	InitForwardResponse(&r, status, ms)
	return &r
}
