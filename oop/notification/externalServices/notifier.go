package externalServices

type NotificationType string

const (
	Email NotificationType = "Email"
	Sms   NotificationType = "Sms"
)

type Notifier interface {
	SendNotify(reciver string, notify_type NotificationType, message string)
}

func NewNotifier() Notifier {
	// var notifier Notifier = porslineService{}
	var notifier Notifier = kavenegarService{}
	return notifier
}
