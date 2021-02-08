package main

import "zinx/ch02/znet"

func main() {
	server := znet.NewServer("zinx_ch02")
	server.Serve()
}
