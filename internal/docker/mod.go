package docker

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

func X() {
    apiClient, err := client.NewClientWithOpts(client.FromEnv)
    if err != nil {
        panic(err)
    }
    defer apiClient.Close()
    containers, err := apiClient.ContainerList(context.Background(), container.ListOptions{All: true})
    if err != nil {
        panic(err)
    }
    for _, ctr := range containers {
        fmt.Printf("%s %s (status: %s)\n", ctr.ID, ctr.Image, ctr.Status)
    }
}
