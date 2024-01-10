package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type PingHandler struct {
	Message string `json:"message"`
	Status  int    `json:-`
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	ping := &PingHandler{
		Message: `Pong`,
		Status:  http.StatusOK,
	}

	pingBytes, err := json.Marshal(ping)

	if err != nil {
		fmt.Errorf("Internal Server Error....", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(ping.Status)

	w.Write(pingBytes)
}

func main() {

	port := os.Getenv("PORT")
	fmt.Println("PORT IS: ", port)
	http.HandleFunc("/ping", pingHandler)
	http.ListenAndServe(":8080", nil)
}
