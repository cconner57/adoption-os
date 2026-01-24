package main

import (
	"fmt"
	"log"

	"github.com/SherClockHolmes/webpush-go"
)

func main() {
	privateKey, publicKey, err := webpush.GenerateVAPIDKeys()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("VAPID_PRIVATE_KEY=%s\n", privateKey)
	fmt.Printf("VAPID_PUBLIC_KEY=%s\n", publicKey)
	fmt.Println("VAPID_SUBJECT=mailto:admin@idohr.app")
}
