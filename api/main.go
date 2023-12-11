package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type workflow struct {
	id               string
	workflowsequence []string
}

func workflowHandler(w http.ResponseWriter, r *http.Request) {
	// Extracting the ID from the URL path
	id := strings.TrimPrefix(r.URL.Path, "/api/workflow/")

	switch r.Method {
	case http.MethodGet:
		getWorkflow(w, id)
	case http.MethodPost:
		createWorkflow(w, r, id)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func simulateHandler(w http.ResponseWriter, r *http.Request) {
	// Extracting the ID from the URL path
	id := strings.TrimPrefix(r.URL.Path, "/api/workflow/")

	switch r.Method {
	case http.MethodPost:
		Simulatedata(w, r, id)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func getWorkflow(w http.ResponseWriter, id string) {
	// Reading the workflow content from the file
	filePath := filepath.Join("..", "config", "workflows", id+".json")
	content, err := os.ReadFile(filePath)
	if err != nil {
		http.Error(w, "Workflow not found", http.StatusNotFound)
		return
	}

	// Responding with the workflow content
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(content)
}

func createWorkflow(w http.ResponseWriter, r *http.Request, id string) {
	// Reading the request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}
	// body_json = json.NewEncoder(string(body))

	fmt.Printf("Serrunnn :%d\n", string(body))
	fmt.Printf("iddddd :%d\n", id)
	// newWorkflow := workflow{id: id, workflowsequence: sequence}

	// Saving the workflow to a file
	saveWorkflowToFile(id, body)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(body)
}

func saveWorkflowToFile(id string, workflow []byte) {
	// Creating the directory if it doesn't exist
	err := os.MkdirAll(filepath.Join("..", "config", "workflows"), os.ModePerm)
	if err != nil {
		fmt.Println("Error creating directory:", err)
		return
	}

	// Writing the workflow content to a file
	filePath := filepath.Join("..", "config", "workflows", id+".json")
	fmt.Printf("Se:%s\n", filePath)
	err = os.WriteFile(filePath, []byte(workflow), 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
	}
}

func main() {
	// r := mux.NewRouter()

	// r.HandleFunc("/workflow/{id}",getWorkflow).Methods("GET")
	// r.HandleFunc("/operator/{id}",getOperator).Methods("GET")
	// r.HandleFunc("/workflow/{id}",createWorkflow).Methods("POST")
	// r.HandleFunc("/operator/{id}",createOperator).Methods("POST")
	http.HandleFunc("/api/workflow/", workflowHandler)
	http.HandleFunc("/api/simulate/", simulateHandler)

	port := 3000
	fmt.Printf("Server is running on :%d\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}

}
