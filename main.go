package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/NullPrice/kingpinger/kingping"
	"github.com/gorilla/mux"

	"github.com/sparrc/go-ping"

	"github.com/spf13/viper"
)

type Config struct {
	port    string
	host    string
	service string
}

func handler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	test := kingping.JobRequest{}
	err := decoder.Decode(&test)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	// w.WriteHeader(200)
	fmt.Fprintf(w, "%+v", test)
}

func handlerUnused(w http.ResponseWriter, r *http.Request) {
	pinger, err := ping.NewPinger("www.google.com")
	pinger.SetPrivileged(true)
	if err != nil {
		panic(err)
	}
	pinger.Count = 3
	pinger.Run()
	fmt.Fprintf(w, "%+v", pinger.Statistics())
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
