package docker

import (
	"fmt"
	"testing"
)

var api *API

func init() {
	api = NewAPI()
	api.Debug = true
}

/******************
	Image
******************/

func Test_ImageList(t *testing.T) {

	result, err := api.ImageList()
	if err != nil {
		fmt.Printf("Test_NewAPI=> %s\n", err.Error())
		return
	}

	for _, ele := range *result {
		fmt.Printf("%s-%v-%v-%v\n", ele.ID, ele.RepoTags, ele.Size/(1024*1024), ele.Containers)
	}
}

/******************
	Container
******************/

func Test_ContainerList(t *testing.T) {

	result, err := api.ContainerList()
	if err != nil {
		fmt.Printf("Test_ContainerList=> %s\n", err.Error())
		return
	}

	for _, ele := range *result {
		fmt.Printf("%s-%s-%s-%v\n", ele.GetID(), ele.Names, ele.ImageID, ele.Ports)
	}
}

func Test_ContainerGet(t *testing.T) {

	result, err := api.ContainerGet("cbea06bebab4")
	if err != nil {
		fmt.Printf("Test_ContainerList=> %s\n", err.Error())
		return
	}

	fmt.Printf("%v\n", result)
}

func Test_ContainerRun(t *testing.T) {

	result, err := api.ContainerRun("cbea06bebab4", nil)
	if err != nil {
		fmt.Printf("Test_ContainerRun=> %s\n", err.Error())
		return
	}

	fmt.Printf("%v\n", result)
}

func Test_ContainerStop(t *testing.T) {

	err := api.ContainerStop("cbea06bebab4")
	if err != nil {
		fmt.Printf("Test_ContainerStop=> %s\n", err.Error())
		return
	}
}

func Test_ContainerLog(t *testing.T) {

	result, err := api.ContainerLog("cbea06bebab4")
	if err != nil {
		fmt.Printf("ContainerLog=> %s\n", err.Error())
		return
	}
	fmt.Printf(result)
}
