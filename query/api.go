package query

import "time"

// What you need to pull a query result
type Conn interface {

	Pull() (Result, error)

	Rx(via * Via) (numOfRx int, err error)

	SetDeadline(time time.Time) error

	Close() error
}

func Dial(driverType string, addr string) (Conn, error) {
	driver, ok := registry[driverType]
	if !ok {
		return nil, errWrongType
	} else {
		return driver.Dial(addr)
	}
}

type Driver interface {

	Dial(addr string) (Conn, error)
}

var registry = make(map[string] Driver)

func PutDriver(driverType string, driver Driver) {
	registry[driverType] = driver
}
