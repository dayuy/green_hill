package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"time"
)

func main1(){
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!")
	})
	http.HandleFunc("/time/", func(w http.ResponseWriter, r *http.Request) {
		t:=time.Now()
		timeStr:=fmt.Sprintf("{\"time\":\"%s\"}", t)
		w.Write([]byte(timeStr))
	})
	http.ListenAndServe(":8080", nil)
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params)  {
	fmt.Fprint(w, "Welcome!\n")
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params )  {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func main2() {
	router:=httprouter.New()
	router.GET("/", Index)
	router.GET("/hello/:name", Hello)

	log.Fatal(http.ListenAndServe(":8080", router))
}