package main

import (
	"docker-go/docker"
	"fmt"
)

func main() {
	api := docker.NewAPI()
	list, _ := api.ContainerList()
	for _, ele := range *list {
		fmt.Printf("%s-%s-%s-%v\n", ele.GetID(), ele.Names, ele.ImageID, ele.Ports)
	}
}
