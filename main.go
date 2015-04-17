package main

import (
	"log"
	"os"

	"github.com/cargonauts/go-dockerclient"
	"github.com/codegangsta/cli"
)

func getClient(endpoint string) *docker.Client {
	client, _ := docker.NewClient(endpoint)
	return client
}

func getContainers(client *docker.Client) []docker.APIContainers {
	containers, _ := client.ListContainers(docker.ListContainersOptions{All: false})
	return containers
}

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

		client := getClient(endpoint)
		containers := getContainers(client)

		for _, container := range containers {
			log.Printf("ID: %s", container.ID)
			chan c := client.StatsContainer(container.ID)
		}
	}
	app.Run(os.Args)
}
