package server

import (
	"fmt"
	"gorouter/types"
	"io"
	"net/http"
)

type HttpServer struct  {
	host string
}

func NewHTTPServer(hostname string) *HttpServer {
	return &HttpServer{host : hostname}
}

func (this *HttpServer)Run() {
	http.HandleFunc("/foo", fooHandler)
	http.HandleFunc("/", rootHandler)
	go http.ListenAndServe(this.host, nil)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	rw := types.NewRequestWriter()
	var writer io.Writer = rw
	r.Write(writer)
	output := fmt.Sprintf("Root Handler \n\nRequest = %v  \n\n", rw.Data())

	w.Write([]byte(output))
}

//自定义
func fooHandler(w http.ResponseWriter, r *http.Request) {
	json := `{ "result_code":0,` +
			`"app_state":{ ` +
			`"mysql:":"normal", ` + 	
			`"router:":"normal", ` + 
			`"router_conn:":4, ` + 
			`"nginx:":"normal" ` +  `}}`
	w.Write([]byte(json))
}