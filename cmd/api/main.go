package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/user/car-project/internal/config"
	"github.com/user/car-project/internal/db"
	"github.com/user/car-project/internal/routes"
	"github.com/user/car-project/internal/utils"
)

// @title           Car Project API
// @version         1.0
// @description     This is a sample server for a car management system.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

func main() {
	// Initialize logger
	utils.InitLogger()
	logger := utils.GetLogger()

	// Load config
	cfg := config.LoadConfig()

	// Initialize database
	// Note: Uncomment this if you have a running postgres instance
	if cfg.DBURL != "" {
		db.InitDB(cfg.DBURL)
	}

	// Setup routes
	r := routes.SetupRouter(cfg.JWTSecret)

	srv := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: r,
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		logger.Printf("Starting server on port %s\n", cfg.Port)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be caught, so no need to add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Println("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatalf("Server forced to shutdown: %v", err)
	}

	logger.Println("Server exiting")
}
