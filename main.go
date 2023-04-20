package main

import (
	"circuitbreaker/vinservice"
	"errors"
	"fmt"
	"net/http"
	"time"

	"circuitbreaker/circuitbreaker"

	"github.com/gin-gonic/gin"
)

var startTime time.Time = time.Now()

func server() {
	router := gin.Default()
	router.GET("/v1/vin-details/:id", func(ctx *gin.Context) {
		if time.Since(startTime) < 10*time.Second {
			ctx.String(http.StatusInternalServerError, "Server not ready yet")
			fmt.Println("Server not ready yet")
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

func DoReq() error {
	resp, err := http.Get("http://localhost:8080/v1/vin-details/5J6RM4H50GL105806")
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return errors.New("bad response")
	}

	return nil
}

func main() {
	go server()
	cb := circuitbreaker.NewCircuitBreaker()

	fmt.Println("Calling the circuit breaker")
	for i := 0; i < 100; i++ {
		_, err := cb.Execute(func() (interface{}, error) {
			err := DoReq()
			return nil, err
		})
		if err != nil {
			fmt.Println(err)
		}
		time.Sleep(1500 * time.Millisecond)

	}
}
