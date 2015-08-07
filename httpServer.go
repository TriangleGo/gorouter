package main

import (
	"fmt"
	"gorouter/types"
	"io"
	"net/http"
)

func HTTPServer() {
	//http.HandleFunc("/foo", fooHandler)
	http.HandleFunc("/", rootHandler)
	http.ListenAndServe(":8080", nil)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	rw := types.NewRequestWriter()
	var writer io.Writer = rw
	r.Write(writer)
	output := fmt.Sprintf("Root Handler \n\nRequest = %v  \n\n", rw.Data())

	w.Write([]byte(output))
}
