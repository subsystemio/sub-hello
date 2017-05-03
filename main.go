package main

import (
	"github.com/subsystemio/subsystem"
)

func main() {
	s := SubSystem.New("localhost:8081")

	enc := json.NewEncoder(s.Body.Manager.Connection)
	enc.Encode(Message{Action: "Token", Data: "Test"})
	s.Body.Manager.Connection.Close()
}
