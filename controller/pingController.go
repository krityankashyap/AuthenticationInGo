package controller

import "net/http"

func PingHandler(w http.ResponseWriter, r *http.Request){ // ResponseWriter is a object used to write on ur response
	w.Write([]byte("pong"))  // Write writes the data to the connection as part of an HTTP reply.

}