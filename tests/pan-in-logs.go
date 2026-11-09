package rules

import "log"

type Card struct {
	CardNumber string
	CVV        string
	ID         string
}

func charge(c Card, cardNumber string) {
	// ruleid: pan-in-logs-go
	log.Printf("charging %s", cardNumber)
	// ruleid: pan-in-logs-go
	log.Println("declined", c.CVV)

	// ok: pan-in-logs-go
	log.Printf("charging card %s", c.ID)
}
