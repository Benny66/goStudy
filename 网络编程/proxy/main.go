package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync"
	"time"
)

// Server represents a backend server and its associated weight.
type Server struct {
	URL    *url.URL
	Weight float64
}

// LoadBalancer represents a dynamic weighted round-robin load balancer.
type LoadBalancer struct {
	servers []*Server
	mutex   sync.RWMutex
}

// NewLoadBalancer returns a new LoadBalancer instance.
func NewLoadBalancer() *LoadBalancer {
	return &LoadBalancer{
		servers: make([]*Server, 0),
	}
}

// AddServer adds a new server with the given URL and weight to the load balancer.
func (lb *LoadBalancer) AddServer(serverURL *url.URL, weight float64) {
	lb.mutex.Lock()
	defer lb.mutex.Unlock()

	server := &Server{
		URL:    serverURL,
		Weight: weight,
	}

	lb.servers = append(lb.servers, server)
}

// RemoveServer removes the server at the given index from the load balancer.
func (lb *LoadBalancer) RemoveServer(index int) {
	lb.mutex.Lock()
	defer lb.mutex.Unlock()

	if index < 0 || index >= len(lb.servers) {
		return
	}

	lb.servers = append(lb.servers[:index], lb.servers[index+1:]...)
}

// NextServer returns the next server to handle a request, selected using a dynamic
// weighted round-robin algorithm.
func (lb *LoadBalancer) NextServer() (*Server, error) {
	lb.mutex.Lock()
	defer lb.mutex.Unlock()

	if len(lb.servers) == 0 {
		return nil, fmt.Errorf("no servers available")
	}

	totalWeight := 0.0
	for _, server := range lb.servers {
		totalWeight += server.Weight
	}

	rand.Seed(time.Now().UnixNano())
	rnd := rand.Float64() * totalWeight

	for i := range lb.servers {
		server := lb.servers[i]
		rnd -= server.Weight
		if rnd <= 0 {
			return server, nil
		}
	}

	// Should never happen.
	return nil, fmt.Errorf("unable to select a server")
}

// ReverseProxy is a custom implementation of httputil.ReverseProxy that uses
// the dynamic weighted round-robin load balancer to select a backend server.
type ReverseProxy struct {
	loadBalancer *LoadBalancer
}

// NewReverseProxy returns a new ReverseProxy instance with the given load balancer.
func NewReverseProxy(loadBalancer *LoadBalancer) *ReverseProxy {
	return &ReverseProxy{
		loadBalancer: loadBalancer,
	}
}

// ServeHTTP handles an incoming HTTP request by selecting a backend server using the
// dynamic weighted round-robin algorithm and forwarding the request to it.
func (p *ReverseProxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	server, err := p.loadBalancer.NextServer()
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	proxy := httputil.NewSingleHostReverseProxy(server.URL)
	proxy.ServeHTTP(w, r)
}

func main() {
	// Initialize the load balancer with some backend servers.
	lb := NewLoadBalancer()
	lb.AddServer(&url.URL{
		Scheme: "https",
		Host:   "admin.janime.cn/",
	}, 3.0)
	lb.AddServer(&url.URL{
		Scheme: "https",
		Host:   "www.baidu.com",
	}, 1.0)
	lb.AddServer(&url.URL{
		Scheme: "https",
		Host:   "www.qq.com/",
	}, 2.0)

	// Create a new HTTP server that uses the reverse proxy to handle requests.
	proxy := NewReverseProxy(lb)
	server := &http.Server{
		Addr:    ":8080",
		Handler: proxy,
	}

	// Start the HTTP server.
	log.Printf("Starting server on %s...\n", server.Addr)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
