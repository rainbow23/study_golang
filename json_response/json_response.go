// Package main provides ...
package main

import (
	"encoding/json"
	// "error"
	"fmt"
	"net/http"
)

type Ping struct {
	Status int    `json:"status"`
	Result string `json:"result"`
}

type Macromill_id struct {
	Id int `json:"macromill_id"`
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	ping := Ping{http.StatusOK, "ok"}
	fmt.Fprintf(w, "ping.Status:%v\n", ping.Status)
	fmt.Fprintf(w, "ping.Result:%v\n", ping.Result)

	mid := Macromill_id{123456}
	fmt.Fprintf(w, "mid :%v\n", mid.Id)
	res, err := json.Marshal(mid)
	fmt.Fprintf(w, "mid json %v\n", string(res))

	// res, err := json.Marshal(ping)
	// fmt.Fprint(w, res.Status)
	// fmt.Fprintf(w, "string(res):%v\n", string(res))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func main() {
	mid := Macromill_id{Id: 12345}
	resMid, err := json.Marshal(mid)
	// fmt.Println("macromill_id:%v\n", mid.id)

	if err != nil {
		fmt.Println("json Marshal fail")
	}

	fmt.Println(string(resMid))

	// http.HandleFunc("/ping", pingHandler)
	// http.ListenAndServe(":8000", nil)
}

// curl localhost:8000/ping
/*
 * ping.Status:200
 * ping.Result:ok
 * string(res):{"status":200,"result":"ok"}
 * {"status":200,"result":"ok"}
 */
