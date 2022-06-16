package helpers

import (
	"strconv"
	"time"
)

func ID() string {
	_no := time.Now().UnixMicro()
	time.Sleep(11 * time.Microsecond)
	return strconv.FormatInt(_no, 10)
}
