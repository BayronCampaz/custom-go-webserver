package webserver

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func HandleRoot(wt http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(wt, "Hello World server")
}

func HandleHome(wt http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(wt, "This is home page")
}

func PostRequest(wt http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var metadata MetaData
	err := decoder.Decode(&metadata)
	if err != nil {
		fmt.Fprintf(wt, "error: %v", err)
		return
	}
	fmt.Fprintf(wt, "Payload : %v\n", metadata)
}

func UserPostRequest(wt http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var user User
	err := decoder.Decode(&user)
	if err != nil {
		fmt.Fprintf(wt, "error: %v", err)
		return
	}
	response, err := user.ToJson()
	if err != nil {
		wt.WriteHeader(http.StatusInternalServerError)
		return
	}
	wt.Header().Set("Content-Type", "application/json")
	wt.Write(response)
}
