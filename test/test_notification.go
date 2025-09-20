package main

import (
	"log"
	"github.com/gen2brain/beeep"
)

func main() {
	log.Println("Testing desktop notification...")
	
	err := beeep.Notify("DegenMon Test", "Desktop notifications are working!", "")
	if err != nil {
		log.Printf("Failed to send notification: %v", err)
	} else {
		log.Println("Notification sent successfully!")
	}
}
