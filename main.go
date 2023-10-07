package main

import (
	"fmt"
	"net/http"

	"github.com/fanialfi/golang-web-service-api-server/routing"
)

var port string = ":8080"

func main() {
	http.HandleFunc("/users", routing.Users)
	http.HandleFunc("/user", routing.User)

	fmt.Printf("web service running on localhost%s\n", port)
	http.ListenAndServe(port, nil)
}
