package docker

import "encoding/json"

//Image 镜像
type Image struct {
	Containers  int         `json:"Containers"`
	Created     int         `json:"Created"`
	ID          string      `json:"Id"`
	Labels      interface{} `json:"Labels"`
	ParentID    string      `json:"ParentId"`
	RepoDigests []string    `json:"RepoDigests"`
	RepoTags    []string    `json:"RepoTags"`
	SharedSize  int         `json:"SharedSize"`
	Size        int         `json:"Size"`
	VirtualSize int         `json:"VirtualSize"`
}

//GetID 容器ID
func (c *Image) GetID() string {
	return c.ID[0:12]
}

//ImageList 获取image列表
func (a *API) ImageList() (*[]Image, error) {
	//curl --unix-socket /var/run/docker.sock http:/v1.24/images/json
	pkg, err := a.Get("/images/json")
	if err != nil {
		return nil, err
	}

	result := make([]Image, 0)
	err = json.Unmarshal([]byte(pkg.Body), &result)
	return &result, err
}
