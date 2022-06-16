package model

import (
	"errors"
	"github.com/jinzhu/gorm"
	"mylib/helpers"
	"time"
)

type Domain struct {
	BaseGormModel

	DomainID   string `json:"domainID" gorm:"Column:domain_id;Type:varchar(60);primary_key;not null"`
	AccountID  string `json:"accountID" gorm:"Column:account_id;Type:varchar(60);not null"`
	DomainName string `json:"domainName" gorm:"Column:domain_name;Type:varchar(100);not null"`

	CreatedAt time.Time `json:"createdAt" gorm:"Column:created_at;Type:timestamptz;not null"`
	BaseUrl   string    `json:"baseUrl" gorm:"Column:base_url;Type:text;not null"`

	CreatedBy string     `json:"createdBy" gorm:"Column:created_by;Type:varchar(60); not null"`
	UpdatedAt *time.Time `json:"updatedAt" gorm:"Column:updated_at;Type:timestamptz;null"`
	UpdatedBy *string    `json:"updatedBy" gorm:"Column:updated_by;Type:varchar(60); null"`

	DeletedUnique string     `json:"deletedUnique" gorm:"Column:deleted_unique;Type:varchar(30);null"`
	DeletedAt     *time.Time `json:"deletedAt" gorm:"Column:deleted_at;Type:timestamptz;null"`
	DeletedBy     *string    `json:"deletedBy" gorm:"Column:deleted_by;Type:varchar(60); null"`
}

func CreateDomain(db *gorm.DB) Domain {
	model := Domain{}
	model.DB = db
	return model
}

func (model Domain) Add(accountID string,
	createdBy string,
	domainName string,
	baseUrl string,

) error {

	_q := model.DB.Where(Domain{
		DomainName: domainName,
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

		model.DomainID = helpers.ID()
		model.AccountID = accountID
		model.DomainName = domainName
		model.BaseUrl = baseUrl
		model.CreatedBy = createdBy
		model.CreatedAt = time.Now().UTC()
		_q = model.DB.Create(&model)

		return _q.Error

	}
}

func (model Domain) Update(accountID string,
	updatedBy string,
	domainID string,
	domainName string,
	baseUrl string) error {

	err := model.FindByID(accountID, domainName)
	if err != nil {
		return err
	}

	_q := model.DB.Where(Domain{
		DomainID:  domainID,
		AccountID: accountID,
	}).Updates(map[string]interface{}{
		"domain_name": domainName,
		"base_url":    baseUrl,
		"updated_at":  time.Now().UTC(),
		"updated_by":  updatedBy,
	})

	return _q.Error

}

func (model Domain) Delete(accountID string,
	domainID string) error {

	err := model.FindByID(accountID, domainID)
	if err != nil {
		return err
	}

	_q := model.DB.Where(Domain{
		DomainID:  domainID,
		AccountID: accountID,
	}).Unscoped().Delete(&model)

	return _q.Error

}

func (model *Domain) FindByID(accountID string, domainID string) error {

	_q := model.DB.Where(Domain{
		DomainID:  domainID,
		AccountID: accountID,
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

func (model *Domain) FindList(accountID string) ([]Domain, error) {

	var _domains []Domain

	_q := model.DB.Where(Domain{
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
