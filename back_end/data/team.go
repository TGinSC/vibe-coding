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

func (members Members) ToStore() (members__ database.Members) {
	for _, item := range members {
		str := strconv.Itoa(int(item))
		members__ = append(members__, str)
	}
	return
}

func (items Items) ToStore() (items__ database.Items) {
	for _, item := range items {
		str := strconv.Itoa(int(item))
		items__ = append(items__, str)
	}
	return
}

func (team Team) ToStore() *database.TeamModel {
	return &database.TeamModel{
		TeamUID:        team.TeamUID,
		TeamLeader:     team.TeamLeader,
		TeamPassword:   team.TeamPassword,
		MembersInclude: team.MembersInclude.ToStore(),
		ItemsInclude:   team.ItemsInclude.ToStore(),
	}
}
