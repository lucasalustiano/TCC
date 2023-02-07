package main

import (
	"flag"
	"math/rand"
	"net/http"
	"time"
)

var numRoutines = flag.Int("n", 0, "Número de requisições a serem feitas para cada endpoint (padrão número randômico)")
var sleepTime = flag.Int("s", 15, "Número em segundos de espera entre as requisições (padrão 15 segundos)")

func main() {
	flag.Parse()

	for {
		if *numRoutines == 0 {
			num := rand.Intn(10) + 1
			for i := 0; i < num; i++ {
				_, err1 := http.Get("http://localhost:8080/func1")
				if err1 != nil {
					println(err1)
				}

				_, err2 := http.Get("http://localhost:8080/func2")
				if err2 != nil {
					println(err2)
				}
			}
		} else {
			for i := 0; i < *numRoutines; i++ {
				_, err1 := http.Get("http://localhost:8080/func1")
				if err1 != nil {
					println(err1)
				}

				_, err2 := http.Get("http://localhost:8080/func2")
				if err2 != nil {
					println(err2)
				}
			}
		}
		time.Sleep(time.Duration(*sleepTime) * time.Second)
	}
}
