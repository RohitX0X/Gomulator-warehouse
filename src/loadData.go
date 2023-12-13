// eventhub_connect.go

package src

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
)

type eventsequence struct {
	Id               string `json:"id"`
	Workflowsequence []int  `json:"workflowsequence"`
}

// loads data,mapping before sending to an Azure Event Hub.
func LoadData(w http.ResponseWriter, r *http.Request, id string) error {

	filePath := filepath.Join("..", "config", "templates")
	assignmentstartcontent, err := os.ReadFile(filepath.Join(filePath, "assignmentstart.json"))
	if err != nil {
		panic(err)
	}
	assignmentstopcontent, err := os.ReadFile(filepath.Join(filePath, "assignmentstop.json"))
	sessionstartcontent, err := os.ReadFile(filepath.Join(filePath, "sessionstart.json"))
	sessionstopcontent, err := os.ReadFile(filepath.Join(filePath, "sessionstop.json"))
	pickcontent, err := os.ReadFile(filepath.Join(filePath, "pick.json"))
	travelcontent, err := os.ReadFile(filepath.Join(filePath, "travel.json"))
	breakcontent, err := os.ReadFile(filepath.Join(filePath, "break.json"))
	deliverycontent, err := os.ReadFile(filepath.Join(filePath, "delivery.json"))

	workflowseqcontent, err := os.ReadFile(filepath.Join("..", "config", "workflows", id+".json"))
	var workflowseqjson eventsequence

	if err := json.Unmarshal(workflowseqcontent, &workflowseqjson); err != nil {
		panic(err)
	}

	eventmapcontent, err := os.ReadFile(filepath.Join("..", "config", "workflowmap.json"))

	var eventmapcontentjson = make(map[string]string)

	if err := json.Unmarshal(eventmapcontent, &eventmapcontentjson); err != nil {
		panic(err)
	}

	var event_ref_mapping = make(map[string]*[]byte)

	for key, values := range eventmapcontentjson {

		switch values {

		case ("sessionstart"):
			event_ref_mapping[key] = &sessionstartcontent
		case ("sessionstop"):
			event_ref_mapping[key] = &sessionstopcontent
		case ("assignmentstart"):
			event_ref_mapping[key] = &assignmentstartcontent
		case ("assignmentstop"):
			event_ref_mapping[key] = &assignmentstopcontent
		case ("pick"):
			event_ref_mapping[key] = &pickcontent
		case ("travel"):
			event_ref_mapping[key] = &travelcontent
		case ("break"):
			event_ref_mapping[key] = &breakcontent
		case ("delivery"):
			event_ref_mapping[key] = &deliverycontent
		}
	}

	err_t := transformData(id, &event_ref_mapping, workflowseqjson.Workflowsequence)

	if err_t != nil {
		panic(err_t)
	}

	return nil
}
