package main

import "github.com/sdttttt/go-sso/entry"

import "net/rpc"

import "net"

import "net/http"

func main() {
	authModule := new(entry.AuthenticationModule)

	err := rpc.Register(authModule)

	if err != nil {
		panic(err.Error())
	}

	rpc.HandleHTTP()

	listen, err := net.Listen("tcp", ":8001")

	if err != nil {
		panic(err.Error())
	}

	http.Serve(listen, nil)

}
