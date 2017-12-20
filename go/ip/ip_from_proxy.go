package main

import (
	"net/http"
)

func getIPFromProxy(resp http.ResponseWriter, req *http.Request) {
	ip := req.Header.Get("X-real-ip")
	if ip == "" {
		ip = "ip为空"
	}
	resp.Write([]byte(ip))
}

func startServer() {
	http.HandleFunc("/", getIPFromProxy)
	http.ListenAndServe(":8564", nil)
}

func main() {
	startServer()
}
