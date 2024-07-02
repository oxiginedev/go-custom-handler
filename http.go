package main

import (
	"encoding/json"
	"net/http"
)

type Handler func(w http.ResponseWriter, r *http.Request) error

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h(w, r); err != nil {
		HandleError(w, r, err)
	}
}

func HandlerFunc(hn Handler) http.HandlerFunc {
	return http.HandlerFunc(Handler(hn).ServeHTTP)
}

type HTTPError struct {
	StatusCode int
	Message    string
}

func (he *HTTPError) Error() string {
	return he.Message
}

func HandleError(w http.ResponseWriter, r *http.Request, err error) {
	statusCode := http.StatusInternalServerError

	if sc, ok := err.(*HTTPError); ok {
		statusCode = sc.StatusCode
	}

	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(err.Error()))
}

func RespondJSON(w http.ResponseWriter, statusCode int, d interface{}) error {
	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")

	return json.NewEncoder(w).Encode(d)
}
