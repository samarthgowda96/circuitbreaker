package main

import (
	"fmt"
	"net/http"
	"time"

	"circuitbreaker/vinservice"

	"github.com/gin-gonic/gin"
)

var startTime time.Time = time.Now()

func main() {
	router := gin.Default()
	router.GET("/v1", func(ctx *gin.Context) {
		ctx.String(http.StatusInternalServerError, "Server is healthy")

	})

	router.GET("/v1/vin-details/:id", func(ctx *gin.Context) {
		if time.Since(startTime) < 5*time.Second {
			ctx.String(http.StatusInternalServerError, "Server not ready yet")
			return
		}
		id := ctx.Param("id")
		body, err := vinservice.NewVinService(id)
		if err != nil {
			ctx.String(http.StatusInternalServerError, "Failed to fetch VIN data")
			return
		}

		ctx.Data(http.StatusOK, "application/json", body)
	})

	fmt.Printf("Starting server at port 8080\n")
	router.Run(":8080")
}
