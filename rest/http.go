package rest

import (
	"encoding/json"
	"log"
	"net/http"
)

type Request struct {
	Id      int64  `json:"id"`
	Message string `json:"message"`
}

type Response struct {
	Code int
	Data Request
}

func Serve() {
	http.HandleFunc("/", Create)
	log.Println(http.ListenAndServe(":8888", nil))
}

func Create(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	request := Request{}
	decoder.Decode(&request)
	defer r.Body.Close()

	response := Response{
		Code: 200,
		Data: request,
	}
	json.NewEncoder(w).Encode(response)
}
