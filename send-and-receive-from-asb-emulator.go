package main

import (
	"context"
	_ "embed"
	"io"
	"log"
	"strings"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus"
	"github.com/testcontainers/testcontainers-go/modules/azure/servicebus"
)

//go:embed servicebus-config.json
var serviceBusConfigJSON string

const queueName = "queue.1"

func main() {
	ctx := context.Background()

	// 1. Start Service Bus Emulator – module handles MSSQL + network automatically!
	// Use default wait strategy
	sbContainer, err := servicebus.Run(
		ctx,
		"mcr.microsoft.com/azure-messaging/servicebus-emulator:1.1.2",
		servicebus.WithAcceptEULA(),
		servicebus.WithConfig(strings.NewReader(serviceBusConfigJSON)),
	)
	if err != nil {
		// Print logs on failure
		logs, logErr := sbContainer.Logs(ctx)
		if logErr == nil {
			if logBytes, readErr := io.ReadAll(logs); readErr == nil {
				log.Printf("Emulator logs:\n%s\n", string(logBytes))
			}
			logs.Close()
		}
		log.Fatalf("Failed to start Service Bus emulator: %v", err)
	}
	defer sbContainer.Terminate(ctx) // Auto-terminates MSSQL + network too!

	// 2. Get connection string (logs show it's ready with this)
	connStr, err := sbContainer.ConnectionString(ctx)
	if err != nil {
		log.Fatalf("Failed to get connection string: %v", err)
	}
	log.Printf("Connected to emulator: %s", connStr[:60]+"...")

	// 3. Create client
	client, err := azservicebus.NewClientFromConnectionString(connStr, nil)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close(ctx)

	// 4. Send message
	sender, err := client.NewSender(queueName, nil)
	if err != nil {
		log.Fatalf("Sender error: %v", err)
	}
	defer sender.Close(ctx)

	msg := &azservicebus.Message{Body: []byte("Hello message sent!")}
	if err := sender.SendMessage(ctx, msg, nil); err != nil {
		log.Fatalf("Send failed: %v", err)
	}
	log.Println("Message sent")

	// 5. Receive message (with context timeout)
	receiver, err := client.NewReceiverForQueue(queueName, nil)
	if err != nil {
		log.Fatalf("Receiver error: %v", err)
	}
	defer receiver.Close(ctx)

	receiveCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	messages, err := receiver.ReceiveMessages(receiveCtx, 1, nil)
	if err != nil {
		log.Fatalf("Receive failed: %v", err)
	}

	log.Printf("Received %d message(s)", len(messages))
	for _, m := range messages {
		log.Printf("Body: %s", string(m.Body))
		if err := receiver.CompleteMessage(ctx, m, nil); err != nil {
			log.Printf("Complete failed: %v", err)
		}
	}

	log.Println("Demo completed successfully!")
}
