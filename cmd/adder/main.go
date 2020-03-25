package main

import (
	"fmt"
	"github.com/Dr-Evans/adder-service/internal/adderserver"
	"github.com/Dr-Evans/adder-service/rpc/adder"
	"net/http"
)

func main() {
	server := &adderserver.Server{}

	twirpHandler := adder.NewAdderServer(server, nil)

	fmt.Println("Service is running on port :8080")
	err := http.ListenAndServe(":8080", twirpHandler)
	if err != nil {
		panic(err)
	}
}
