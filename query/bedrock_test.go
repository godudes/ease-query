package query

import (
	"testing"
	"fmt"
	"time"
)

func TestPull(t *testing.T) {
	a := func(addr string) {
		conn, err := Dial("mc-bedrock", addr)
		if err != nil {
			fmt.Println(err)
			return
		}
		err = conn.SetDeadline(time.Now().Add(time.Duration(1 * time.Second)))
		if err != nil {
			fmt.Println(err)
			return
		}
		res, err := conn.Pull()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(res)
	}
	a( "play.ease"+"cation.net:19132")
	a("play.lb"+"sg.net:19132")
	a("bw.fe"+"craft.cc:19132")
	a("no.server.here:19132") // dial udp: no such host
	a("www.google.com:19132") // read udp: i/o timeout
}