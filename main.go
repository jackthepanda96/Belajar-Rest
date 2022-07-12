package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type User struct {
	Nama     string
	Email    string
	Password string
}

var (
	arrData []User
)

func HelloWorld(writer http.ResponseWriter, reader *http.Request) {
	switch reader.Method {
	case "GET":
		writer.Header().Set("content-type", "application/json")
		msg := "Hello World"
		writer.Write([]byte(msg))
	}
}

func UserHandle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	switch r.Method {
	case "POST":
		var tmp User
		decode := json.NewDecoder(r.Body)
		err := decode.Decode(&tmp)
		if err != nil {
			log.Println("Cannot parse", err.Error())
		}
		log.Println(tmp)
		arrData = append(arrData, tmp)
		res := map[string]interface{}{
			"message": "Success input data",
			"data":    tmp,
		}
		send, err := json.Marshal(res)
		if err != nil {
			log.Println("Cannot send", err.Error())
		}
		w.Write(send)
	case "GET":
		res := map[string]interface{}{
			"message": "Get all data",
			"data":    arrData,
		}
		send, err := json.Marshal(res)
		if err != nil {
			log.Println("Cannot send", err.Error())
		}
		w.Write(send)
	}
}

func UserSpecificHandle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	res := map[string]interface{}{}
	params := mux.Vars(r)
	cnv, err := strconv.Atoi(params["id"])

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		log.Println("Cannot convert to int", err.Error())
		return
	}

	if cnv > len(arrData) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		log.Println("Index out of range")
		return
	}

	switch r.Method {
	case "GET":
		res = map[string]interface{}{
			"message": "Get all data",
			"data":    arrData[cnv-1],
		}
	case "DELETE":
		arrData = arrData[1:]
		res = map[string]interface{}{
			"message": "success delete first data",
		}
	}
	send, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		log.Println("Cannot send", err.Error())
	}
	w.Write(send)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HelloWorld)
	r.HandleFunc("/user", UserHandle).Methods("POST", "GET")
	r.HandleFunc("/user/{id}", UserSpecificHandle).Methods("GET", "PUT", "DELETE")

	fmt.Println("Menjalankan program ....")
	err := http.ListenAndServe(":8000", r)
	if err != nil {
		log.Fatal(err.Error())
	}
}
