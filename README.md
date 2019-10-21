### goGravatar

A gRPC based web service that consumes Gravatar's API to generates URLs mapped to a globally unique email address and personalized avatar. Test coverage only covers the `hasher` package.

* Go
* gRPC
* Protobuf
* Gravatar

To start the application, you would need to enter the downloaded project directory, and then run the `server` and `client` in separate terminals respectively i.e.

* for the `server`, execute `go run server/main.go`
* for the `client`, execute `go run client/main.go`
