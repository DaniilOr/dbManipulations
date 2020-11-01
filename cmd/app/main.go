package main

import (
	"github.com/DaniilOr/webGo/cmd/app/Server"
	"github.com/DaniilOr/webGo/src/service"
	"log"
	"net"
	"net/http"
	"os"
)
const defaultPort = "8888"
const defaultHost = "0.0.0.0"

func main() {
	os.Setenv("PORT", defaultPort)
	os.Setenv("HOST", defaultHost)
	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = defaultPort
	}

	host, ok := os.LookupEnv("HOST")
	if !ok {
		host = defaultHost
	}

	log.Println(host)
	log.Println(port)

	if err := execute(net.JoinHostPort(host, port)); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
func execute(addr string)(err error) {
	service, err := service.CreateNewService()
	if err != nil{
		log.Println(err)
		return err
	}
	mux := http.NewServeMux()
	application := Server.NewServer(service, mux)
	application.Init()
	server := &http.Server{
		Addr: addr,
		Handler: application,
	}
	return server.ListenAndServe()
}
