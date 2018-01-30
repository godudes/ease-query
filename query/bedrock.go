package query

import (
	"net"
	"bytes"
	"encoding/binary"
	"math/rand"
	"strings"
	"strconv"
)

var magic = []byte("\x00\xff\xff\x00\xfe\xfe\xfe\xfe\xfd\xfd\xfd\xfd\x12\x34\x56\x78")

type bedrockConn struct {
	*conn
}

func (c bedrockConn) Pull() (Result, error) {
	return bedrockPull(c.addr)
}

func newBedrockConn(addr string) bedrockConn {
	return bedrockConn{
		&conn{
			McBedrock, addr,
		},
	}
}

func bedrockPull(addr string) (Result, error) {
	var err error
	conn, err := net.Dial("udp", addr)
	if err != nil {
		return nilResult, err
	}
	defer func() {conn.Close()} ()

	reqPacket := new(bytes.Buffer)
	reqPacket.WriteByte(0x01)
	pingId :=  rand.Uint32()
	binary.Write(reqPacket, binary.BigEndian, pingId)
	reqPacket.Write(magic)
	reqPacket.WriteByte(0)
	binary.Write(reqPacket, binary.BigEndian, reqPacket.Len())
	conn.Write(reqPacket.Bytes())

	respBytes := make([]byte, 4096)
	n, err := conn.Read(respBytes)
	if err != nil {
		return nilResult, err
	}
	defer func() {conn.Close()} ()
	respPacket := bytes.NewBuffer(respBytes[35:n]) // jump header
	res := respPacket.String()
	arr := strings.Split(res, ";")
	c := converse{func(e error) {
		err = e
	}}
	ret := &resultImpl{
		typ: McBedrock,
		msgOfToday: arr[1],
		onlineCount: c.strToInt(arr[4]),
		maxCount: c.strToInt(arr[5]),
	}
	return ret, err
}

type converse struct{
	onErr func(error)
}

func (c *converse) strToInt(in string) int {
	o, err := strconv.Atoi(in)
	if err != nil {
		c.onErr(err)
	}
	return o
}
