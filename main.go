package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync"
)

// LoadBalancer represents a simple round-robin load balancer
type LoadBalancer struct {
	servers []string
	mu      sync.Mutex
	index   int
}

// NewLoadBalancer creates a new LoadBalancer
func NewLoadBalancer(servers []string) *LoadBalancer {
	return &LoadBalancer{
		servers: servers,
		index:   0,
	}
}

// GetNextServer returns the next server in round-robin order
func (lb *LoadBalancer) GetNextServer() string {
	lb.mu.Lock()
	defer lb.mu.Unlock()
	server := lb.servers[lb.index]
	lb.index = (lb.index + 1) % len(lb.servers)
	return server
}

// ServeHTTP handles HTTP requests and forwards them to the selected server
func (lb *LoadBalancer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	target := lb.GetNextServer()
	fmt.Printf("Forwarding request to: %s\n", target)
	url, _ := url.Parse("http://" + target)
	proxy := httputil.NewSingleHostReverseProxy(url)
	proxy.ServeHTTP(w, r)
}

func main() {
	servers := []string{"localhost:8081", "localhost:8082", "localhost:8083"}
	lb := NewLoadBalancer(servers)

	http.HandleFunc("/", lb.ServeHTTP)
	fmt.Println("Load balancer started on :8080")
	http.ListenAndServe(":8080", nil)
}
