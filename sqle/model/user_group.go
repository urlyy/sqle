package model

import "github.com/jinzhu/gorm"

type UserGroup struct {
	Model
	Name  string  `json:"name" gorm:"index"`
	Desc  string  `json:"desc" gorm:"column:description"`
	Users []*User `gorm:"many2many:user_group_users"`
	Stat  uint    `json:"stat" gorm:"comment:'0:active,1:disabled'"`
	Roles []*Role `gorm:"many2many:user_group_roles"`
}

func (ug *UserGroup) TableName() string {
	return "user_groups"
}

func (ug *UserGroup) SetStat(stat int) {
	ug.Stat = uint(stat)
}

func (ug *UserGroup) IsDisabled() bool {
	return ug.Stat == Disabled
}

func (s *Storage) GetUserGroupByName(name string) (
	userGroup *UserGroup, isExist bool, err error) {
	userGroup = &UserGroup{}

	err = s.db.Where("name = ?", name).First(userGroup).Error
	if gorm.IsRecordNotFoundError(err) {
		return nil, false, nil
	}

	return userGroup, true, err
}
