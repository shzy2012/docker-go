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
结果
```text
cbea06bebab4-[/shzy-redis]-a55fbf438dfd-[{0.0.0.0 6379 6379 tcp}]
16fb7fcead21-[/shzy-mysql]-7bb2586065cd-[{0.0.0.0 3306 3306 tcp} { 33060 0 tcp}]
```
