package model

import (
	"errors"
	"github.com/jinzhu/gorm"
	"time"
)

type RemoteConfig struct {
	BaseGormModel

	ConfigKey string    `json:"configKey" gorm:"Column:config_key;Type:varchar(100);primary_key;not null"`
	AccountID string    `json:"accountID" gorm:"Column:account_id;Type:varchar(60);primary_key;not null"`
	CreatedAt time.Time `json:"createdAt" gorm:"Column:created_at;Type:timestamptz;not null"`

	ConfigValue string `json:"configValue" gorm:"Column:config_value;Type:text;not null"`

	CreatedBy string     `json:"createdBy" gorm:"Column:created_by;Type:varchar(60); not null"`
	UpdatedAt *time.Time `json:"updatedAt" gorm:"Column:updated_at;Type:timestamptz;null"`
	UpdatedBy *string    `json:"updatedBy" gorm:"Column:updated_by;Type:varchar(60); null"`

	DeletedUnique string     `json:"deletedUnique" gorm:"Column:deleted_unique;Type:varchar(30);null"`
	DeletedAt     *time.Time `json:"deletedAt" gorm:"Column:deleted_at;Type:timestamptz;null"`
	DeletedBy     *string    `json:"deletedBy" gorm:"Column:deleted_by;Type:varchar(60); null"`
}

func CreateRemoteConfig(db *gorm.DB) RemoteConfig {
	model := RemoteConfig{}
	model.DB = db
	return model
}

func (model RemoteConfig) Add(accountID string,
	createdBy string,
	configKey string,
	configValue string,

) error {

	_q := model.DB.Where(RemoteConfig{
		ConfigKey: configKey,
		AccountID: accountID,
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
		return errors.New("config_key_" + configKey + "_already_exists")
	} else {

		model.ConfigKey = configKey
		model.AccountID = accountID
		model.ConfigValue = configValue
		model.CreatedBy = createdBy
		model.CreatedAt = time.Now().UTC()
		_q = model.DB.Create(&model)

		return _q.Error

	}
}

func (model RemoteConfig) Update(accountID string,
	updatedBy string,
	configKey string,
	configValue string) error {

	err := model.FindByID(accountID, configKey)
	if err != nil {
		return err
	}

	_q := model.DB.Where(RemoteConfig{
		ConfigKey: configKey,
		AccountID: accountID,
	}).Updates(map[string]interface{}{
		"config_value": configValue,
		"updated_at":   time.Now().UTC(),
		"updated_by":   updatedBy,
	})

	return _q.Error

}

func (model RemoteConfig) Delete(accountID string,
	configKey string) error {

	err := model.FindByID(accountID, configKey)
	if err != nil {
		return err
	}

	_q := model.DB.Where(RemoteConfig{
		ConfigKey: configKey,
		AccountID: accountID,
	}).Unscoped().Delete(&model)

	return _q.Error

}

func (model *RemoteConfig) FindByID(accountID string,
	configKey string) error {

	_q := model.DB.Where(RemoteConfig{
		ConfigKey: configKey,
		AccountID: accountID,
	}).Find(&model)

	if _q.Error != nil {
		if gorm.IsRecordNotFoundError(_q.Error) {
			return errors.New("cannot_find_config_key_" + configKey)
		} else {
			return _q.Error
		}
	} else {
		return nil
	}
}

func (model *RemoteConfig) FindList(accountID string) ([]RemoteConfig, error) {

	var _remoteConfigs []RemoteConfig

	_q := model.DB.Where(RemoteConfig{
		AccountID: accountID,
	}).Find(&_remoteConfigs)

	if _q.Error != nil {
		if gorm.IsRecordNotFoundError(_q.Error) {
			return nil, errors.New("cannot_find_remote_configs_with_account_id_" + accountID)
		} else {
			return nil, _q.Error
		}
	} else {
		return _remoteConfigs, nil
	}

}
