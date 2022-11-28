package main

import (
	"fmt"
	"log"
	"net/http"



	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main(){
	//TODO　collectorにメトリクスを作成しここで配信をする。
	prometheus.MustRegister(c)
	http.Handle("/metrics", promhttp.Handler())

	log.Println("Listening on ", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal(err)
	}
}