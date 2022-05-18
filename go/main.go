package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/weather", weatherEndpoint)
	http.HandleFunc("/health", getHealthEndpoint)

	fmt.Println("Will serve...")

	http.ListenAndServe(":8899", nil)
}
