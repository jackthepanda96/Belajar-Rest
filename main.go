package main

import (
	"fmt"
	"log"
	"net/http"
)

func UserHandle1(writer http.ResponseWriter, reader *http.Request) {
	switch reader.Method {
	case "GET":
		writer.Header().Set("content-type", "application/json")
		msg := "Hello World"
		writer.Write([]byte(msg))
	}
}

func main() {
	http.HandleFunc("/user", UserHandle1)

	fmt.Println("Menjalankan program ....")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal(err.Error())
	}
}
