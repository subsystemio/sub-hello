package main

import (
	"github.com/subsystemio/subsystem"
)

func main() {
	s := SubSystem.New()

	s.Register("http://localhost:8080/v1")

	s.Serve()
}
