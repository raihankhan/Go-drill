package hello_webserver

import (
	"fmt"
	"net/http"
)

func helloServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
		if err != nil {
			return
		}
	})

	err := http.ListenAndServe(":80", nil)
	if err != nil {
		return
	}
}
