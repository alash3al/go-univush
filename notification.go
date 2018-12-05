package univush

import (
	"strconv"
	"strings"
	"time"

	"firebase.google.com/go/messaging"
	"github.com/sideshow/apns2"
	"github.com/sideshow/apns2/payload"
)

// Notification - a universal notification format
type Notification struct {
	Title       string            `json:"title"`
	Body        string            `json:"body"`
	Icon        string            `json:"icon"`
	Sound       string            `json:"sound"`
	Key         string            `json:"key"`
	Priority    string            `json:"priority"`
	DeviceToken string            `json:"device_token"`
	DeviceApp   string            `json:"device_app"`
	OnClick     string            `json:"OnClick"`
	TTL         int               `json:"ttl"`
	CustomData  map[string]string `json:"custom_data"`
}

// ToAPNS - transform the notification to APNS format
func (n Notification) ToAPNS() *apns2.Notification {
	pr := apns2.PriorityLow
	npr := strings.ToUpper(n.Priority)
	if npr == "HIGH" {
		pr = apns2.PriorityHigh
	}
	payload := payload.NewPayload()
	if n.Sound != "" {
		payload.SoundName(n.Sound)
	}
	for k, v := range n.CustomData {
		payload.Custom(k, v)
	}
	return &apns2.Notification{
		CollapseID:  n.Key,
		DeviceToken: n.DeviceToken,
		Topic:       n.DeviceApp,
		Expiration:  time.Now().Add(time.Duration(n.TTL) * time.Second),
		Priority:    pr,
		Payload:     payload,
	}
}

// ToAndroid - transform the notification to android format
func (n Notification) ToAndroid() *messaging.Message {
	ttl := time.Duration(n.TTL)
	return &messaging.Message{
		Data:  n.CustomData,
		Token: n.DeviceToken,
		Notification: &messaging.Notification{
			Title: n.Title,
			Body:  n.Body,
		},
		Android: &messaging.AndroidConfig{
			CollapseKey: n.Key,
			Priority:    n.Priority,
			TTL:         &ttl,
			Notification: &messaging.AndroidNotification{
				Sound:       n.Sound,
				ClickAction: n.OnClick,
				Icon:        n.Icon,
			},
		},
	}
}

// ToWeb - transform the notification to a webpush format
func (n Notification) ToWeb() *messaging.Message {
	return &messaging.Message{
		Data:  n.CustomData,
		Token: n.DeviceToken,
		Notification: &messaging.Notification{
			Title: n.Title,
			Body:  n.Body,
		},
		Webpush: &messaging.WebpushConfig{
			Headers: map[string]string{
				"TTL":     strconv.Itoa(n.TTL),
				"Urgency": strings.ToLower(n.Priority),
			},
			Notification: &messaging.WebpushNotification{
				Icon: n.Icon,
				Actions: []*messaging.WebpushNotificationAction{
					{
						Title:  "View",
						Action: n.OnClick,
					},
				},
			},
		},
	}
}
