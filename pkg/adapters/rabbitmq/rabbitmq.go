package rabbitmq

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/spf13/viper"

	"github.com/NullPrice/kingpinger/pkg/pinger"
	ping "github.com/sparrc/go-ping"
	"github.com/streadway/amqp"
)

type RabbitMQ struct {
	Result      pinger.Result
	PingRequest pinger.PingRequest
	Ping        *ping.Pinger
}

func logError(err error, msg string) {
	if err != nil {
		log.Printf("%s: %s", msg, err)
	}
}

// ProcessResult - Handles dealing with the result
func (adapter *RabbitMQ) ProcessResult() {
	// auth := amqp.PlainAuth{Username: viper.GetString("username"), Password: viper.GetString("Something")}
	// log.Print(auth.Response())
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s/", viper.GetString("username"), viper.GetString("password"), viper.GetString("host"), viper.GetString("rabbit_port")))
	logError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	logError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"results", // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	logError(err, "Failed to declare a queue")

	body, err := json.Marshal(adapter.Result)
	logError(err, "Failed to marshal JSON result data")

	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(body),
		})
	log.Printf(" [x] Sent %s", body)
	logError(err, "Failed to publish a message")
}

// SetResult - Sets the result value
func (adapter *RabbitMQ) SetResult(result pinger.Result) {
	adapter.Result = result
}

// GetResult - Gets result value
func (adapter *RabbitMQ) GetResult() pinger.Result {
	return adapter.Result
}

// SetPingRequest - Sets ping request
func (adapter *RabbitMQ) SetPingRequest(request pinger.PingRequest) {
	adapter.PingRequest = request
}

// GetPingRequest - Gets ping request
func (adapter *RabbitMQ) GetPingRequest() pinger.PingRequest {
	return adapter.PingRequest
}

// Run - Runs a ping process and sets updates the result struct
func (adapter *RabbitMQ) Run() {
	if adapter.Ping == nil {
		// If pinger has not been set manually
		adapter.SetPingDependency(ping.NewPinger(adapter.GetPingRequest().Target))
	}
	adapter.Ping.Run()
	adapter.Result = pinger.Result{JobID: adapter.PingRequest.JobID, Statistics: adapter.Ping.Statistics()}
}

// SetPingDependency - Sets the ping dependency
func (adapter *RabbitMQ) SetPingDependency(x *ping.Pinger, err error) {
	if err != nil {
		// TODO: We should handle this: this does an os.exit behind the scenes, we want to handle all errors as this is a client
		log.Fatalln(err)
	}
	x.SetPrivileged(true)
	x.Count = adapter.PingRequest.Count
	adapter.Ping = x
}

// GetPingDependency - Gets the ping dependency struct
func (adapter *RabbitMQ) GetPingDependency() *ping.Pinger {
	return adapter.Ping
}
