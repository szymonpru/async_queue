## Experimental Async Queue using Go,Docker,Gin,Go-Machinery
This is a simple, no-frills application designed for testing the capabilities of Golang and a few libraries. 
Using Docker, Gin, and Go-Machinery v2, I've put together a trio of containers: Redis, API, and Worker.
It's a bit rough around the edges, but it gets the job done for some casual Golang experimentation.

### Quick Start Guide

1. Grab the Code
```
git clone https://github.com/szymonpru/async_queue
cd async_queue
```
2. Fire it Up
```
docker-compose up -d
```

### Playtime

#### Launch a Task
To launch a task, toss a POST request at `/tasks/sum` with a payload of numbers.

Example using curl:
```
curl -X POST -H "Content-Type: application/json" -d '{"numbers": [1, 2, 3]}' http://localhost:8080/tasks/sum
```

#### Spy on Task Progress and Results
Send a GET request to `/tasks/<task_id>/status` to peek at the task's status and results.

Example with curl:
```
curl http://localhost:8080/tasks/<task_id>/status
```

### API Whims
* POST `/tasks/sum`: Drop a task with a bunch of numbers.
* GET `/tasks/<task_id>/status`: Check out how your task is doing.