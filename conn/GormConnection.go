package conn

import (
	"github.com/jinzhu/gorm"
	"time"
)

type GormConnConfig struct {
	host             string
	port             int
	driver           string
	dbName           string
	user             string
	password         string
	sslMode          string
	connString       string
	maxIdleConns     int
	maxOpenConns     int
	maxConnsLifeTime time.Duration
}

func (cfg *GormConnConfig) SetHost(host string) {
	cfg.host = host
}
func (cfg GormConnConfig) GetHost() string {
	return cfg.host
}
func (cfg *GormConnConfig) SetPort(port int) {
	cfg.port = port
}
func (cfg GormConnConfig) GetPort() int {
	return cfg.port
}
func (cfg *GormConnConfig) SetDriver(driver string) {
	cfg.driver = driver
}
func (cfg GormConnConfig) GetDriver() string {
	return cfg.driver
}

func (cfg *GormConnConfig) SetDbName(dbName string) {
	cfg.dbName = dbName
}
func (cfg GormConnConfig) GetDbName() string {
	return cfg.dbName
}

func (cfg *GormConnConfig) SetUser(user string) {
	cfg.user = user
}
func (cfg GormConnConfig) GetUser() string {
	return cfg.user
}

func (cfg *GormConnConfig) SetPassword(password string) {
	cfg.password = password
}
func (cfg GormConnConfig) GetPassword() string {
	return cfg.password
}

func (cfg *GormConnConfig) SetSSLMode(sslMode string) {
	cfg.sslMode = sslMode
}
func (cfg GormConnConfig) GetSSLMode() string {
	return cfg.sslMode
}

func (cfg *GormConnConfig) SetConnectionString(connString string) {
	cfg.connString = connString
}
func (cfg GormConnConfig) GetConnectionString() string {
	return cfg.connString
}

func (cfg *GormConnConfig) SetMaxIdleConns(maxIdleConns int) {
	cfg.maxIdleConns = maxIdleConns
}
func (cfg GormConnConfig) GetMaxIdleConns() int {
	return cfg.maxIdleConns
}

func (cfg *GormConnConfig) SetMaxOpenConns(maxOpenConns int) {
	cfg.maxOpenConns = maxOpenConns
}
func (cfg GormConnConfig) GetMaxOpenConns() int {
	return cfg.maxOpenConns
}

func (cfg *GormConnConfig) SetConnMaxLifetime(maxConnsLifeTime time.Duration) {
	cfg.maxConnsLifeTime = maxConnsLifeTime
}
func (cfg GormConnConfig) GetConnMaxLifeTime() int {
	return cfg.maxOpenConns
}

type GormConnection struct {
	db *gorm.DB

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
