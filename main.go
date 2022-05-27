package main

import (
	bot "discord-go-bot/bot/src/main"
  "net/http"
  "fmt"
  // "log"
)

func main() {

  // Running bot
	bot.Start()

  //Create the default mux
	mux := http.NewServeMux()

	//Handling the /v1/teachers. The handler is a function here
	mux.HandleFunc("/v1/teachers", teacherHandler)

	//Handling the /v1/students. The handler is a type implementing the Handler interface here
	sHandler := studentHandler{}
	mux.Handle("/v1/students", sHandler)

	//Create the server. 
	s := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	s.ListenAndServe()

  // // Go server
  // http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
  // 	fmt.Fprintf(w, "Hello, world")
  // })

  // log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))

  

	<-make(chan struct{})
	return
}


func teacherHandler(res http.ResponseWriter, req *http.Request) {
  fmt.Fprintf(res, "Hello, world")
	data := []byte("V1 of teacher's called")
	res.WriteHeader(200)
	res.Write(data)
}

type studentHandler struct{}

func (h studentHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
  fmt.Fprintf(res, "Hello, world")
	data := []byte("V1 of student's called")
	res.WriteHeader(200)
	res.Write(data)
}