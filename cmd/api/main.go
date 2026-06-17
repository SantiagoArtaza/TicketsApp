package main

import (
	"TicketsApp/internal/config"
	"TicketsApp/internal/database"
	"TicketsApp/internal/users"
	"log"
	"net/http"
)

func main() {
	cfg := config.Load()

	db, err := database.Connect(cfg.DatabaseURL())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	userRepository := users.NewRepository(db)
	userService := users.NewService(userRepository)
	userHandler := users.NewHandler(userService)

	http.HandleFunc("/users", userHandler.HandleUsers)

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("API running"))
	})

	log.Println("Server running on port", cfg.AppPort)

	err = http.ListenAndServe(":"+cfg.AppPort, nil)
	if err != nil {
		log.Fatal(err)
	}
}
