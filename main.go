package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type DefaultResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ResponseFormat struct {
	Name     string
	Github   string
	Linkedin string
}

var data = []ResponseFormat{
	{"Herlianto", "https://github.com/herlianto-github", "https://www.linkedin.com/in/herlianto-%E2%80%8D-829aa284/"},
}

func Welcome(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":

		responses := DefaultResponse{200, "Successful Operation", data}
		res, _ := json.Marshal(responses)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(res)

	default:
		http.Error(w, "Bad Request", 400)
	}
}

func main() {

	http.HandleFunc("/", Welcome)

	fmt.Println("Server Started at localhost:8080")
	http.ListenAndServe(fmt.Sprintf(":%v", 8080), nil)
}
