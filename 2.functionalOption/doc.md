Functional Options is a design pattern in Go that provides a flexible way to set options for struct configuration, primarily during its initialization. This pattern involves creating functions that modify a struct's fields, and then applying these functions to the struct when it's being instantiated. This pattern enhances readability and maintainability, especially when dealing with a lot of optional parameters.

## Server Example

```go
package main

import "fmt"

// Server struct with optional configuration fields
type Server struct {
    Port    int
    Timeout int
}

// ServerOption represents a function that applies a configuration to the Server
type ServerOption func(*Server)

// WithPort returns a ServerOption that sets the port
func WithPort(port int) ServerOption {
    return func(s *Server) {
        s.Port = port
    }
}

// WithTimeout returns a ServerOption that sets the timeout
func WithTimeout(timeout int) ServerOption {
    return func(s *Server) {
        s.Timeout = timeout
    }
}

// NewServer creates a new Server with the provided options
func NewServer(opts ...ServerOption) *Server {
    server := &Server{
        Port:    8080,  // default port
        Timeout: 30,    // default timeout in seconds
    }

    // Apply each option to the server
    for _, opt := range opts {
        opt(server)
    }

    return server
}

func main() {
    // Create a server with default settings
    defaultServer := NewServer()
    fmt.Printf("Default Server: %+v\n", defaultServer)

    // Create a server with a custom port and timeout
    customServer := NewServer(WithPort(9090), WithTimeout(60))
    fmt.Printf("Custom Server: %+v\n", customServer)
}
```
