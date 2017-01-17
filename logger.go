package glug

import (
	"log"
	"net/http"
	"time"
)

func Logger(conn *Connection) bool {
	start := time.Now()
	addr := conn.Request.Header.Get("X-Real-IP")
	if addr == "" {
		addr = conn.Request.Header.Get("X-Forwarded-For")
		if addr == "" {
			addr = conn.Request.RemoteAddr
		}
	}
	log.Printf("Started %s %s for %s", conn.Request.Method, conn.Request.URL.Path, addr)
	fun := func(resp *Resp) {
		status := resp.Status
		statusText := http.StatusText(int(status))
		duration := time.Since(start)
		log.Printf("Completed %v %s in %v\n", status, statusText, duration)
	}
	conn.Register(fun)
	return true
}
