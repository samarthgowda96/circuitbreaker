package circuitbreaker

import (
	"fmt"
	"time"

	"github.com/sony/gobreaker"
)

// Create a new CircuitBreaker
func NewCircuitBreaker() *gobreaker.CircuitBreaker {

	cfg := gobreaker.Settings{
		Name: "Default circuit breaker configuration",

		ReadyToTrip: func(counts gobreaker.Counts) bool {
			failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
			return counts.Requests >= 3 && failureRatio >= 0.6
		},
		MaxRequests: 3,
		Timeout:     2 * time.Second,
		Interval:    5 * time.Second,
		OnStateChange: func(name string, from gobreaker.State, to gobreaker.State) {
			fmt.Printf("CircuitBreaker '%s' changed from '%s' to '%s'\n", name, from, to)
		},
	}
	return gobreaker.NewCircuitBreaker(cfg)
}
