package query

import (
	"testing"
	"fmt"
)

var ecAddr = "play.easecation.net:19132"

func TestPull(t *testing.T) {
	conn, err := Dial(McBedrock, ecAddr)
	if err != nil {
		t.Error(err)
	}
	res, err := conn.Pull()
	fmt.Println(err)
	fmt.Println(res)
}