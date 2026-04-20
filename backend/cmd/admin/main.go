package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/SHP-Association/E-learningWeb/backend/pkg/log"
	"github.com/SHP-Association/E-learningWeb/backend/pkg/services"
	"github.com/SHP-Association/E-learningWeb/backend/ent/user"
)

// main creates a new admin user with the email passed in via the flag.
func main() {
	// Start a new container.
	c := services.NewContainer()
	defer func() {
		// Gracefully shutdown all services.
		if err := c.Shutdown(); err != nil {
			log.Default().Error("shutdown failed", "error", err)
		}
	}()

	var email string
	flag.StringVar(&email, "email", "", "email address for the admin user")
	flag.Parse()

	if len(email) == 0 {
		invalid("email is required")
	}

	// Generate a password.
	pw, err := c.Auth.RandomToken(10)
	if err != nil {
		invalid("failed to generate a random password")
	}

	// Check if the user already exists.
	u, err := c.ORM.User.
		Query().
		Where(user.Email(email)).
		Only(context.Background())

	if err == nil {
		// Update existing user.
		err = u.Update().
			SetPassword(pw).
			SetAdmin(true).
			SetVerified(true).
			Exec(context.Background())
		
		if err != nil {
			invalid(err.Error())
		}
		fmt.Println("")
		fmt.Println("-- ADMIN USER UPDATED --")
	} else {
		// Create the admin user.
		err = c.ORM.User.
			Create().
			SetEmail(email).
			SetUsername(email).
			SetFirstName("Admin").
			SetAdmin(true).
			SetVerified(true).
			SetPassword(pw).
			Exec(context.Background())

		if err != nil {
			invalid(err.Error())
		}
		fmt.Println("")
		fmt.Println("-- ADMIN USER CREATED --")
	}

	fmt.Printf("Email: %s\n", email)
	fmt.Printf("Password: %s\n", pw)
	fmt.Println("----")
	fmt.Println("")
}

func invalid(msg string) {
	fmt.Printf("[ERROR] %s\n", msg)
	os.Exit(1)
}
