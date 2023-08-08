package main

import (
	"fmt"
	"strconv"

	"github.com/hashicorp/consul/api"
)

func Register(address string, port int, name string, tags []string, id string) {
	cfg := api.DefaultConfig()
	cfg.Address = "go-consul" + ":8500"
	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	registration := new(api.AgentServiceRegistration)
	registration.Name = name
	registration.ID = id
	registration.Port = port
	registration.Tags = tags
	registration.Address = address
	fmt.Println(strconv.Itoa(port), port)
	fmt.Println("http://" + address + ":" + strconv.Itoa(port) + "/health")
	check := &api.AgentServiceCheck{
		CheckID: id,
		// HTTP:    "http://localhost:8021/health",
		HTTP:                           "http://" + address + ":" + strconv.Itoa(port) + "/health",
		Timeout:                        "5s",
		Interval:                       "5s",
		DeregisterCriticalServiceAfter: "400s",
	}
	registration.Check = check

	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		panic(err)
	}
}

func AllServices() {
	cfg := api.DefaultConfig()
	cfg.Address = "go-consul" + ":8500"
	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	data, err := client.Agent().Services()
	if err != nil {
		panic(err)
	}

	for key, _ := range data {
		fmt.Println(key)
	}
}

func FilterService() {
	cfg := api.DefaultConfig()
	cfg.Address = "go-consul" + ":8500"
	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	// data, err := client.NewClient()
	// if err != nil {
	// 	panic(err)
	// }

	data, err := client.Agent().ServicesWithFilter(`Service == "mxshop-web1"`)

	if err != nil {
		panic(err)
	}

	for key, _ := range data {
		fmt.Println(key)
	}
}

func main1() {
	// FilterService()
	// AllServices()
	Register("172.18.0.3", 8021, "user_web", []string{"mxshop", "uccs"}, "user_web")
}
