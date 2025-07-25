package database

import (
	"database/sql/driver"
	"encoding/json"
)

// 包括 teamUID(uint), score(uint), percentComplete(uint)
// 格式 "teamUID|score|percentComplate"
// 新建一个 TeamBelongModel 来实现转换
type TeamBelong []string

type Mession []string

type TeamOwn []string

type UserModel struct {
	UserUID      uint       `gorm:"unique;primarykey" json:"userUID"`
	UserPassWord uint       `json:"userPassword"`
	TeamBelong   TeamBelong `json:"teamBelong"`
	Mession      Mession    `json:"mession"`
	TeamOwn      TeamOwn    `json:"teamOwn"`
}

func (t *TeamBelong) Scan(value interface{}) error {
	bytesValue, _ := value.([]byte)
	return json.Unmarshal(bytesValue, t)
}

func (t TeamBelong) Value() (driver.Value, error) {
	return json.Marshal(t)
}
func (t *Mession) Scan(value interface{}) error {
	bytesValue, _ := value.([]byte)
	return json.Unmarshal(bytesValue, t)
}

func (t Mession) Value() (driver.Value, error) {
	return json.Marshal(t)
}
func (t *TeamOwn) Scan(value interface{}) error {
	bytesValue, _ := value.([]byte)
	return json.Unmarshal(bytesValue, t)
}

func (t TeamOwn) Value() (driver.Value, error) {
	return json.Marshal(t)
}
