package conn

import "time"

type IGormConnectionConfig interface {
	IConnectionConfig

	SetEnableLog(enableLog bool)
	GetEnableLog() bool

	SetSingularTable(singularTable bool)
	GetSingularTable() bool
}

type IConnectionConfig interface {
	SetHost(host string)
	GetHost() string

	SetPort(port int)
	GetPort() int

	SetDriver(driver string)
	GetDriver() string

	SetDbName(dbName string)
	GetDbName() string

	SetUser(user string)
	GetUser() string

	SetPassword(password string)
	GetPassword() string

	SetSSLMode(sslMode string)
	GetSSLMode() string

	SetConnectionString(connString string)
	GetConnectionString() string

	SetMaxIdleConns(maxIdleConns int)
	GetMaxIdleConns() int

	SetMaxOpenConns(maxOpenConns int)
	GetMaxOpenConns() int

	SetConnMaxLifetime(duration time.Duration)
	GetConnMaxLifeTime() time.Duration
}

type IConnection interface {
	OpenConnection() error
	CloseConnection() error

	SetConnectionType(connectionType int)
	GetConnectionType() int

	SetDevelopmentPhase(developmentPhase int)
	GetDevelopmentPhase() int

	SetConnectionConfig(connectionConfig IConnectionConfig)
	GetConnectionConfig() IConnectionConfig
}
