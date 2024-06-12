package entities

import "notification/externalServices"

type Order struct {
	Id               int
	Price            int
	NotifyId         string
	NotificationType externalServices.NotificationType
}
