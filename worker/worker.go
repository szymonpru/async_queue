package worker

import (
	"github.com/RichardKnop/machinery/v2"
)

func StartWorker(server *machinery.Server) error {
	consumerTag := "machinery_worker"

	// The second argument is a consumer tag
	// Ideally, each worker should have a unique tag (worker1, worker2 etc)
	worker := server.NewWorker(consumerTag, 0)

	return worker.Launch()

}
