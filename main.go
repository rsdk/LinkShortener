package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"linkShortener/config"
	"linkShortener/controllers"
	"linkShortener/models"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	config.Init()
	models.InitDB()

	// Create context that listens for the interrupt signal from the OS.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	router := gin.Default()
	var urlPrefix string
	if len(viper.GetString("urlPrefix")) > 0 {
		urlPrefix += "/" + viper.GetString("urlPrefix")
	}
	router.PUT(urlPrefix+"/shorten", controllers.ShortenURL)
	router.GET(urlPrefix+"/:id", controllers.GetURL)

	serverAddress := viper.GetString("IP") + ":" + viper.GetString("Port")
	srv := &http.Server{
		Addr:    serverAddress,
		Handler: router,
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	// Listen for the interrupt signal.
	<-ctx.Done()

	// Restore default behavior on the interrupt signal and notify user of shutdown.
	stop()
	log.Println("shutting down gracefully, press Ctrl+C again to force")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}
	log.Println("Server exiting")
}
