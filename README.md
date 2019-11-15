# Docker API 

docker API client，用于操作docker
```text
使用unix方式与docker daemon通讯
详细:https://docs.docker.com/engine/reference/commandline/dockerd/
````


example
```golang
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
```