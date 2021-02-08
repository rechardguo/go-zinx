package main

import "zinx/ch01/znet"

func main() {
	server := znet.NewServer("zinx")
	server.Serve()
}
