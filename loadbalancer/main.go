package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync"
)

type Backend struct {
	URL         *url.URL
	Active      bool
	Connections int
}

var backends = []*Backend{
	{URL: mustParseURL("http://backend-1:8082"), Active: true},
	// {URL: mustParseURL("http://backend-2:8083"), Active: true},
	// {URL: mustParseURL("http://backend-3:8084"), Active: true},
}

var mu sync.Mutex
var roundRobinIndex int

func mustParseURL(rawURL string) *url.URL {
	u, _ := url.Parse(rawURL)
	return u
}

func loadBalancerHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	backend := backends[roundRobinIndex]
	roundRobinIndex = (roundRobinIndex + 1) % len(backends)
	mu.Unlock()

	proxy := httputil.NewSingleHostReverseProxy(backend.URL)
	proxy.ServeHTTP(w, r)
}

func main() {
	http.HandleFunc("/", loadBalancerHandler)
	log.Println("Load Balancer démarré sur :8080")
	http.ListenAndServe(":8080", nil)
}
