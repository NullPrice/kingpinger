package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	httpadapter "github.com/NullPrice/kingpinger/pkg/adapters/http"
	rabbitmqadapter "github.com/NullPrice/kingpinger/pkg/adapters/rabbitmq"
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

func rabbitMqHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	job := pinger.PingRequest{}
	err := decoder.Decode(&job)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	// Create a new HTTP adapter since we are making an HTTP callback ping
	rabbitMqAdapter := &rabbitmqadapter.RabbitMQ{PingRequest: job}
	rabbitMqAdapter.SetPingDependency(ping.NewPinger(rabbitMqAdapter.GetPingRequest().Target))
	go pinger.Process(rabbitMqAdapter)
	w.WriteHeader(200)
}

func main() {
	viper.SetConfigName("config")
	viper.AddConfigPath("$HOME/.kingpinger")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		// Log but ignore -- config either malformed or missing
		log.Print(err)
	}
	viper.SetEnvPrefix("pinger")
	viper.BindEnv("rabbitmq")
	viper.BindEnv("port")
	viper.SetDefault("port", "8080")
	viper.BindEnv("host")
	viper.BindEnv("service")
	if viper.IsSet("rabbitmq") && viper.GetBool("rabbitmq") {
		log.Print("Starting in RabbitMQ mode")
		// TODO: We will want to pull from a queue, right now lets just push results to a queue and trigger via http
		var router = mux.NewRouter()
		router.HandleFunc("/", rabbitMqHandler).Methods("POST")
		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", viper.Get("port")), router))
	} else {
		log.Print("Starting in HTTP callback mode")
		var router = mux.NewRouter()
		router.HandleFunc("/", httpHandler).Methods("POST")
		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", viper.Get("port")), router))
	}

}
