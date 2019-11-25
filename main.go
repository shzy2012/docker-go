package main

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
)

func main() {
	//api := docker.NewAPI()
	// list, _ := api.ContainerList()
	// for _, ele := range *list {
	// 	fmt.Printf("%s-%s-%s-%v\n", ele.GetID(), ele.Names, ele.GetImageID(), ele.Ports)
	// }

	body := `
{
    "Image": "mysql:latest",
    "ExposedPorts": {
		"3306/tcp": {},
		"3307/tcp": {},
		"3308/tcp": {}
	},
	"Env":["MYSQL_ROOT_PASSWORD=my-secret-pw"],
	"PortBindings": {
        "3306/tcp": [
			{ "HostPort": "6009" },
			{ "HostPort": "6001" }
		],
		"3307/tcp": [
			{ "HostPort": "6000" }
		],
		"3308/tcp": [
			{ "HostPort": "6008" }
        ]
    }
}`
	conn, _ := net.Dial("unix", "/var/run/docker.sock")

	urlCmd := fmt.Sprintf("%s %s HTTP/1.0\r\n", "POST", "/containers/create?name=shzy")
	//处理body
	if body != "" {
		urlCmd = urlCmd + "Content-Type: application/json\r\n"
		urlCmd = urlCmd + "Content-Length: " + strconv.Itoa(len(body)) + "\r\n\r\n"
		urlCmd = urlCmd + body
	}

	conn.Write([]byte(urlCmd))
	status, _ := bufio.NewReader(conn).ReadString('\t')

	fmt.Printf("%s\n", status)

	// body := `{"Image":"jwilder/nginx-proxy:latest","ExposedPorts":{"8080/tcp":{}} }`
	// conn, _ := net.Dial("unix", "/var/run/docker.sock")

	// urlCmd := fmt.Sprintf("%s %s HTTP/1.0\r\n", "POST", "/containers/shzy2012/start")
	// //处理body
	// if body != "" {
	// 	urlCmd = urlCmd + "Content-Type: application/json\r\n"
	// 	urlCmd = urlCmd + "Content-Length: " + strconv.Itoa(len(body)) + "\r\n\r\n"
	// 	urlCmd = urlCmd + body
	// }

	// conn.Write([]byte(urlCmd))
	// status, _ := bufio.NewReader(conn).ReadString('\t')

	// fmt.Printf("%s\n", status)

}
