package main

import (
	"github.com/kittichok/go-auth/internal/pkg"
	"github.com/kittichok/go-auth/server"
)


func main() {
	//TODO: use inject instead
	pkg.InitTrace()
	pkg.SetupTransacName("auth")
	
	server.Init()
}

