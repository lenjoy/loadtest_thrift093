package main

import (
	"github.com/rcrowley/go-metrics"
	"log"
	"os"
	"time"
)

func main() {
	g := metrics.NewCounter()
	metrics.Register("bar", g)
	g.Inc(1)

	go metrics.Log(metrics.DefaultRegistry,
		2*time.Second,
		log.New(os.Stdout, "metrics: ", log.Lmicroseconds))

	var j int64
	j = 1
	for true {
		time.Sleep(time.Second * 1)
		g.Inc(1)
		j++
	}
}
