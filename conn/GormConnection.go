package conn

import (
	"github.com/jinzhu/gorm"
)

type GormConnection struct {
	db               *gorm.DB
	connectionType   int
	developmentPhase int
	connectionConfig IGormConnectionConfig
}

func (conn *GormConnection) SetConnectionConfig(connectionConfig IConnectionConfig) {
	conn.connectionConfig = connectionConfig.(IGormConnectionConfig)
}

func (conn GormConnection) GetConnectionConfig() IConnectionConfig {
	return conn.connectionConfig
}

func (conn *GormConnection) OpenConnection() error {

	connection, err := gorm.Open(conn.connectionConfig.GetDriver(),
		conn.connectionConfig.GetConnectionString())

	if err != nil {
		return err
	}
	connection.LogMode(conn.connectionConfig.GetEnableLog())
	connection.SingularTable(conn.connectionConfig.GetEnableLog())

	conn.db = connection
	conn.db.DB().SetMaxIdleConns(conn.connectionConfig.GetMaxIdleConns())
	conn.db.DB().SetMaxOpenConns(conn.connectionConfig.GetMaxOpenConns())
	conn.db.DB().SetConnMaxLifetime(conn.connectionConfig.GetConnMaxLifeTime())

	return nil

}

func (conn *GormConnection) CloseConnection() error {
	if conn.db != nil {
		err := conn.db.Close()
		if err != nil {
			return err
		}
		conn.db = nil
	}
	return nil
}

func (conn *GormConnection) SetConnectionType(connType int) {
	conn.connectionType = connType
}

func (conn GormConnection) GetConnectionType() int {
	return conn.connectionType
}

func (conn *GormConnection) SetDevelopmentPhase(developmentPhase int) {
	conn.developmentPhase = developmentPhase
}

func (conn GormConnection) GetDevelopmentPhase() int {
	return conn.developmentPhase
}

func (conn *GormConnection) SetConnConfig(connConfig IGormConnectionConfig) {
	conn.connectionConfig = connConfig
}

func (conn GormConnection) GetConnConfig() IGormConnectionConfig {
	return conn.connectionConfig
}
