package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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

func main() {
	http.HandleFunc("/", HelloWorld)
	http.HandleFunc("/user", UserHandle)

	fmt.Println("Menjalankan program ....")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal(err.Error())
	}
}
