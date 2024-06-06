package main

import (
	"fmt"
	"net/http"
)

func main() {
	//awsAccessKey := "AKIA6ODU5DHTQ3ATG57C"
	//awsSecretKey := "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"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
		if err != nil {
			fmt.Printf("oh noes: %v", err)
		}
	})
	fmt.Println("Listening on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(fmt.Sprintf("oh noes: %v", err))
	}
}
