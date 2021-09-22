package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
	"github.com/mohammadne/middleman/cmd/client"
	loadbalancer "github.com/mohammadne/middleman/cmd/load_balancer"
	"github.com/mohammadne/middleman/cmd/server"
	"github.com/mohammadne/middleman/pkg/file"
)

const rootDir = "./static"

func main() {
	err := file.CreateDirIfMissed(rootDir)
	if err != nil {
		log.Fatal("Error creating root directory")
	}

	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	host := os.Getenv("HOST")

	// SETUP SERVERS
	servers := strings.Split(os.Getenv("SERVER_PORTS"), ",")
	for index, port := range servers {
		servers[index] = fmt.Sprintf("%s:%s", host, port)
		go server.Setup(
			fmt.Sprintf("%s/%s", rootDir, port),
			servers[index],
		)
	}

	// SETUP LOAD_BALANCER
	loadbalancerPort := os.Getenv("LOAD_BALANCER_PORT")
	go loadbalancer.Setup(
		fmt.Sprintf("%s:%s", host, loadbalancerPort),
		servers,
	)

	// SETUP CLIENT
	requestsNum, _ := strconv.Atoi(os.Getenv("CLIENT_REQUESTS"))
	go client.Setup(
		fmt.Sprintf("http://%s:%s/objects", host, loadbalancerPort),
		requestsNum,
	)

	blockingChannel := make(chan int)
	<-blockingChannel
}
