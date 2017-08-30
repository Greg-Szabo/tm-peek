package main

import (
	"net/http"
	"flag"
	"log"
	"fmt"
	"time"
)

func main() {
	var address = flag.String("address", "127.0.0.1:8000", "Address to listen on. (IP:port)")
	flag.Parse()

	router := NewRouter()

//	go func() {
		log.Fatal(http.ListenAndServe(*address, router))
//	}()


	for true {
		fmt.Println(dataStore)
		time.Sleep(time.Second*2)
	}


}
