package univush

import (
	"context"
	"errors"

	"firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"github.com/sideshow/apns2"
	"google.golang.org/api/option"
)

// ClientType - the type of the client "fcm" or "apns2"
type ClientType int

const (
	// ClientTypeFCM - FCM client
	ClientTypeFCM ClientType = 1

	// ClientTypeAPNS2 - APNS2 client
	ClientTypeAPNS2 ClientType = 2
)

// Client - a client connection manager
type Client struct {
	fcm        *messaging.Client
	apns2      *apns2.Client
	clientType ClientType
}

// NewFCMClient - constructs a new FCM client
func NewFCMClient(ctx context.Context, appID, appToken string) (*Client, error) {
	app, err := firebase.NewApp(ctx, nil, option.WithAPIKey(appToken))
	if err != nil {
		return nil, err
	}

	client, err := app.Messaging(ctx)
	if err != nil {
		return nil, err
	}
	return &Client{
		fcm:        client,
		clientType: ClientTypeFCM,
	}, nil
}

// NewAPNS2CLient - constructs a new APNS2 client
func NewAPNS2CLient(ctx context.Context, cert []byte, passphrase string) (*Client, error) {
	crt, err := CertBytes(cert, passphrase)
	if err != nil {
		return nil, err
	}
	return &Client{
		apns2:      apns2.NewClient(crt),
		clientType: ClientTypeAPNS2,
	}, nil
}

// Send - push the message
func (c *Client) Send(ctx context.Context, n *Notification) (*ClientResponse, error) {
	switch c.clientType {
	case ClientTypeAPNS2:
		resp, err := c.apns2.PushWithContext(ctx, n.ToAPNS())
		if err != nil {
			return nil, err
		}
		return NewClientResponseFromAPNS2(resp), nil
	case ClientTypeFCM:
		id, err := c.fcm.Send(ctx, n.ToAndroid())
		if err != nil {
			return nil, err
		}
		return NewClientResponseFromFCM(id), nil
	}

	return nil, errors.New("invalid client specified")
}
