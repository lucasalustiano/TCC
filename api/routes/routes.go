package routes

import (
	"log"
	"net/http"

	"github.com/lucasalustiano/TCC/api/functions"
)

func HandleResquests() {
  http.HandleFunc("/func1", functions.Func1)
	http.HandleFunc("/func2", functions.Func2)
  
	log.Fatal(http.ListenAndServe(":8000", nil))
}