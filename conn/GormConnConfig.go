package conn

import (
	"encoding/json"
	"io/ioutil"
	"time"
)

type GormConnConfig struct {
	host             string
	port             int
	driver           string
	dbName           string
	user             string
	password         string
	enableLog        bool
	singularTable    bool
	sslMode          string
	connString       string
	maxIdleConns     int
	maxOpenConns     int
	maxConnsLifeTime time.Duration
}

func CreateGormConnConfigWithConfFile(confFile string) (*GormConnConfig, error) {

	bFile, err := ioutil.ReadFile(confFile)
	if err != nil {
		return nil, err
	}

	_gormConfig := &GormConnConfig{}
	err = json.Unmarshal(bFile, &_gormConfig)
	if err != nil {
		return nil, err
	}

	return _gormConfig, nil
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
func (cfg GormConnConfig) GetConnMaxLifeTime() time.Duration {
	return cfg.maxConnsLifeTime
}

func (cfg *GormConnConfig) SetEnableLog(enableLog bool) {
	cfg.enableLog = enableLog
}
func (cfg GormConnConfig) GetEnableLog() bool {
	return cfg.enableLog
}

func (cfg *GormConnConfig) SetSingularTable(singularTable bool) {
	cfg.singularTable = singularTable
}
func (cfg GormConnConfig) GetSingularTable() bool {
	return cfg.singularTable
}
