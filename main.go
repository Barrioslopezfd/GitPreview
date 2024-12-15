package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main(){

    if len(os.Args) > 2 {
	log.Fatalf("Too many arguments, expected None or 1 -- Received: %d", len(os.Args))
    }

    var PORT string
    var err error
    if len(os.Args) == 2 {
	PORT, err=GetPort(os.Args[1])
	if err != nil {
	    log.Fatal(err)
	}
    } else {
	PORT = "8080"
    }

    mux:=http.NewServeMux()
    mux.Handle("/", http.FileServer(http.Dir(".")))
    server:=&http.Server{
        Handler: mux,
        Addr: fmt.Sprintf(":%s",PORT),
    }
    log.Fatal(server.ListenAndServe())
}