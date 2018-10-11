package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	httpadapter "github.com/NullPrice/kingpinger/pkg/adapters/http"
	ping "github.com/sparrc/go-ping"

	pinger "github.com/NullPrice/kingpinger/pkg/pinger"
	"github.com/gorilla/mux"

	"github.com/spf13/viper"
)

func httpHandler(w http.ResponseWriter, r *http.Request) {
	// Handle decoding the HTTP request body
	decoder := json.NewDecoder(r.Body)
	job := pinger.PingRequest{}
	err := decoder.Decode(&job)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	// Create a new HTTP adapter since we are making an HTTP callback ping
	httpAdapter := &httpadapter.HTTP{PingRequest: job}
	httpAdapter.SetPingDependency(ping.NewPinger(httpAdapter.GetPingRequest().Target))
	go pinger.Process(httpAdapter)
	w.WriteHeader(200)
}

func main() {
	viper.SetEnvPrefix("pinger")
	viper.BindEnv("port")
	viper.BindEnv("host")
	viper.BindEnv("service")
	viper.SetDefault("port", "8080")
	var router = mux.NewRouter()
	router.HandleFunc("/", httpHandler).Methods("POST")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", viper.Get("port")), router))
}
