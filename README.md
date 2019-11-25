# docker-go 
```text
docker-go是docker-api的封装，提供基础的函数，用于操作Docker Engine 

优点：简单，看得懂、改的动，能够满足个性化需求

缺点：封装函数较少，封装了常见参数，不是所有的参数
```

----


## Docker 引擎如何与外部api通讯？
```text
dockerd daemon提供了三种不同的socket方式与api通讯，分别是unix,tcp和fd.默认是unix.

unix:
  优点：效率高，数据直接从一个进程copy到另外一个进程
  缺点：只能本机运行

tcp：
  优点：能够远程操作docker
  缺点：接口实现复杂，需要配置https

fs：
  优点：
  缺点：

详细:https://docs.docker.com/engine/reference/commandline/dockerd/
````

---

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
