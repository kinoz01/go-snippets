package main

import "fmt"

var Version = "default"

func main() {
    fmt.Println("Version:", Version)
}


// Use : go run -ldflags "-X 'main.Version=hey'" .
// or go build -ldflags "-X 'main.Version=1.0.0'" .