package data

import (
	"contribution/database"
	"fmt"
	"strconv"
)

type TeamBelong struct {
	TeamUID         uint `json:"teamUID"`
	Score           uint `json:"score"`
	PercentComplete uint `json:"percentComplete"`
}

type TeamBelongs []TeamBelong
type Messions []uint
type TeamsOwns []uint

type User struct {
	UserUID      uint        `json:"userUID"`
	UserPassWord uint        `json:"userPassword"`
	TeamsBelong  TeamBelongs `json:"teamsBelong"`
	Messions     Messions    `json:"messions"`
	TeamsOwn     TeamsOwns   `json:"teamsOwn"`
}

// 从参与运算的对象转为存储形式
func (tbm TeamBelongs) ToStore() (tb database.TeamsBelong) {
	for _, item := range tbm {
		str := fmt.Sprintf("%d|%d|%d", item.TeamUID, item.Score, item.PercentComplete)
		tb = append(tb, str)
	}
	return
}

func (messions Messions) ToStore() (messions__ database.Messions) {
	for _, item := range messions {
		str := strconv.Itoa(int(item))
		messions__ = append(messions__, str)
	}
	return
}

func (teamsOwns TeamsOwns) ToStore() (teamsOwns__ database.TeamsOwn) {
	for _, item := range teamsOwns {
		str := strconv.Itoa(int(item))
		teamsOwns__ = append(teamsOwns__, str)
	}
	return
}

func (user User) ToStore() *database.UserModel {
	return &database.UserModel{
		UserUID:      user.UserUID,
		UserPassWord: user.UserPassWord,
		TeamsBelong:  user.TeamsBelong.ToStore(),
		Messions:     user.Messions.ToStore(),
		TeamsOwn:     user.TeamsOwn.ToStore(),
	}
}
