package main

import (
	"fmt"
	"net"
	_ "server/routers"
	"server/socket"
	"strings"

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

	//start socket server
	go func() {
		listen, err := net.Listen("tcp", path)
		if err != nil {
			logs.Error("socket server error: ", err)
		}
		defer listen.Close()

		logs.Debug("socket server start success...")

		for {
			conn, err := listen.Accept()
			if err != nil {
				continue
			}
			logs.Debug("TCP connect: ", conn.RemoteAddr().String())
			hosts := beego.AppConfig.Strings("allow")
			flag := true
			for _, ip := range hosts {
				if strings.HasPrefix(conn.RemoteAddr().String(), ip) {
					flag = false
				}
			}
			if flag {
				logs.Debug("IP: ", conn.RemoteAddr().String(), " is Not allow")
				conn.Close()
			}
			socket.Handle(conn)
		}
	}()
}
