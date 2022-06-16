package conntests

import (
	"mylib/conn"
	"testing"
)

func TestGormConn (t *testing.T) {

	connConfig :=

	conn, err := conn.CreateConnection(conn.CONN_TYPE_GORM,
		conn.DEV_PHASE_DEV,
		)
	if err!=nil {
		t.Fatal(err)
	}

}
