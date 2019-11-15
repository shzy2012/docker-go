package docker

//PKG 数据包
type PKG struct {
	OK                 bool
	Code               string
	Protocol           string
	APIVersion         string
	ContentLength      string
	ContentType        string
	Date               string
	DockerExperimental string
	Ostype             string
	Server             string
	Body               string
}

//ToByte convert body to bytes
func (x *PKG) ToByte() []byte {
	return []byte(x.Body)
}
