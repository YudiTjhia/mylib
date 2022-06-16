package conn

import "errors"

const CONN_TYPE_GORM = 1

const DEV_PHASE_DEV = 1
const DEV_PHASE_STAGING = 2
const DEV_PHASE_PROD = 3

func CreateConnection(connType int,
	developmentPhase int,
	connConfig IConnectionConfig) (IConnection, error) {

	switch connType {

	case CONN_TYPE_GORM:
		conn := GormConnection{}
		conn.SetConnectionType(connType)
		conn.SetDevelopmentPhase(developmentPhase)
		conn.SetConnConfig(connConfig.(IGormConnectionConfig))
		return &conn, nil
	}

	return nil, errors.New("invalid_connection_type")
}
