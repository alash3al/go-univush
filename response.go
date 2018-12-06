package univush

import (
	"github.com/sideshow/apns2"
)

// ClientResponse - client response
type ClientResponse struct {
	ID   string
	Code int
	Sent bool
}

// NewClientResponseFromAPNS2 - create a APNS2 response from client respose
func NewClientResponseFromAPNS2(resp *apns2.Response) *ClientResponse {
	return &ClientResponse{
		ID:   resp.ApnsID,
		Code: resp.StatusCode,
		Sent: resp.Sent(),
	}
}

// NewClientResponseFromFCM - create a FCM response from client response
func NewClientResponseFromFCM(id string) *ClientResponse {
	return &ClientResponse{
		ID:   id,
		Code: 200,
		Sent: true,
	}
}
