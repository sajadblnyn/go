package externalServices

import "fmt"

type kavenegarService struct {
}

func (kave kavenegarService) SendNotify(receptor string, notify_type string, message string) {
	fmt.Printf("sent %s : %s to %s by kavenegar\n", notify_type, message, receptor)
}
