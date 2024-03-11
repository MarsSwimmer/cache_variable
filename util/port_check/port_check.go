package port_check

import (
	"fmt"
	"net"
	"time"
)

// 检查端口是否可用
func CheckPortAvaiable(port int32) bool {
	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return false
	}
	defer ln.Close()

	time.Sleep(100 * time.Millisecond)
	return true
}
