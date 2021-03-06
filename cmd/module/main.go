package module

import (
	"log"
	"os"
	"time"

	"github.com/qmstr/synclib/pkg/module/rabbitmq"
)

//nolint:deadcode,unused
// This function is intentionally unused,
// since this repository is a GitHub Template
// (https://docs.github.com/en/github/creating-cloning-and-archiving-repositories/creating-a-template-repository),
// and other module repositories are supposed
// to be generated from this one.
func main() {

	// Connecting to RabbitMQ
	conn, ch := rabbitmq.Connect(os.Getenv("RABBITMQ_ADDRESS"))
	defer conn.Close()
	defer ch.Close()

	// Declaring queues
	queue, _ := rabbitmq.DeclareQueues(ch, os.Getenv("QUEUE_NAME"))

	// Dummy callback function
	dummyCallbackFunction := func() error {
		dummyComputationTimeInSeconds := 15
		log.Printf("Performing dummy computation for %d seconds...", dummyComputationTimeInSeconds)
		time.Sleep(time.Duration(dummyComputationTimeInSeconds) * time.Second)
		return nil // dummy function produces no error
	}

	// Executing the callback function upon receiving the order message
	responded := rabbitmq.OnMessageReceive(ch, queue.Name, dummyCallbackFunction)
	<-responded
}
