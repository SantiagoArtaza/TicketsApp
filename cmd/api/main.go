package main

import (
	"TicketsApp/internal/database"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	dbURL := "postgres://support_user:support_pass@localhost:5432/support_task_manager?sslmode=disable"

	db, err := database.Connect(dbURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		err := db.Ping(context.Background())
		if err != nil {
			http.Error(w, "database not connected", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("API running"))
	})

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("Server running on port", port)

	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
