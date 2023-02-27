package client

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/lrm25/plGoDistributed/service"
)

func postService(serviceName string) error {

	serviceA := service.NewService(serviceName)
	serviceABytes, err := json.Marshal(serviceA)
	if err != nil {
		return err
	}
	if _, err = http.Post("http://localhost:8080/service", "application/json", bytes.NewBuffer(serviceABytes)); err != nil {
		return err
	}
	return nil
}

func Test() error {

	if err := postService("service A"); err != nil {
		return err
	}
	if err := postService("service B"); err != nil {
		return err
	}
	return nil
}
