/*
 * example
 *
 * description example.
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package main

import (
	"log"
	"net/http"

	openapi "ghost/go"
)

func main() {
	log.Printf("Server started")

	ServiceApiService := openapi.NewServiceApiService()
	ServiceApiController := openapi.NewServiceApiController(ServiceApiService)

	router := openapi.NewRouter(ServiceApiController)

	log.Fatal(http.ListenAndServe(":8089", router))
}
