package univush

import (
	"github.com/sideshow/apns2"
)

type ClientResponse struct {
	ID   string
	Code int
	Sent bool
}

func NewClientResponseFromAPNS2(resp *apns2.Response) *ClientResponse {
	return &ClientResponse{
		ID:   resp.ApnsID,
		Code: resp.StatusCode,
		Sent: resp.Sent(),
	}
}

func NewClientResponseFromFCM(id string) *ClientResponse {
	return &ClientResponse{
		ID:   id,
		Code: 200,
		Sent: true,
	}
}
