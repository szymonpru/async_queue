package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/RichardKnop/machinery/v2"
	"github.com/RichardKnop/machinery/v2/tasks"
	"github.com/gin-gonic/gin"

	"async_queue/utils"
)

func StartApi(config *utils.Configuration, server *machinery.Server) {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// Define a GET endpoint to initiate the task
	r.GET("/task/dummy", func(c *gin.Context) {
		task := tasks.Signature{
			Name: "dummyTask",
		}

		// Send the task to the worker asynchronously
		asyncResult, err := server.SendTask(&task)
		if err != nil {
			log.Fatal(err)
		}

		// Respond with the task ID
		c.JSON(http.StatusOK, gin.H{
			"task_id": asyncResult.GetState().TaskUUID,
		})
	})

	// Define a GET endpoint to check the status of the task by ID
	r.GET("/task/:taskID/status", func(c *gin.Context) {
		taskID := c.Param("taskID")

		// Use the task ID to query the status
		taskState, err := server.GetBackend().GetState(taskID)
		if err != nil {
			log.Fatal(err)
		}

		// Respond with the task status
		c.JSON(http.StatusOK, gin.H{
			"task_id":    taskID,
			"task_state": taskState,
		})
	})

	r.Run(fmt.Sprintf("%s:%s", config.ApiHost, config.ApiPort))
}
