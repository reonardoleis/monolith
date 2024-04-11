package main

import "github.com/reonardoleis/views/internal/server"

func main() {
	server := server.NewServer()
	if err := server.Run(":3000"); err != nil {
		panic(err)
	}
}
