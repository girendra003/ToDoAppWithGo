package main

import (

	"github.com/girendra003/gin-mongodb-api/config"
	"github.com/girendra003/gin-mongodb-api/cronjobs"
	"github.com/girendra003/gin-mongodb-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	//run database
	config.DBInstance()

	cronjobs.Ok()

	//routes
	routes.UserRoute(router)

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to Gin MongoDB API",
		})
	})


	router.Run(":8000")

}




//graceful shutdown
// package main

// import (
// 	"context"
// 	"log"
// 	"os"
// 	"os/signal"
// 	"syscall"
// 	"time"

// 	"github.com/gin-gonic/gin"
// 	"github.com/girendra003/gin-mongodb-api/config"
// 	"github.com/girendra003/gin-mongodb-api/routes"
// 	"net/http"
// )

// func main() {
// 	// Create Gin router
// 	router := gin.Default()

// 	// Initialize MongoDB connection (also sets global config.Client)
// 	config.DBInstance()

// 	// Define API routes
// 	routes.UserRoute(router)

// 	// Basic test route
// 	router.GET("/", func(c *gin.Context) {
// 		c.JSON(http.StatusOK, gin.H{
// 			"message": "Welcome to Gin MongoDB API",
// 		})
// 	})

// 	// Create custom HTTP server so we can shut it down gracefully
// 	srv := &http.Server{
// 		Addr:    ":8000",
// 		Handler: router,
// 	}

// 	// Start server in a separate goroutine
// 	go func() {
// 		log.Println("🚀 Server is running at http://localhost:8000")
// 		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
// 			log.Fatalf("❌ Server error: %v", err)
// 		}
// 	}()

// 	// Set up signal channel to catch OS interrupts (Ctrl+C, etc.)
// 	quit := make(chan os.Signal, 1)
// 	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

// 	// Block until signal is received
// 	<-quit
// 	log.Println("🔻 Shutdown signal received...")

// 	// Create a context with timeout for shutdown tasks (server, DB)
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()

// 	// Gracefully shutdown HTTP server
// 	if err := srv.Shutdown(ctx); err != nil {
// 		log.Fatalf("❌ Server forced to shutdown: %v", err)
// 	}

// 	// Gracefully disconnect MongoDB
// 	if err := config.Client.Disconnect(ctx); err != nil {
// 		log.Printf("⚠️ MongoDB disconnect error: %v", err)
// 	} else {
// 		log.Println("✅ MongoDB disconnected successfully")
// 	}

// 	log.Println("✅ Server exited cleanly")
// }
