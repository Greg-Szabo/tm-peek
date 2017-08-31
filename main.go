package main

import (
	"github.com/Greg-Szabo/tm-peek/restapi"
	"flag"
)

func main() {

	var address = flag.String("address", "127.0.0.1:8000", "Address to listen on. (IP:port)")
	flag.Parse()

	restapi.Start(*address,"localhost:46657", false)

}
