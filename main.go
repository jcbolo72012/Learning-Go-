package main

import(
	"fmt"
	"log"
	"net/http"
)

func main()(

	fileServer = http.fileServer(http.Dir("./static"))

	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at port 8080\n")

)