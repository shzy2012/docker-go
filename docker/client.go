package docker

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"net"
	"strconv"
	"strings"
	"sync"

	"github.com/shzy2012/common/log"
)

/* docker api example https://docs.docker.com/develop/sdk/examples/
 * 通过http对docker api进行封装
 */

//Action 请求类型
type Action string

const (
	//POST 类型
	POST Action = "POST"
	//GET 类型
	GET Action = "GET"
)

//API Docker API
type API struct {
	Ctx   context.Context
	Debug bool
	cmdCH chan string

	sync.Mutex
	conn net.Conn
}

//NewAPI 创建Docker API
func NewAPI() *API {

	api := &API{
		cmdCH: make(chan string),
	}

	return api
}

/*Reqeust 负责与docker引擎交互，使用unix方式
 */
func (a *API) Reqeust(action Action, url, body string) (*PKG, error) {
	var err error

	if a.conn == nil {
		log.Infoln("实例化conn")
		//如果nil，重新实例化
		a.Lock()
		a.conn, err = net.Dial("unix", "/var/run/docker.sock")
		a.Unlock()

		if err != nil {
			log.Errorf("[实例化 conn]=> %s\n", err.Error())
			return nil, err
		}
	}

	urlCmd := fmt.Sprintf("%s %s HTTP/1.0\r\n\r\n", action, url)

	//处理body
	if body != "" {
		urlCmd = urlCmd + "Content-Type: application/json\r\n"
		urlCmd = urlCmd + "Content-Length: " + strconv.Itoa(len(body)) + "\r\n\r\n"
		urlCmd = urlCmd + body
	}

	if a.Debug {
		log.Debugf("[Docker send cmd]=>\n%s\n", urlCmd)
	}

	a.conn.Write([]byte(urlCmd))
	status, err := bufio.NewReader(a.conn).ReadString('\t')

	if err != nil && err.Error() != "EOF" {
		log.Errorf("[conn read]=> %s\n", err.Error())
		return nil, err
	}

	if a.Debug {
		log.Debugf("[Message from server]=>\n%s\n", status)
	}

	if strings.TrimSpace(status) == "" {
		log.Errorf("[conn read]=> %s\n", "empty")
		return nil, err
	}

	list := strings.Split(status, "\r\n")
	if a.Debug {
		log.Debugf("[Split status len]=> %v\n", len(list))
	}

	dpkg := &PKG{}
	for i, ele := range list {

		if strings.Contains(ele, "HTTP") {

			/*
				HTTP/1.0 404 Not Found
				HTTP/1.0 200 OK
				HTTP/1.0 204 No Content
				201 – no error
				400 – bad parameter
				404 – no such container
				406 – impossible to attach (container not running)
				409 – conflict
				500 – server error
			*/
			dpkg.Protocol = ele
			if strings.Contains(ele, "200") {
				dpkg.OK = true
				dpkg.Code = "200"
			} else if strings.Contains(ele, "204") {
				dpkg.OK = true
				dpkg.Code = "204"
				dpkg.Body = "No Content"
			}

		} else if strings.Contains(ele, "Api-Version") {
			dpkg.APIVersion = ele
		} else if strings.Contains(ele, "Content-Type") {
			dpkg.ContentType = ele
		} else if strings.Contains(ele, "Content-Length") {
			dpkg.ContentLength = ele
		} else if strings.Contains(ele, "Date") {
			dpkg.Date = ele
		} else if strings.Contains(ele, "Docker-Experimental") {
			dpkg.DockerExperimental = ele
		} else if strings.Contains(ele, "Ostype") {
			dpkg.Ostype = ele
		} else if strings.Contains(ele, "Server") {
			dpkg.Server = ele
		} else if strings.TrimSpace(ele) == "" {
			dpkg.Body = list[i+1]
			break
		}
	}

	if !dpkg.OK {
		return dpkg, errors.New(dpkg.Body)
	}

	return dpkg, nil
}

/*Get 执行Get操作
@url：cmd的路径 类似/images/json
*/
func (a *API) Get(url string) (*PKG, error) {
	return a.Reqeust(GET, url, "")
}

/*Post 执行POST操作
@url：cmd的路径 类似/containers/create
*/
func (a *API) Post(url string) (*PKG, error) {
	return a.Reqeust(POST, url, "")
}
