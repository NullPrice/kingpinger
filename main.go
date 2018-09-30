package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/NullPrice/kingpinger/kingping"
	"github.com/gorilla/mux"

	"github.com/spf13/viper"
)

type Config struct {
	port    string
	host    string
	service string
}

func handler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	job := kingping.JobRequest{}
	err := decoder.Decode(&job)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	job.Process()
	log.Printf("%+v", &job)
	w.WriteHeader(200)
}

func main() {
	viper.SetEnvPrefix("pinger")
	viper.BindEnv("port")
	viper.BindEnv("host")
	viper.BindEnv("service")
	viper.SetDefault("port", "8080")
	var router = mux.NewRouter()
	router.HandleFunc("/", handler).Methods("POST")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", viper.Get("port")), router))
}
