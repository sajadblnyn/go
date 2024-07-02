package externalServices

import "fmt"

type porslineService struct {
}

func (kave porslineService) SendNotify(reciver string, notify_type string, message string) {
	fmt.Printf("sent %s : %s to %s by PorseLine\n", notify_type, message, reciver)
}
