package main

import (
	"itvdn/nethttp"
	"time"
	"itvdn/mux"
	"itvdn/gingonic"
)

func main(){
	go nethttp.Run()
	go mux.Run()
	go gingonic.Run()

	for {
		time.Sleep(5 * time.Second)
	}

}
