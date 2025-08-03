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
type TeamOwns []uint

type User struct {
	UserUID      uint        `json:"userUID"`
	UserPassword uint        `json:"userPassword"`
	TeamsBelong  TeamBelongs `json:"teamsBelong"`
	Messions     Messions    `json:"messions"`
	TeamsOwn     TeamOwns    `json:"teamsOwn"`
}

func NewUser() *User {
	return &User{}
}

func (*User) Get(uid uint) (res User, err error) {
	__user__, err := database.NewUserModel().Get(uid)
	res = UserToUse(__user__)
	return
}

func (*User) Create(item *User) error {
	return database.NewUserModel().Create(item.ToStore())
}

func (*User) Delete(uid uint) error {
	return database.NewUserModel().Delete(uid)
}

func (*User) Updata(item *User) error {
	return database.NewUserModel().Updata(item.ToStore())
}

func (*User) DataName() string {
	return database.NewUserModel().TableName()
}

// 从参与运算的对象转为存储形式

// TeamBelongs 从参与运算的对象转为存储形式
func (__tb__ TeamBelongs) ToStore() (tb database.TeamsBelong) {
	for _, item := range __tb__ {
		str := fmt.Sprintf("%d|%d|%d", item.TeamUID, item.Score, item.PercentComplete)
		tb = append(tb, str)
	}
	return
}

// Messions 从参与运算的对象转为存储形式
func (__messions__ Messions) ToStore() (messions database.Messions) {
	for _, item := range __messions__ {
		str := strconv.Itoa(int(item))
		messions = append(messions, str)
	}
	return
}

// Teamsowns 从参与运算的对象转为存储形式
func (__teamOwns__ TeamOwns) ToStore() (teamsOwns database.TeamsOwn) {
	for _, item := range __teamOwns__ {
		str := strconv.Itoa(int(item))
		teamsOwns = append(teamsOwns, str)
	}
	return
}

// User 从参与运算的对象转为存储形式
func (user User) ToStore() *database.UserModel {
	return &database.UserModel{
		UserUID:      user.UserUID,
		UserPassword: user.UserPassword,
		TeamsBelong:  user.TeamsBelong.ToStore(),
		Messions:     user.Messions.ToStore(),
		TeamsOwn:     user.TeamsOwn.ToStore(),
	}
}
