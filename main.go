package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/florian42/me-api/internal/cmd"
	"github.com/florian42/me-api/internal/presence"
)

type presenceData struct {
	Status     presence.PresenceStatus `json:"status"`
	FocusedApp string                  `json:"focused_app,omitempty"`
}

func meHandler(w http.ResponseWriter, r *http.Request) {
	runner := cmd.Runner()

	focusedApp, _ := presence.GetFocusedApp(runner)

	json.NewEncoder(w).Encode(presenceData{
		Status:     presence.GetStatus(runner),
		FocusedApp: focusedApp,
	})
}

func main() {
	http.HandleFunc("/me", meHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
