package database

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

// 包括 teamUID(uint), score(uint), percentComplete(uint)
// 格式 "teamUID|score|percentComplate"
// 新建一个 TeamBelongModel 来实现转换
type TeamsBelong []string

type Messions []string

type TeamsOwn []string

type UserModel struct {
	UserUID      uint        `gorm:"unique;primarykey" json:"userUID"`
	UserPassWord uint        `json:"userPassword"`
	TeamsBelong  TeamsBelong `json:"teamBelong"`
	Messions     Messions    `json:"messions"`
	TeamsOwn     TeamsOwn    `json:"teamOwn"`
}

func NewUserModel() *UserModel {
	return &UserModel{}
}

func (*UserModel) Get(uid uint) (res UserModel, err error) {
	err = database.Model(&UserModel{}).First(&res, uid).Error
	return
}

func (*UserModel) Create(item *UserModel) error {
	return database.Model(&UserModel{}).Create(item).Error
}

func (*UserModel) Delete(id uint) error {
	return database.Delete(&UserModel{}, id).Error
}

func (*UserModel) Updata(item *UserModel) error {
	return database.Model(&UserModel{}).Where("uid = ?", item.UserUID).Updates(item).Error
}

func (*UserModel) TableName() string {
	return "user"
}

// 创建新的 TeamBelong
func CreateTeamBelong(teamUID uint, score uint, percentComplate uint) string {
	str := fmt.Sprintf("%d|%d|%d", teamUID, score, percentComplate)
	return str
}

// 重载函数，实现 string 数组写入

func (t *TeamsBelong) Scan(value interface{}) error {
	bytesValue, _ := value.([]byte)
	return json.Unmarshal(bytesValue, t)
}
func (t TeamsBelong) Value() (driver.Value, error) {
	return json.Marshal(t)
}

func (t *Messions) Scan(value interface{}) error {
	bytesValue, _ := value.([]byte)
	return json.Unmarshal(bytesValue, t)
}
func (t Messions) Value() (driver.Value, error) {
	return json.Marshal(t)
}

func (t *TeamsOwn) Scan(value interface{}) error {
	bytesValue, _ := value.([]byte)
	return json.Unmarshal(bytesValue, t)
}
func (t TeamsOwn) Value() (driver.Value, error) {
	return json.Marshal(t)
}
