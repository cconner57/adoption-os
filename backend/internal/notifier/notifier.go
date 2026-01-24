package notifier

import (
	"encoding/json"
	"log/slog"
	"os"

	"github.com/SherClockHolmes/webpush-go"
	"github.com/cconner57/adoption-os/backend/internal/data"
)

type Notifier struct {
	Models    data.Models
	Logger    *slog.Logger
	VapidKeys VAPIDKeys
}

type VAPIDKeys struct {
	PublicKey  string
	PrivateKey string
	Subject    string
}

func New(models data.Models, logger *slog.Logger) *Notifier {
	return &Notifier{
		Models: models,
		Logger: logger,
		VapidKeys: VAPIDKeys{
			PublicKey:  os.Getenv("VAPID_PUBLIC_KEY"),
			PrivateKey: os.Getenv("VAPID_PRIVATE_KEY"),
			Subject:    os.Getenv("VAPID_SUBJECT"),
		},
	}
}

func (n *Notifier) SendToAll(message string) {
	subs, err := n.Models.Notifications.GetAll()
	if err != nil {
		n.Logger.Error("Notifier: Failed to fetch subscriptions", "error", err)
		return
	}

	for _, sub := range subs {
		go n.Send(sub, message)
	}
}

// Send sends a notification to a specific subscription
func (n *Notifier) Send(sub *data.NotificationSubscription, message string) {
	// Simple payload structure
	payload := map[string]string{
		"title": "AdoptionOS Alert",
		"body":  message,
		"icon":  "/images/paw.svg",
	}

	jsonPayload, _ := json.Marshal(payload)

	s := &webpush.Subscription{
		Endpoint: sub.Endpoint,
		Keys: webpush.Keys{
			P256dh: sub.P256dh,
			Auth:   sub.Auth,
		},
	}

	resp, err := webpush.SendNotification(jsonPayload, s, &webpush.Options{
		Subscriber:      n.VapidKeys.Subject,
		VAPIDPublicKey:  n.VapidKeys.PublicKey,
		VAPIDPrivateKey: n.VapidKeys.PrivateKey,
		TTL:             30,
	})

	if err != nil {
		n.Logger.Error("Notifier: Failed to send", "endpoint", sub.Endpoint, "error", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == 410 || resp.StatusCode == 404 {
		n.Logger.Info("Notifier: Removing invalid subscription", "endpoint", sub.Endpoint)
		_ = n.Models.Notifications.Delete(sub.Endpoint)
	}
}
