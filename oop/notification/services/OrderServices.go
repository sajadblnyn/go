package services

import (
	"notification/entities"
	"notification/externalServices"
)

func CreateOrder(id int, price int, notifyId string, notifyType externalServices.NotificationType) entities.Order {
	order := entities.Order{Id: id, Price: price, NotifyId: notifyId, NotificationType: notifyType}

	var notifier externalServices.Notifier = externalServices.NewNotifier()
	notifier.SendNotify(order.NotifyId, string(order.NotificationType), "your purchase committed")
	return order
}
