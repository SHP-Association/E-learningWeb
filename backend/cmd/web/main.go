package main

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"github.com/SHP-Association/E-learningWeb/backend/pkg/handlers"
	"github.com/SHP-Association/E-learningWeb/backend/pkg/log"
	"github.com/SHP-Association/E-learningWeb/backend/pkg/services"
	"github.com/SHP-Association/E-learningWeb/backend/pkg/tasks"
	"net"
	"os/exec"
	"time"
)

func main() {
	// Start a new container.
	c := services.NewContainer()
	defer func() {
		// Gracefully shutdown all services.
		fatal("shutdown failed", c.Shutdown())
	}()

	// Build the router.
	if err := handlers.BuildRouter(c); err != nil {
		fatal("failed to build the router", err)
	}

	// Register all task queues.
	tasks.Register(c)

	// Start the task runner to execute queued tasks.
	c.Tasks.Start(context.Background())

	// Clear the port if it's already in use.
	clearPort(c.Config.HTTP.Hostname, c.Config.HTTP.Port)

	// Start the server.
	go func() {
		srv := http.Server{
			Addr:         fmt.Sprintf("%s:%d", c.Config.HTTP.Hostname, c.Config.HTTP.Port),
			Handler:      c.Web,
			ReadTimeout:  c.Config.HTTP.ReadTimeout,
			WriteTimeout: c.Config.HTTP.WriteTimeout,
			IdleTimeout:  c.Config.HTTP.IdleTimeout,
		}

		if c.Config.HTTP.TLSEnabled {
			certs, err := tls.LoadX509KeyPair(c.Config.HTTP.TLSCertificate, c.Config.HTTP.TLSKey)
			fatal("cannot load TLS certificate", err)

			srv.TLSConfig = &tls.Config{
				Certificates: []tls.Certificate{certs},
			}
		}

		if err := c.Web.StartServer(&srv); errors.Is(err, http.ErrServerClosed) {
			fatal("shutting down the server", err)
		}
	}()

	// Wait for interrupt signal to gracefully shut down the web server and task runner.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	signal.Notify(quit, os.Kill)
	<-quit
}

// fatal logs an error and terminates the application, if the error is not nil.
func fatal(msg string, err error) {
	if err != nil {
		log.Default().Error(msg, "error", err)
		os.Exit(1)
	}
}

// clearPort checks if a port is in use and kills the process using it after a countdown.
func clearPort(hostname string, port uint16) {
	addr := fmt.Sprintf("%s:%d", hostname, port)
	ln, err := net.Listen("tcp", addr)
	if err == nil {
		_ = ln.Close()
		return
	}

	fmt.Printf("\n⚠️  Port %d is already in use. Automatically clearing in 10 seconds...\n", port)
	for i := 10; i > 0; i-- {
		fmt.Printf("🔥 Clearing in %d seconds...\r", i)
		time.Sleep(time.Second)
	}
	fmt.Printf("\n🚀 Clearing port %d now!\n", port)

	// Execute fuser -k to kill the process on the port
	cmd := exec.Command("fuser", "-k", fmt.Sprintf("%d/tcp", port))
	_ = cmd.Run()

	// Give the OS a moment to fully release the port
	time.Sleep(time.Second * 2)
}
