package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"time"
)

func main(){
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/tl-hasher/sha256", processRequest).Methods("POST")
	log.Println("Listening on port 8080")
	log.Println("Application Started")
	log.Fatal(http.ListenAndServe(":8080", myRouter))

}

type Response struct {
	Hash string `json:"hash"`
	TimeTakenMS int64 `json:"time_taken_ms"`
}

func processRequest(writer http.ResponseWriter, request *http.Request) {
	startTime:=time.Now().UnixNano() / int64(time.Millisecond)
	defer request.Body.Close()
	hash:=sha256.New()
	file,_,_:=request.FormFile("file")
	if _, err := io.Copy(hash, file); err != nil {
		log.Fatal(err)
	}
	sum := hash.Sum(nil)
	finishTime:=time.Now().UnixNano() / int64(time.Millisecond)
	response:= Response{
		Hash:        fmt.Sprintf("%x", sum),
		TimeTakenMS: finishTime - startTime,
	}

	writer.Header().Add("Content-type", "application/json")
	json,_:=json.Marshal(response)
	writer.Write(json)
}
