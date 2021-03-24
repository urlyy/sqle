package model

import (
	"database/sql"
	"fmt"
	"strings"
	"sync"
	"time"

	"actiontech.cloud/universe/sqle/v4/sqle/errors"
	"actiontech.cloud/universe/sqle/v4/sqle/log"
	"github.com/jinzhu/gorm"
	"github.com/jmoiron/sqlx"
	"github.com/jmoiron/sqlx/reflectx"
)

var storage *Storage

var storageMutex sync.Mutex

const dbDriver = "mysql"

func InitStorage(s *Storage) {
	storageMutex.Lock()
	defer storageMutex.Unlock()
	storage = s
}

func GetStorage() *Storage {
	storageMutex.Lock()
	defer storageMutex.Unlock()
	return storage
}

func UpdateStorage(newStorage *Storage) {
	storageMutex.Lock()
	defer storageMutex.Unlock()
	storage.db.Close()
	storage = newStorage
}

func GetDb() *gorm.DB {
	return storage.db
}

func GetSqlxDb() *sqlx.DB {
	db := sqlx.NewDb(storage.db.DB(), dbDriver)
	db.Mapper = reflectx.NewMapperFunc("json", strings.ToLower)
	return db
}

type Model struct {
	ID        uint       `json:"id" gorm:"primary_key" example:"1"`
	CreatedAt time.Time  `json:"created_at" gorm:"default:current_timestamp" example:"2018-10-21T16:40:23+08:00"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"default:current_timestamp on update current_timestamp" example:"2018-10-21T16:40:23+08:00"`
	DeletedAt *time.Time `json:"-" sql:"index"`
}

func NewStorage(user, password, host, port, schema string, debug bool) (*Storage, error) {
	log.Logger().Infof("connecting to storage, host: %s, port: %s, user: %s, schema: %s",
		host, port, user, schema)
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user, password, host, port, schema))
	if err != nil {
		log.Logger().Errorf("connect to storage failed, error: %v", err)
		return nil, errors.New(errors.CONNECT_STORAGE_ERROR, err)
	}
	if debug {
		db.SetLogger(log.Logger().WithField("type", "sql"))
		db.LogMode(true)
	}
	log.Logger().Info("connected to storage")
	return &Storage{db: db}, errors.New(errors.CONNECT_STORAGE_ERROR, err)
}

type Storage struct {
	db *gorm.DB
}

func (s *Storage) AutoMigrate() error {
	err := s.db.AutoMigrate(
		&Instance{},
		&RuleTemplate{},
		&Rule{},
		&Task{},
		&ExecuteSQL{},
		&RollbackSQL{},
		&SqlWhitelist{},
		&User{},
		&Role{},
		&WorkflowTemplate{},
		&WorkflowStepTemplate{},
		&Workflow{},
		&WorkflowRecord{},
		&WorkflowStep{},
	).Error
	return errors.New(errors.CONNECT_STORAGE_ERROR, err)
}

func (s *Storage) CreateRulesIfNotExist(rules []Rule) error {
	allRules, err := s.GetAllRule()
	if err != nil {
		return err
	}
	allRulesMap := make(map[string]string)
	for _, rule := range allRules {
		allRulesMap[rule.Name] = rule.Value
	}
	for _, rule := range rules {
		/*If rule exist
		1.value is used, skip init
		2.new value is empty,skip init (otherwise it will cause a panic)
		*/
		if existRuleValue, ok := allRulesMap[rule.Name]; ok && (existRuleValue != "" || rule.Value == "") {
			continue
		}
		err = s.Save(rule)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Storage) CreateDefaultTemplate(rules []Rule) error {
	_, exist, err := s.GetRuleTemplateByName("all")
	if err != nil {
		return err
	}
	if !exist {
		t := &RuleTemplate{
			Name: "all",
			Desc: "default template for all rule",
		}
		if err := s.Save(t); err != nil {
			return err
		}
		return s.UpdateRuleTemplateRules(t, rules...)
	}
	return nil
}

const defaultAdminUser = "admin"

func (s *Storage) CreateAdminUser() error {
	_, exist, err := s.GetUserByName(defaultAdminUser)
	if err != nil {
		return err
	}
	if !exist {
		return s.Save(&User{
			Name:     defaultAdminUser,
			Password: "admin",
		})
	}
	return nil
}

func (s *Storage) Exist(model interface{}) (bool, error) {
	var count int
	err := s.db.Model(model).Where(model).Count(&count).Error
	if err != nil {
		return false, errors.New(errors.CONNECT_STORAGE_ERROR, err)
	}
	return count > 0, nil
}

func (s *Storage) Create(model interface{}) error {
	return errors.New(errors.CONNECT_STORAGE_ERROR, s.db.Create(model).Error)
}

func (s *Storage) Save(model interface{}) error {
	return errors.New(errors.CONNECT_STORAGE_ERROR, s.db.Save(model).Error)
}

func (s *Storage) Update(model interface{}, attrs ...interface{}) error {
	return errors.New(errors.CONNECT_STORAGE_ERROR, s.db.Model(model).UpdateColumns(attrs).Error)
}

func (s *Storage) Delete(model interface{}) error {
	return errors.New(errors.CONNECT_STORAGE_ERROR, s.db.Delete(model).Error)
}

func (s *Storage) HardDelete(model interface{}) error {
	return errors.New(errors.CONNECT_STORAGE_ERROR, s.db.Unscoped().Delete(model).Error)
}

func (s *Storage) TxExec(fn func(tx *sql.Tx) error) error {
	db := s.db.DB()
	tx, err := db.Begin()
	if err != nil {
		return errors.New(errors.CONNECT_STORAGE_ERROR, err)
	}
	err = fn(tx)
	if err != nil {
		tx.Rollback()
		return errors.New(errors.CONNECT_STORAGE_ERROR, err)
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return errors.New(errors.CONNECT_STORAGE_ERROR, err)
	}
	return nil
}

type RowList []string

func (r *RowList) Scan(src interface{}) error {
	var data string
	switch src := src.(type) {
	case nil:
		data = ""
	case string:
		data = src
	case []byte:
		data = string(src)
	default:
		return fmt.Errorf("scan: unable to scan type %T into []string", src)
	}
	*r = []string{}
	if data != "" {
		l := strings.Split(data, ",")
		for _, v := range l {
			if v != "" {
				*r = append(*r, v)
			}
		}
	}
	return nil
}
