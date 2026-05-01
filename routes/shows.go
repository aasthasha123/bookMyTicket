package routes

import (
	"bookMyTicket/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func ShowsRoute() http.HandlerFunc {
	events := models.GetEvents()
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			w.Write(eventsToJSON(events))
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}
func GetShowsRouteByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := mux.Vars(r)["name"]
		fmt.Println("NAME ", name)
		events := models.GetEvents()
		for _, event := range events {
			if event.Name == name {
				fmt.Println("EVENT ", event)
				w.Write(eventsToJSON([]models.Event{event}))
			}
		}
	}
}

func eventsToJSON(events []models.Event) []byte {
	jsonData, err := json.Marshal(events)
	if err != nil {
		return []byte("[]")
	}
	return jsonData
}
