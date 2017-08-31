package main

import (
	"net/http"
	"flag"
	"log"
	"time"
	"github.com/Greg-Szabo/tm-peek/restapi"
)

func main() {

	var address = flag.String("address", "127.0.0.1:8000", "Address to listen on. (IP:port)")
	flag.Parse()

	go func() {
		log.Fatal(http.ListenAndServe(*address, restapi.Router()))
	}()


	for true {
		time.Sleep(time.Second*2)
	}


}
