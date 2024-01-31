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

	// Define a POST endpoint to trigger the add task
	r.POST("/tasks/sum", func(c *gin.Context) {
		var request struct {
			Numbers []int64 `json:"numbers" binding:"required"`
		}

		// Bind JSON request to the struct
		if err := c.BindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Prepare task parameters
		task := tasks.Signature{
			Name: "sum",
			Args: []tasks.Arg{
				{
					Type:  "[]int64",
					Value: request.Numbers,
				},
			},
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

	r.GET("/tasks/:taskID/status", func(c *gin.Context) {
		taskID := c.Param("taskID")

		// Use the task ID to query the status
		taskState, err := server.GetBackend().GetState(taskID)
		if err != nil {
			// Handle the error and respond with an appropriate HTTP status
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Task not found",
			})
			return
		}

		// Respond with the task status
		c.JSON(http.StatusOK, gin.H{
			"task_id":    taskID,
			"task_state": taskState,
		})
	})

	r.Run(fmt.Sprintf("%s:%s", config.ApiHost, config.ApiPort))
}
