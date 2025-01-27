package main

import (
	"bacchus/node"
	"bacchus/task"
	"bacchus/manager"
	"bacchus/worker"
	"fmt"
	"time"

	"github.com/golang-collections/collections/queue"
    "github.com/google/uuid"
)

func main() {
	t := task.Task{
		ID: uuid.New(),
		Name: "Sipserver",
		State: task.Pending,
		Image: "sipserver-1",
		Memory: 1024,
		Disk: 1,

	}

	te := task.TaskEvent
}