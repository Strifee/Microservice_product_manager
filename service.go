package main

import (
	"log"
	"net/http"
	"service/models"
	"service/repository"
)

func main() {
	repository.AddProduct(models.Product{
		Name:      "laptop",
		UnitPrice: 9000.0,
	})
	repository.AddProduct(models.Product{
		Name:      "phone",
		UnitPrice: 3000.0,
	})

	http.HandleFunc("/", handlers.HandleRequest)

	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Println(err)
		return
	}
}
