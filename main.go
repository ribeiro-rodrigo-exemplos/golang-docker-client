package main

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func main() {
	cli, err := client.NewEnvClient()

	if err != nil {
		fmt.Println(err)
	}

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{
		All: true,
	})

	if err != nil {
		panic(err)
	}

	for _, container := range containers {
		fmt.Println(container, container.Image)
	}
}
