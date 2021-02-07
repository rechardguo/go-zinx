package main

import "zinx/ch02/znet"

func main() {
	server := znet.NewServer("zinx")
	server.Serve()
}
