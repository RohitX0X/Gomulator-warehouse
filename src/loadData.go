// eventhub_connect.go

package eventhub

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

const operatorstring string = "operator"

// PushToEventHub sends data to an Azure Event Hub.
func loadData(id string) error {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	filePath := filepath.Join("..", "config", "templates")
	assignmentstartcontent, err := os.ReadFile(filepath.Join(filePath, "assignmentstart.json"))
	assignmentstopcontent, err := os.ReadFile(filepath.Join(filePath, "assignmentstop.json"))
	sessionstartcontent, err := os.ReadFile(filepath.Join(filePath, "sessionstart.json"))
	sessionstopcontent, err := os.ReadFile(filepath.Join(filePath, "sessionstop.json"))
	pickcontent, err := os.ReadFile(filepath.Join(filePath, "pick.json"))
	travelcontent, err := os.ReadFile(filepath.Join(filePath, "travel.json"))
	breakcontent, err := os.ReadFile(filepath.Join(filePath, "break.json"))
	deliverycontent, err := os.ReadFile(filepath.Join(filePath, "delivery.json"))

	eventmapcontent, err := os.ReadFile(filepath.Join("..", "config", "workflowmap.json"))

	var eventmapcontentjson = make(map[string]string)

	if err := json.Unmarshal(eventmapcontent, &eventmapcontentjson); err != nil {
		panic(err)
	}

	var event_ref_mapping = make(map[string]*[]byte)

	for key, values := range eventmapcontentjson {

		switch values {

		case ("sessionstartcontent"):
			event_ref_mapping[key] = &sessionstartcontent
		case ("sessionstopcontent"):
			event_ref_mapping[key] = &sessionstopcontent
		case ("assignmentstartcontent"):
			event_ref_mapping[key] = &assignmentstartcontent
		case ("assignmentstopcontent"):
			event_ref_mapping[key] = &assignmentstopcontent
		case ("pickcontent"):
			event_ref_mapping[key] = &pickcontent
		case ("travelcontent"):
			event_ref_mapping[key] = &travelcontent
		case ("breakcontent"):
			event_ref_mapping[key] = &breakcontent
		case ("deliverycontent"):
			event_ref_mapping[key] = &deliverycontent
		}
	}

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
