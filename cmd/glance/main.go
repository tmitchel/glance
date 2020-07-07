package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/pacedotdev/oto/otohttp"
	"github.com/sirupsen/logrus"
	"github.com/tmitchel/glance"
	"github.com/tmitchel/glance/generated"
)

func main() {
	// load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		logrus.Error("Error loading .env file")
	}

	db, err := glance.OpenDatabase(fmt.Sprintf("host=db port=5432 user=%s "+
		"password=%s dbname=%s sslmode=disable",
		os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB")))
	if err != nil {
		logrus.Fatal(err)
	}

	c, _ := glance.NewCreateService(db)
	g, _ := glance.NewGetService(db)
	server := otohttp.NewServer()
	generated.RegisterCreateService(server, c)
	generated.RegisterGetService(server, g)

	r := mux.NewRouter()
	r.PathPrefix("/build").Handler(http.FileServer(http.Dir("frontend/public/")))
	r.PathPrefix("/oto").Handler(server)
	r.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "frontend/public/index.html")
	})

	logrus.Fatal(http.ListenAndServe(":8080", r))
}
