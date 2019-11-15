package docker

import (
	"encoding/json"
	"errors"
	"fmt"
)

//ContainerList 获取container列表
func (a *API) ContainerList() (*[]Container, error) {
	//curl --unix-socket /var/run/docker.sock http:/v1.24/containers/json
	pkg, err := a.Get("/containers/json")
	if err != nil {
		return nil, err
	}

	result := make([]Container, 0)
	err = json.Unmarshal([]byte(pkg.Body), &result)
	return &result, err
}

/*ContainerGet 获取一个容器
@id：容器ID
*/
func (a *API) ContainerGet(id string) (*ContainerDetail, error) {
	//curl --unix-socket /var/run/docker.sock  http:/v1.24/containers/4fa6e0f0c678/json
	cmdURL := fmt.Sprintf("/containers/%s/json", id)
	pkg, err := a.Get(cmdURL)
	if err != nil {
		return nil, err
	}

	result := &ContainerDetail{}
	err = json.Unmarshal([]byte(pkg.Body), result)
	return result, err
}

/*ContainerRun 运行容器
@id：容器ID
*/
func (a *API) ContainerRun(image string, cmds []string) (*ContainerDetail, error) {
	/*
		curl --unix-socket /var/run/docker.sock -H "Content-Type: application/json" \
		  -d '{"Image": "alpine", "Cmd": ["echo", "hello world"]}' \
		  -X POST http:/v1.24/containers/create

		{"Id":"1c6594faf5","Warnings":null}
	*/

	cmdURL := fmt.Sprintf(`/containers/create`)
	pkg, err := a.Post(cmdURL)
	if err != nil {
		return nil, err
	}

	result := &ContainerDetail{}
	err = json.Unmarshal([]byte(pkg.Body), result)
	return result, err
}

/*ContainerStop 停止容器
@id：容器ID
*/
func (a *API) ContainerStop(id string) error {
	//curl --unix-socket /var/run/docker.sock  -X POST http:/v1.24/containers/ae63e8b89a26/stop
	cmdURL := fmt.Sprintf("/containers/%s/stop", id)
	pkg, err := a.Post(cmdURL)
	if err != nil {
		return err
	}

	if !pkg.OK {
		return errors.New(pkg.Body)
	}

	return nil
}

/*ContainerLog 打印容器日志
@id：容器ID
*/
func (a *API) ContainerLog(id string) (string, error) {
	//curl --unix-socket /var/run/docker.sock "http:/v1.24/containers/ca5f55cdb/logs?stdout=1"
	cmdURL := fmt.Sprintf("/containers/%s/logs?stdout=1", id)
	pkg, err := a.Get(cmdURL)
	if err != nil {
		return "", err
	}

	if !pkg.OK {
		return "", errors.New(pkg.Body)
	}

	return pkg.Body, nil
}
