package query

// Type of a query. For example, McBedrock, McJava, etc.
type Type int

const (
	Nil = iota
	McBedrock
)

// What you need to pull a query result
type Conn interface {

	Pull() (Result, error)
}

type conn struct {
	typ Type
	addr string
}

func (c *conn) Pull() (Result, error) {
	return nil, nil
}

func Dial(typ Type, addr string) (Conn, error) {
	switch typ {
	case McBedrock:
		return newBedrockConn(addr), nil
	default:
		return nil, errWrongType
	}
}


