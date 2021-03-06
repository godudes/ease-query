package query_test

import (
	"testing"
	"fmt"
	"time"

	"github.com/godudes/ease-query/query"
)

func TestPull(t *testing.T) {
	a := func(addr string) {
		conn, err := query.Dial("mc-bedrock", addr)
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
	a("play.ease"+"cation.net:19132")
	a("play.lb"+"sg.net:19132")
	a("bw.fe"+"craft.cc:19132")
	a("no.server.here:19132") // dial udp: no such host
	a("www.google.com:19132") // read udp: i/o timeout
}

func TestRx(t *testing.T) {

	a := func(addr string) {
		conn, err := query.Dial("mc-bedrock", addr)
		err = conn.SetDeadline(time.Now().Add(time.Duration(5 * time.Second)))
		if err != nil {
			fmt.Println(err)
			return
		}
		var (
			onlineCount int32
			maxCount int32
		)
		via := &query.Via{
			OnlineCount: &onlineCount,
			MaxCount: &maxCount,
		}
		i := 0
		for t := range time.Tick(100 * time.Millisecond){
			n, err := conn.Rx(via)
			if err != nil {
				fmt.Println(err)
				break
			}
			fmt.Printf("Time=[%s], N=[%d], Count=[%d/%d]\n", t.String(), n, onlineCount, maxCount)
			i++; if i >= 20 { break }
		}
	}
	a( "play.ease"+"cation.net:19132")
}