package query

// What you need to pull a query result
type Conn interface {

	Pull() (Result, error)

	Close() error
}

func Dial(driverType string, addr string) (Conn, error) {
	val, ok := registry[driverType]
	if !ok {
		return nil, errWrongType
	} else {
		return val.Dial(addr)
	}
}

type Driver interface {

	Dial(addr string) (Conn, error)
}

var registry = make(map[string] Driver)

func PutDriver(driverType string, driver Driver) {
	registry[driverType] = driver
}
