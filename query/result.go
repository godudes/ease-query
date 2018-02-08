package query

// What should be returned in a query
type Result interface {
	GetServerId() uint64

	GetMsgOfToday() string

	GetOnlineCount() int32

	GetMaxCount() int32

	GetBedrockNetVer() int32

	GetBedrockGameVer() string

	String() string
}

type Via struct {
	ServerId *uint64
	MsgOfToday *string
	OnlineCount *int32
	MaxCount *int32
	BedrockNetVer *int32
	BedrockGameVer *string
}

