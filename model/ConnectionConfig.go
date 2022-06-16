package model

import (
	"errors"
	"github.com/jinzhu/gorm"
	"mylib/helpers"
	"time"
)

type ConnectionConfig struct {
	BaseGormModel

	ConfigID  string `json:"configID" gorm:"Column:config_id;Type:varchar(60);primary_key;not null"`
	AccountID string `json:"accountID" gorm:"Column:account_id;Type:varchar(60);primary_key;not null"`

	ConfigName string    `json:"configName" gorm:"Column:config_name;Type:varchar(100);not null"`
	CreatedAt  time.Time `json:"createdAt" gorm:"Column:created_at;Type:timestamptz;not null"`

	Host           string `json:"host" gorm:"Column:host;Type:varchar(15);not null"`
	Port           int    `json:"port" gorm:"Column:port;Type:integer;not null"`
	Driver         string `json:"driver" gorm:"Column:driver;Type:varchar(60);not null"`
	User           string `json:"user" gorm:"Column:user;Type:varchar(60);not null"`
	Password       string `json:"password" gorm:"Column:password;Type:text;not null"`
	ConnectionType int    `json:"connectionType" gorm:"Column:connection_type;Type:integer;not null"`
	DbType         int    `json:"dbType" gorm:"Column:db_type;Type:integer;not null"`

	enableLog        bool   `json:"enableLog" gorm:"Column:enable_log;null"`
	singularTable    bool   `json:"singularTable" gorm:"Column:singular_table;null"`
	sslMode          string `json:"sslMode" gorm:"Column:ssl_mode;Type:varchar(60);null"`
	connString       string `json:"connString" gorm:"Column:conn_string;Type:text;null"`
	maxIdleConns     int    `json:"maxIdleConns" gorm:"Column:max_idle_conns;not null"`
	maxOpenConns     int    `json:"maxOpenConns" gorm:"Column:max_open_cons;not null"`
	maxConnsLifeTime int    `json:"maxConnsLifeTime" gorm:"Column:max_conns_life_time;not null"`

	CreatedBy string     `json:"createdBy" gorm:"Column:created_by;Type:varchar(60); not null"`
	UpdatedAt *time.Time `json:"updatedAt" gorm:"Column:updated_at;Type:timestamptz;null"`
	UpdatedBy *string    `json:"updatedBy" gorm:"Column:updated_by;Type:varchar(60); null"`

	DeletedUnique string     `json:"deletedUnique" gorm:"Column:deleted_unique;Type:varchar(30);null"`
	DeletedAt     *time.Time `json:"deletedAt" gorm:"Column:deleted_at;Type:timestamptz;null"`
	DeletedBy     *string    `json:"deletedBy" gorm:"Column:deleted_by;Type:varchar(60); null"`
}

func CreateConnectionConfig(db *gorm.DB) ConnectionConfig {
	model := ConnectionConfig{}
	model.DB = db
	return model
}

func (model ConnectionConfig) Add(accountID string,
	createdBy string,
	configName string,
	host string,
	port int,
	driver string,
	user string,
	password string,
	connectionType int,
	dbType int,
	enableLog bool,
	singularTable bool,
	sslMode string,
	connString string,
	maxIdleConns int,
	maxOpenConns int,
	maxConnsLifeTime int,

) error {

	_q := model.DB.Where(ConnectionConfig{
		ConfigName: domainName,
		AccountID:  accountID,
	}).Find(&model)

	_exist := false
	if _q.Error != nil {
		if !gorm.IsRecordNotFoundError(_q.Error) {
			return _q.Error
		} else {
		}
	} else {
		_exist = true
	}

	if _exist {
		return errors.New("domain_name_" + domainName + "_already_exists")
	} else {

		model.ConnectionConfigID = helpers.ID()
		model.AccountID = accountID
		model.ConfigName = domainName
		model.BaseUrl = baseUrl
		model.CreatedBy = createdBy
		model.CreatedAt = time.Now().UTC()
		_q = model.DB.Create(&model)

		return _q.Error

	}
}

func (model ConnectionConfig) Update(accountID string,
	updatedBy string,
	domainID string,
	domainName string,
	baseUrl string) error {

	err := model.FindByID(accountID, domainName)
	if err != nil {
		return err
	}

	_q := model.DB.Where(ConnectionConfig{
		ConnectionConfigID: domainID,
		AccountID:          accountID,
	}).Updates(map[string]interface{}{
		"domain_name": domainName,
		"base_url":    baseUrl,
		"updated_at":  time.Now().UTC(),
		"updated_by":  updatedBy,
	})

	return _q.Error

}

func (model ConnectionConfig) Delete(accountID string,
	domainID string) error {

	err := model.FindByID(accountID, domainID)
	if err != nil {
		return err
	}

	_q := model.DB.Where(ConnectionConfig{
		ConnectionConfigID: domainID,
		AccountID:          accountID,
	}).Unscoped().Delete(&model)

	return _q.Error

}

func (model *ConnectionConfig) FindByID(accountID string, domainID string) error {

	_q := model.DB.Where(ConnectionConfig{
		ConnectionConfigID: domainID,
		AccountID:          accountID,
	}).Find(&model)

	if _q.Error != nil {
		if gorm.IsRecordNotFoundError(_q.Error) {
			return errors.New("cannot_find_domain_id_" + domainID)
		} else {
			return _q.Error
		}
	} else {
		return nil
	}
}

func (model *ConnectionConfig) FindList(accountID string) ([]ConnectionConfig, error) {

	var _domains []ConnectionConfig

	_q := model.DB.Where(ConnectionConfig{
		AccountID: accountID,
	}).Find(&_domains)

	if _q.Error != nil {
		if gorm.IsRecordNotFoundError(_q.Error) {
			return nil, errors.New("cannot_find_domains_with_account_id_" + accountID)
		} else {
			return nil, _q.Error
		}
	} else {
		return _domains, nil
	}
}
