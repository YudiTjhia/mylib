package conntest

import (
	"mylib/conn"
	"testing"
)

func TestGormConn_OpenCloseConnection(t *testing.T) {

	_connConfig, err := conn.CreateGormConnConfigWithConfFile("../conf/db.conf.dev.json")
	if err != nil {
		t.Fatal(err)
	}

	conn, err := conn.CreateConnection(conn.CONN_TYPE_GORM,
		conn.DEV_PHASE_DEV,
		_connConfig,
	)
	if err != nil {
		t.Fatal(err)
	}

	err = conn.CloseConnection()
	if err != nil {
		t.Fatal(err)
	}

}
