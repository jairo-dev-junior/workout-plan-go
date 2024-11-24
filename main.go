package main

import (
	"log"
	"net/http"

	router "github.com/jairo-dev-junior/workout-plan-go/view"
	"github.com/joho/godotenv"
)

func main() {
	var err = godotenv.Load(".env")
	if err != nil {
		log.Fatal("cannot get environment variables")
	}
	http.ListenAndServe(":3000", router.Route())
}
