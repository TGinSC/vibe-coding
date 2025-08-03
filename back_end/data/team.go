package data

import (
	"contribution/database"
	"strconv"
)

type Members []uint
type Items []uint

type Team struct {
	TeamUID        uint    `json:"teamUID"`
	TeamLeader     uint    `json:"teamLeader"`
	TeamPassword   uint    `json:"teamPassword"`
	MembersInclude Members `json:"membersInclude"`
	ItemsInclude   Items   `json:"itemsInclude"`
}

func NewTeam() *Team {
	return &Team{}
}

func (*Team) Get(uid uint) (res Team, err error) {
	team, err := database.NewTeamModel().Get(uid)
	res = TeamToUse(team)
	return
}

func (*Team) Create(item *Team) error {
	return database.NewTeamModel().Create(item.ToStore())
}

func (*Team) Delete(uid uint) error {
	return database.NewTeamModel().Delete(uid)
}

func (*Team) Updata(item *Team) error {
	return database.NewTeamModel().Updata(item.ToStore())
}

func (*Team) DataName() string {
	return database.NewTeamModel().TableName()
}

// 从参与运算的对象转为存储形式

// Members 从参与运算的对象转为存储形式
func (members Members) ToStore() (members__ database.Members) {
	for _, item := range members {
		str := strconv.Itoa(int(item))
		members__ = append(members__, str)
	}
	return
}

// Items 从参与运算的对象转为存储形式
func (items Items) ToStore() (items__ database.Items) {
	for _, item := range items {
		str := strconv.Itoa(int(item))
		items__ = append(items__, str)
	}
	return
}

// Team 从参与运算的对象转为存储形式
func (team Team) ToStore() *database.TeamModel {
	return &database.TeamModel{
		TeamUID:        team.TeamUID,
		TeamLeader:     team.TeamLeader,
		TeamPassword:   team.TeamPassword,
		MembersInclude: team.MembersInclude.ToStore(),
		ItemsInclude:   team.ItemsInclude.ToStore(),
	}
}
