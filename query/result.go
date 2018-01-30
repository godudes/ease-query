package query

// What should be returned in a query
type Result interface {
	GetServerId() uint64

	GetResultType() Type

	GetMsgOfToday() string

	GetOnlineCount() int32

	GetMaxCount() int32

	GetBedrockNetVer() int32

	GetBedrockGameVer() string

	String() string
}

var nilResult = &nilRes{}

type nilRes struct {}

func (n *nilRes) GetServerId() uint64 {
	return 0
}

func (n *nilRes) GetResultType() Type {
	return Nil
}

func (n *nilRes) GetMsgOfToday() string {
	return ""
}

func (n *nilRes) GetOnlineCount() int32 {
	return 0
}

func (n *nilRes) GetMaxCount() int32 {
	return 0
}

func (n *nilRes) GetBedrockNetVer() int32 {
	return 0
}

func (n *nilRes) GetBedrockGameVer() string {
	return ""
}

func (n *nilRes) String() string {
	return "NilResult"
}