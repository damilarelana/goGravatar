package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"

	mux "github.com/gorilla/mux"
)

// set the Port that the server would be listening/serving from
const port = ":9091"

// integrate the defined GravatarService [from gravatar.proto]
type GravatarService struct {
	// ...
}

// implement Service
func (s *GravatarService) Generate(ctx context.Context, in *pb.GravatarRequest) *pb.GravatarResponse {
	log.Printf("Received email %v with size %v", in.Email, in.Size)
	return &pb.GravatarResponse{
		Url: Gravatar(in.Email, in.Size)
	}
}

// GravatarServerHomePage defines the landing page the server
func ServerHomePage(w http.ResponseWriter, r *http.Request) {
	dataString := "Welcome to the Gravatar Service Server"
	io.WriteString(w, dataString)
}

// custom404PageHandler defines custom 404 page
func custom404PageHandler(w http.ResponseWriter, r *http.Request) {

	// set the content header type
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound) // this automatically generates a 404 status code
	// page content
	data404Page := "This page does not exist ... 404!"
	io.WriteString(w, data404Page)
}

// serviceRequestHandler() defines request handling service [used to aggregate all endpoints before running]
func serviceRequestHandlers() {
	muxRouter := mux.NewRouter().StrictSlash(true)                     // instantiate the gorillamux Router and enforce trailing slash rule i.e. `/path` === `/path/`
	muxRouter.NotFoundHandler = http.HandlerFunc(custom404PageHandler) // customer 404 Page handler scenario
	muxRouter.Handle("/", isAuthorized(jwtServerHomePage))             // uses `http.Handle` because` `IsAuthorized` returns an http.Handler (i.e. an interface NOT a function)
	log.Fatal(http.ListenAndServe(port, muxRouter))                    // set the port where the http server listens and serves. changed `nil` to the instance muxRouter
}

func main() {
	go serviceRequestHandlers() // call and run the server as a goroutine

	// create an artificial pause "to ensure the main function goroutine does not cause the serviceRequestHandler goroutine to exit"
	var tempString string
	fmt.Scanln(&tempString)
}
