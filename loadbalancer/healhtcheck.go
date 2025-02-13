package main

import (
	"net/http"
	"time"
)

func healthCheck() {
	for {
		for _, backend := range backends {
			resp, err := http.Get(backend.URL.String() + "/health")
			backend.Active = err == nil && resp.StatusCode == http.StatusOK
		}
		time.Sleep(5 * time.Second)
	}
}

func init() {
	go healthCheck()
}
