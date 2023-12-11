// eventhub_connect.go

package eventhub

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
)

// EventData represents the structure of data to be sent to the Event Hub.
type EventData struct {
	Message string `json:"message"`
	// Add other fields as needed
}

// PushToEventHub sends data to an Azure Event Hub.
func PushToEventHub(data EventData) error {
	// Azure Event Hub connection string
	eventHubConnStr := os.Getenv("EVENTHUB_CONNECTION_STRING")
	if eventHubConnStr == "" {
		return fmt.Errorf("EVENTHUB_CONNECTION_STRING not set")
	}

	// Event Hub name
	eventHubName := "your_event_hub_name" // Replace with your Event Hub name

	// Create an Event Hub client
	hub, err := eventhubs.NewHubFromConnectionString(eventHubConnStr, eventHubName)
	if err != nil {
		return fmt.Errorf("error creating Event Hub client: %v", err)
	}

	// Serialize data to JSON
	messageBody, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("error serializing data to JSON: %v", err)
	}

	// Create an event to be sent
	event := eventhubs.NewEventFromString(string(messageBody))

	// Send the event
	err = hub.Send(context.Background(), event)
	if err != nil {
		return fmt.Errorf("error sending event to Event Hub: %v", err)
	}

	fmt.Println("Data sent to Event Hub successfully")
	return nil
}
