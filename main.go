package main

import (
	"fmt"
	"log"
	"os"

	"github.com/RichardKnop/machinery/v2"
	"github.com/RichardKnop/machinery/v2/config"
	cli "github.com/urfave/cli/v2"

	redisbackend "github.com/RichardKnop/machinery/v2/backends/redis"
	redisbroker "github.com/RichardKnop/machinery/v2/brokers/redis"
	eagerlock "github.com/RichardKnop/machinery/v2/locks/eager"

	"async_queue/api"
	"async_queue/tasks"
	"async_queue/utils"
	"async_queue/worker"
)

func startMachineryServer() (*machinery.Server, error) {
	cnf := &config.Config{
		DefaultQueue:    "machinery_tasks",
		ResultsExpireIn: 3600,
		Redis: &config.RedisConfig{
			MaxIdle:                3,
			IdleTimeout:            240,
			ReadTimeout:            15,
			WriteTimeout:           15,
			ConnectTimeout:         15,
			NormalTasksPollPeriod:  1000,
			DelayedTasksPollPeriod: 500,
		},
	}

	// Create server instance
	broker := redisbroker.NewGR(cnf, []string{fmt.Sprintf("%s:%s", utils.Config.RedisHost, utils.Config.RedisPort)}, 0)
	backend := redisbackend.NewGR(cnf, []string{fmt.Sprintf("%s:%s", utils.Config.RedisHost, utils.Config.RedisPort)}, 0)
	lock := eagerlock.New()
	server := machinery.NewServer(cnf, broker, backend, lock)

	// Register tasks
	tasksMap := map[string]interface{}{
		"dummyTask": tasks.DummyTask,
	}

	return server, server.RegisterTasks(tasksMap)
}

func main() {

	machineryServer, err := startMachineryServer()
	if err != nil {
		log.Fatalf("%+v\n", err)
	}

	cliApp := &cli.App{
		Name:    "Async queue",
		Usage:   "Asynchronous queue app using machinery and gin for api.",
		Authors: []*cli.Author{{Name: "Szymon Pruszek", Email: "szymon.pruszek@gmail.com"}},
		Version: "0.0.1",
		Commands: []*cli.Command{
			{
				Name:  "api",
				Usage: "Start the api. ",
				Action: func(*cli.Context) error {
					log.Println("Starting api server.")
					api.StartApi(&utils.Config, machineryServer)
					return nil
				},
			},
			{
				Name:  "worker",
				Usage: "Run the worker that will execute tasks.",
				Action: func(*cli.Context) error {
					log.Println("Starting worker.")
					worker.StartWorker(machineryServer)
					return nil
				},
			},
		},
	}

	if err := cliApp.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
