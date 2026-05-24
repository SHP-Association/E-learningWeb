package tasks

import (
	"github.com/SHP-Association/E-learningWeb/backend/pkg/services"
)

// Register registers all task queues with the task client.
func Register(c *services.Container) {
	c.Tasks.Register(NewExampleTaskQueue(c))
	c.Tasks.Register(NewWelcomeEmailTaskQueue(c))
	c.Tasks.Register(NewSendOTPTaskQueue(c))
}
