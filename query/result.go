package query

import "fmt"

// What should be returned in a query
type Result interface {
	GetResultType() Type

	GetMsgOfToday() string

	GetOnlineCount() int

	GetMaxCount() int

	String() string
}

var nilResult = &resultImpl{
	typ: Nil,
}

type resultImpl struct {
	typ Type
	msgOfToday string
	onlineCount int
	maxCount int
}

func (n *resultImpl) GetResultType() Type {
	return n.typ
}

func (n *resultImpl) GetMsgOfToday() string {
	return n.msgOfToday
}

func (n *resultImpl) GetOnlineCount() int {
	return n.onlineCount
}

func (n *resultImpl) GetMaxCount() int {
	return n.maxCount
}

func (n *resultImpl) String() string {
	return fmt.Sprintf("Result{Type=%d, MOTD=%s, OnlineCount=%d, MaxCount=%d}",
		n.typ, n.msgOfToday, n.onlineCount, n.maxCount)
}