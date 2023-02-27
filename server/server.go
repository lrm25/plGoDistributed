package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"

	"github.com/lrm25/plGoDistributed/service"
)

type Server struct {
	portNum      int
	serviceMutex sync.RWMutex
	services     map[string]bool
}

type Register struct {
	ServiceName string
}

type LogMessage struct {
	Sent        *time.Time
	Received    *time.Time
	ServiceName string
	Message     string
}

func NewServer(startingPort int) *Server {
	return &Server{
		portNum:  startingPort,
		services: make(map[string]bool),
	}
}

func (s *Server) AddService(name string) {
	s.serviceMutex.Lock()
	s.services[name] = true
	s.serviceMutex.Unlock()
}

func (s *Server) GetService(name string) bool {
	s.serviceMutex.RLock()
	defer s.serviceMutex.RUnlock()
	if _, ok := s.services[name]; ok {
		return true
	}
	return false
}

func (s *Server) PrintServices(w http.ResponseWriter) {
	s.serviceMutex.RLock()
	defer s.serviceMutex.RUnlock()

	w.Write([]byte("<html><body>"))
	for name, _ := range s.services {
		w.Write([]byte(name + "<br>"))
	}
	w.Write([]byte("</html></body>"))
}

func (s *Server) Start() error {

	http.HandleFunc("/log", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Log message received")
		fmt.Println("received")
		switch r.Method {
		case "POST":
			fmt.Println("Post method")
			bytes, _ := ioutil.ReadAll(r.Body)
			fmt.Println(string(bytes))
		default:
			fmt.Println("Not supported")
		}
	})

	http.HandleFunc("/service", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			defer r.Body.Close()
			bodyBytes, err := ioutil.ReadAll(r.Body)
			if err != nil {
				w.Write([]byte("Error reading message body: " + err.Error()))
				return
			}
			var serviceData service.Service
			if err := json.Unmarshal(bodyBytes, &serviceData); err != nil {
				w.Write([]byte("Error unmarshaling JSON: " + err.Error()))
				return
			}
			s.AddService(serviceData.Name)
		case "GET":
			s.PrintServices(w)
		default:
			fmt.Println("Not supported")
		}
	})

	attempts := 0
	for ; attempts < 7; attempts++ {
		listenPort := fmt.Sprintf(":%d", s.portNum)
		fmt.Printf("Attempting to start server on port %d\n", s.portNum)
		if err := http.ListenAndServe(listenPort, nil); err != nil {
			fmt.Println(err.Error())
			s.portNum++
		}
	}
	if attempts == 7 {
		return errors.New("Unable to start server, exiting program")
	}
	return nil
}
