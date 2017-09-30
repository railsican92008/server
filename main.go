package main

import (
	"fmt"
	"net"
	_ "server/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

func main() {
	beego.Run()
}

func init() {

	host := beego.AppConfig.DefaultString("socket::host", "localhost")
	port := beego.AppConfig.DefaultString("socket::port", "8001")
	path := fmt.Sprintf("%s:%s", host, port)

	go func() {
		listen, err := net.Listen("tcp", path)
		if err != nil {
			logs.Error("Socket Server Error: ", err)
		}
		defer listen.Close()
	}()
}
