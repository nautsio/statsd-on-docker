package main

import (
	"fmt"
	"log"
	"os"

	"github.com/codegangsta/cli"
	"github.com/fsouza/go-dockerclient"
)

func main() {
	app := cli.NewApp()
	app.Name = "docker-stats"
	app.Usage = "Docker Stats to Statsd"
	app.Version = "0.0.1"
	app.Author = "Cargonauts"
	app.Email = "info@cargonauts.io"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "endpoint, e",
			Value: "tcp://127.0.0.1:2375",
			Usage: "endpoint to the Docker API",
		},
	}
	app.Action = func(c *cli.Context) {
		endpoint := c.String("endpoint")

		log.Printf("Starting %s", app.Usage)
		log.Printf("Connecting to '%s'", endpoint)

		client, _ := docker.NewClient(endpoint)
		containers, _ := client.ListContainers(docker.ListContainersOptions{All: false})

		for _, container := range containers {
			fmt.Println("ID: ", container.ID)
			fmt.Println("Names: ", container.Names)
		}
	}
	app.Run(os.Args)
}
