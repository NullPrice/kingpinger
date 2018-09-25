package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/sparrc/go-ping"

	"github.com/spf13/viper"
)

func handler(w http.ResponseWriter, r *http.Request) {
	pinger, err := ping.NewPinger("www.google.com")
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

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", viper.Get("port")), nil))
}
