package main

import (
	"demogrpc/server"
	"fmt"
	"os"
	"strconv"
)

func main() {
	PORT_STR := os.Getenv("PORT")
	port, err := strconv.Atoi(PORT_STR)
	if err != nil {
		panic(fmt.Errorf("PORT must be valid integer: %s", err.Error()))
	}
	server.New(port)
}
