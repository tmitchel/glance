package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pacedotdev/oto/otohttp"
	"github.com/sirupsen/logrus"
	"github.com/tmitchel/glance"
	"github.com/tmitchel/glance/generated"
)

func main() {
	c := glance.CreateService{}
	g := glance.GetService{}
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
