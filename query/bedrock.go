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
	pingId :=  rand.Uint64()
	binary.Write(reqPacket, binary.BigEndian, pingId)
	reqPacket.Write(magic)
	binary.Write(reqPacket, binary.BigEndian, reqPacket.Len())
	conn.Write(reqPacket.Bytes())

	resBytes := make([]byte, 4096)
	n, err := conn.Read(resBytes)
	if err != nil {
		return nilResult, err
	}
	defer func() {conn.Close()} ()

	if n < 33 {
		return nil, errBedrockWrongLen
	}
	if resBytes[0] != 0x1c {
		return nil, errBedrockWrongMsgId
	}
	if !bytes.Equal(resBytes[1:9], reqPacket.Bytes()[1:9]) {
		return nil, errBedrockWrongPongId
	}
	sid := binary.BigEndian.Uint64(resBytes[9:17])
	if !bytes.Equal(resBytes[17:33], magic) {
		return nil, errBedrockWrongMagic
	}
	resLen := binary.BigEndian.Uint16(resBytes[33:35])
	if 35 + int(resLen) > n {
		return nil, errBedrockWrongLen
	}
	res := bytes.NewBuffer(resBytes[35:35+resLen]).String() // jump header
	arr := strings.Split(res, ";")
	if len(arr) < 5 {
		return nil, errBedrockWrongFmt
	}
	ret := &resultImpl{
		serverId: sid,
		typ: McBedrock,
		msgOfToday: arr[1],
		onlineCount: strToInt(arr[4], &err),
		maxCount: strToInt(arr[5], &err),
	}
	return ret, err
}

func strToInt(in string, errOut *error) int {
	o, err := strconv.Atoi(in)
	if err != nil {
		*errOut = err
	}
	return o
}
