package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/na50r/server/sse"
)

type apiFunc func(w http.ResponseWriter, r *http.Request) error

type ApiError struct {
	Error string `json:"error"`
}

type APIServer struct {
	listenAddr string
	broker     *sse.Broker
}

func NewAPIServer(listenAddr string) *APIServer {
	newBroker := sse.NewServer()

	return &APIServer{
		listenAddr: listenAddr,
		broker:     newBroker,
	}
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request: %s %s", r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func (s *APIServer) Run() {
	router := newRouter()
	router.HandleFunc("/wordcount", makeHTTPHandleFunc(s.handleWordCount))
	router.HandleFunc("/cached", makeHTTPHandleFunc(s.handleWordCountCached))
	router.HandleFunc("/set1", makeHTTPHandleFunc(s.handleSet1))

	router.HandleFunc("/messages", s.broker.BroadcastMessage).Methods("POST")
	router.HandleFunc("/stream", s.broker.Stream).Methods("GET")

	log.Println("Starting server on:", s.listenAddr)
	http.ListenAndServe(s.listenAddr, router)

}

func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.Use(loggingMiddleware) // Add the logging middleware
	return r
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}

func (s *APIServer) handleWordCount(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodPost {
		return fmt.Errorf("invalid method %s", r.Method)
	}

	var req WordCountReqest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return err
	}

	resp := WordCountResponse{
		Count: len(req.Word),
	}

	return WriteJSON(w, http.StatusOK, resp)
}

func (s *APIServer) handleWordCountCached(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodPost {
		return fmt.Errorf("invalid method %s", r.Method)
	}

	var req WordCountReqest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return err
	}

	count, err := checkWordWithDB(s.broker, req.Word)
	if err != nil {
		return err
	}

	resp := WordCountResponse{
		Count: count,
	}

	return WriteJSON(w, http.StatusOK, resp)
}

func (s *APIServer) handleSet1(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodPost {
		return fmt.Errorf("invalid method %s", r.Method)
	}

	InitSet1()

	return WriteJSON(w, http.StatusOK, "Set1 initialized")
}
