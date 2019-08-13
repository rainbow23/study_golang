// Package main provides ...
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Ping struct {
	Status int    `json:"status"`
	Result string `json:"result"`
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	ping := Ping{http.StatusOK, "ok"}
	fmt.Fprintf(w, "ping.Status:%v\n", ping.Status)
	fmt.Fprintf(w, "ping.Result:%v\n", ping.Result)

	res, err := json.Marshal(ping)
	// fmt.Fprint(w, res.Status)
	fmt.Fprintf(w, "string(res):%v\n", string(res))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func main() {
	http.HandleFunc("/ping", pingHandler)
	http.ListenAndServe(":8000", nil)
}

// curl localhost:8000/ping
/*
 * ping.Status:200
 * ping.Result:ok
 * string(res):{"status":200,"result":"ok"}
 * {"status":200,"result":"ok"}
 */
