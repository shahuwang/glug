package glug

import (
	"log"
	"time"
)

func Logger(conn *Connection) {
	start := time.Now()
	addr := conn.Request.Header.Get("X-Real-IP")
	if addr == "" {
		addr = conn.Request.Header.Get("X-Forwarded-For")
		if addr == "" {
			addr = conn.Request.RemoteAddr
		}
	}
	log.Printf("Started %s %s for %s", conn.Request.Method, conn.Request.URL.Path, addr)

}

// func LoggerBeforeSend(resp *Resp) {
// 	conn := resp.Conn
// 	r := conn.response
// 	log.Printf("Completed %v %s in %v\n", r.Header().Get("status"), time.Since)
// }
