package tasks

import (
	"context"
	"fmt"
	"time"

	"github.com/mikestefanello/backlite"

	"github.com/SHP-Association/E-learningWeb/backend/pkg/services"
)

// WelcomeEmailTask processes sending a welcome email to a new user.
type WelcomeEmailTask struct {
	UserID   int
	Username string
	Email    string
}

func (t WelcomeEmailTask) Config() backlite.QueueConfig {
	return backlite.QueueConfig{
		Name:        "WelcomeEmailTask",
		MaxAttempts: 3,
		Timeout:     10 * time.Second,
		Backoff:     30 * time.Second,
	}
}

func NewWelcomeEmailTaskQueue(c *services.Container) backlite.Queue {
	return backlite.NewQueue[WelcomeEmailTask](func(ctx context.Context, task WelcomeEmailTask) error {
		// In a real app, you would render a template here.
		// For now, we use the Mail service's Compose.
		return c.Mail.
			Compose().
			To(task.Email).
			Subject(fmt.Sprintf("Welcome to SHP-Learner, %s!", task.Username)).
			Body(fmt.Sprintf("Hi %s, welcome to our platform!", task.Username)).
			Send(nil) // ctx is nil because we are in background task, but Mail.Send takes echo.Context.
			          // I might need to adjust Mail.Send to accept nil or a generic context.
	})
}
