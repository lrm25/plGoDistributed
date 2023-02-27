package main

import (
	"fmt"

	"github.com/lrm25/plGoDistributed/client"
	"github.com/lrm25/plGoDistributed/server"
)

func main() {
	server := server.NewServer(8080)
	done := make(chan bool)
	go func() {
		server.Start()
		done <- true
	}()
	if err := client.Test(); err != nil {
		fmt.Println(err.Error())
	}
	<-done
}
