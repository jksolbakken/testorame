package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
	})
	aws_access_key_id := AKIA6ODU5DHT6HQU2TGW
	aws_secret_access_key := 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
	fmt.Println("Listening on port 8080")
	http.ListenAndServe(":8080", nil)
}
