package query

import (
	"testing"
	"fmt"
)

func TestPull(t *testing.T) {
	a := func(addr string) {
		conn, err := Dial("mc-bedrock", addr)
		if err != nil {
			t.Error(err)
		}
		res, err := conn.Pull()
		fmt.Println(res)
	}
	a( "play.ease"+"cation.net:19132")
	a("play.lb"+"sg.net:19132")
	a("bw.fe"+"craft.cc:19132")
}