package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	dapr "github.com/dapr/go-sdk/client"
	"github.com/gorilla/mux"
)

var (
	STATESTORE_NAME = getEnv("STATESTORE_NAME", "statestore")
	TENANT_ID       = getEnv("TENANT_ID", "tenant-a")
)

type MyValues struct {
	Values []string
}

func readHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	daprClient, err := dapr.NewClient()
	if err != nil {
		panic(err)
	}

	result, err := daprClient.GetState(ctx, STATESTORE_NAME, fmt.Sprintf("%s-%s", TENANT_ID, "values"), nil)
	if err != nil {
		panic(err)
	}
	myValues := MyValues{}
	json.Unmarshal(result.Value, &myValues)

	respondWithJSON(w, http.StatusOK, myValues)
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func main() {
	appPort := os.Getenv("APP_PORT")
	if appPort == "" {
		appPort = "8080"
	}

	r := mux.NewRouter()

	// Dapr subscription routes orders topic to this route
	r.HandleFunc("/", readHandler).Methods("GET")

	// Add handlers for readiness and liveness endpoints
	r.HandleFunc("/health/{endpoint:readiness|liveness}", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	})

	log.Printf("Starting Read App in Port: %s", appPort)
	// Start the server; this is a blocking call
	err := http.ListenAndServe(":"+appPort, r)
	if err != http.ErrServerClosed {
		log.Panic(err)
	}
}

// getEnv returns the value of an environment variable, or a fallback value if
func getEnv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		value = fallback
	}
	return value
}
