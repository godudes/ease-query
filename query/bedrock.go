package query

import (
	"net"
	"bytes"
	"encoding/binary"
	"math/rand"
	"strings"
	"strconv"
	"fmt"
	"math"
	"time"
)

func init() {
	PutDriver("mc-bedrock", new(bedrockDriver))
}
type bedrockDriver struct {}

func (d bedrockDriver) Dial(addr string) (Conn, error) {
	return newBedrockConn(addr)
}

var magic = []byte{0x00, 0xff, 0xff, 0x00, 0xfe, 0xfe, 0xfe, 0xfe, 0xfd, 0xfd, 0xfd, 0xfd, 0x12, 0x34, 0x56, 0x78}
var validDataFormat = "MC"+"PE" // I hate typo checks

type bedrockResult struct {
	serverId uint64
	msgOfToday string
	onlineCount int32
	maxCount int32
	bedrockNetVer int32
	bedrockGameVer string
}

func (n bedrockResult) GetServerId() uint64 {
	return n.serverId
}

func (n bedrockResult) GetMsgOfToday() string {
	return n.msgOfToday
}

func (n bedrockResult) GetOnlineCount() int32 {
	return n.onlineCount
}

func (n bedrockResult) GetMaxCount() int32 {
	return n.maxCount
}

func (n bedrockResult) GetBedrockNetVer() int32 {
	return n.bedrockNetVer
}

func (n bedrockResult) GetBedrockGameVer() string {
	return n.bedrockGameVer
}

func (n *bedrockResult) String() string {
	return fmt.Sprintf("BedrockResult{ServerId=%d, " +
		"NetVer=%d, GameVer=%s, MOTD=%s, OnlineCount=%d, MaxCount=%d}",
		n.serverId, n.bedrockNetVer, n.bedrockGameVer, n.msgOfToday, n.onlineCount, n.maxCount)
}

type bedrockConn struct {
	conn net.Conn
}

func (c bedrockConn) Pull() (Result, error) {
	return bedrockPing(c.conn)
}

func (c bedrockConn) Rx(via * Via) (numOfRx int, err error) {
	numOfRx = 0
	res, err := bedrockPing(c.conn)
	if err != nil {
		return
	}
	if via.ServerId			!= nil { *via.ServerId		= res.GetServerId()		  ;numOfRx++	}
	if via.MsgOfToday		!= nil { *via.MsgOfToday	= res.GetMsgOfToday()	  ;numOfRx++	}
	if via.OnlineCount		!= nil { *via.OnlineCount	= res.GetOnlineCount()	  ;numOfRx++	}
	if via.MaxCount			!= nil { *via.MaxCount		= res.GetMaxCount()		  ;numOfRx++	}
	if via.BedrockNetVer	!= nil { *via.BedrockNetVer	= res.GetBedrockNetVer()  ;numOfRx++	}
	if via.BedrockGameVer	!= nil { *via.BedrockGameVer= res.GetBedrockGameVer() ;numOfRx++	}
	return
}

func (c bedrockConn) SetDeadline(time time.Time) error {
	return c.conn.SetDeadline(time)
}

func (c bedrockConn) Close() error {
	return c.conn.Close()
}

func newBedrockConn(addr string) (c *bedrockConn, err error) {
	conn, err := net.Dial("udp", addr)
	if err != nil {
		return
	}
	c = new(bedrockConn)
	c.conn = conn
	return
}

// Pull RakNet UNCONNECTED_PONG packet
func bedrockPing(conn net.Conn) (res Result, err error) {
	reqPacket := new(bytes.Buffer)
	reqPacket.WriteByte(0x01)
	pingId := rand.Uint64()
	binary.Write(reqPacket, binary.BigEndian, pingId)
	reqPacket.Write(magic)
	binary.Write(reqPacket, binary.BigEndian, reqPacket.Len())
	conn.Write(reqPacket.Bytes())

	resBytes := make([]byte, 4096)
	n, err := conn.Read(resBytes)
	if err != nil {
		return nil, err
	}
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
	resStr := bytes.NewBuffer(resBytes[35:35+resLen]).String() // jump header
	arr := strings.Split(resStr, ";")
	if len(arr) < 5 {
		return nil, errBedrockWrongFmt
	}
	if arr[0] != validDataFormat {
		return nil, errBedrockWrongFmt
	}
	res = &bedrockResult{
		serverId: sid,
		msgOfToday: arr[1],
		bedrockNetVer: strToInt32(arr[2], &err),
		bedrockGameVer: arr[3],
		onlineCount: strToInt32(arr[4], &err),
		maxCount: strToInt32(arr[5], &err),
	}
	return
}

func strToInt32(in string, errOut *error) int32 {
	o, err := strconv.Atoi(in)
	if err != nil {
		*errOut = err
	}
	if o > math.MaxInt32 {
		*errOut = errBedrockWrongFmt
	}
	return int32(o)
}
