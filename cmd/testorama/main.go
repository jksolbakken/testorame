package main

import (
	"fmt"
	"net/http"
)

func main() {
	awsAccessKey := "AKIA6ODU5DHT2EQ2SXNW"
	awsSecretKey := "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"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
	})
	fmt.Println("Listening on port 8080")
	http.ListenAndServe(":8080", nil)
}
