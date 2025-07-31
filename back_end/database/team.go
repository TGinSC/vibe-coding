package database

import (
	"database/sql/driver"
	"encoding/json"
)

type Members []string

type Items []string

type TeamModel struct {
	TeamUID        uint    `grom:"unique;primarykey" json:"teamUID"`
	TeamLeader     uint    `json:"teamLeader"`
	TeamPassword   uint    `json:"teamPassword"`
	MembersInclude Members `json:"membersInclude"`
	ItemsInclude   Items   `json:"itemsInclude"`
}

func NewTeamModel() *TeamModel {
	return &TeamModel{}
}

func (*TeamModel) Get(uid uint) (res TeamModel, err error) {
	err = database.Model(&TeamModel{}).First(&res, uid).Error
	return
}

func (*TeamModel) Create(item *TeamModel) error {
	return database.Model(&TeamModel{}).Create(item).Error
}

func (*TeamModel) Delete(id uint) error {
	return database.Delete(&TeamModel{}, id).Error
}

func (*TeamModel) Updata(item *TeamModel) error {
	return database.Model(&TeamModel{}).Where("team_uid = ?", item.TeamUID).Updates(item).Error
}

func (*TeamModel) TableName() string {
	return "team"
}

func (t *Members) Scan(value interface{}) error {
	bytesValue, _ := value.([]byte)
	return json.Unmarshal(bytesValue, t)
}
func (t Members) Value() (driver.Value, error) {
	return json.Marshal(t)
}

func (t *Items) Scan(value interface{}) error {
	bytesValue, _ := value.([]byte)
	return json.Unmarshal(bytesValue, t)
}
func (t Items) Value() (driver.Value, error) {
	return json.Marshal(t)
}
