package main

import (
	"flag"
	"log"

	"github.com/whatsapp-signal-bridge/bridge"
	"github.com/whatsapp-signal-bridge/signal"
	"github.com/whatsapp-signal-bridge/whatsapp"
	"github.com/whatsapp-signal-bridge/whatsmeow"
)

func main() {
	botNumber := flag.String("bot", "", "phone number of bot (starting with +)")
	receiverNumber := flag.String("receiver", "", "phone number of receiver (starting with +)")
	flag.Parse()

	if *botNumber == "" {
		log.Fatal("Please provide the botNumber")
	}
	if *receiverNumber == "" {
		log.Fatal("Please provide the receiverNumber")
	}

	queue := bridge.NewQueue()

	go signal.StartClient(signal.SignalClientOptions{
		Queue:          queue,
		BotNumber:      *botNumber,
		ReceiverNumber: *receiverNumber,
	})

	go whatsapp.StartClient(whatsapp.WhatsappClientOptions{
		Queue: queue,
	})

	go whatsmeow.StartClient(whatsmeow.WhatsmeowClientOptions{
		Queue: queue,
	})

	select {}
}
