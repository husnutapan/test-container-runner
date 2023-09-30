package internal

import (
	"context"
	"fmt"
	"github.com/testcontainers/testcontainers-go"
)

type Container interface {
	Start()
}

type container struct {
	container    testcontainers.Container
	containerUrl string
	exposePorts  []string
}

func NewContainer(containerUrl string, exposePorts []string) *container {
	var ports []string
	for i := 0; i < len(exposePorts); i++ {
		port := fmt.Sprintf("%s:%s", exposePorts[i], exposePorts[i])
		ports = append(ports, port)
	}
	return &container{
		containerUrl: containerUrl,
		exposePorts:  ports,
	}
}

func (c *container) Start() {
	request := testcontainers.ContainerRequest{
		Image:        c.containerUrl,
		ExposedPorts: c.exposePorts,
	}

	c.container, _ = testcontainers.GenericContainer(context.Background(), testcontainers.GenericContainerRequest{
		ContainerRequest: request,
		Started:          true,
	})
}
