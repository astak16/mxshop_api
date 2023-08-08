package main

import (
	"log"

	_ "github.com/mbobakov/grpc-consul-resolver"

	"google.golang.org/grpc"
)

func main2() {
	conn, err := grpc.Dial("consul:172.18.0.3:8500/user-srv?wait=14s&tag=srv",
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

}
