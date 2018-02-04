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
	serverId *uint64
	msgOfToday *string
	onlineCount *int32
	maxCount *int32
	bedrockNetVer *int32
	bedrockGameVer *string
}

