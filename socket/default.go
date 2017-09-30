package socket

import (
	"net"

	"github.com/astaxie/beego/logs"
)

//Handle data receive function
func Handle(conn net.Conn) {

	buffer := make([]byte, 2048)

	for {

		n, err := conn.Read(buffer)
		if err != nil {
			logs.Debug("connection error: ", err, " IP: ", conn.RemoteAddr().String())
			conn.Close()
			return
		}

		logs.Debug("[DATA] --> ", string(buffer[:n]))
	}
}
