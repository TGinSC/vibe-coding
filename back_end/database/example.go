package database

import (
	"log"
)

// test create
func StoreExampleData() {

	var err error

	// test UserModel create
	err = NewUserModel().Create(&UserModel{
		UserUID:      1,
		UserPassword: 1,
		TeamsBelong:  TeamsBelong{"1|1|1", "1|1|1s"},
		Messions:     Messions{"1"},
		TeamsOwn:     TeamsOwn{"1"},
	})
	if err != nil {
		panic(err)
	}
	log.Println("create UserModel.")

	// test TeamModel create
	err = NewTeamModel().Create(&TeamModel{
		TeamUID:        1,
		TeamLeader:     1,
		TeamPassword:   1,
		MembersInclude: Members{"1"},
		ItemsInclude:   Items{"1"},
	})
	if err != nil {
		panic(err)
	}
	log.Println("create TeamModel.")

	// test ItemModel create
	err = NewItemModel().Create(&ItemModel{
		ItemUID:    1,
		Score:      1,
		ShouldBCB:  1,
		BCB:        1,
		IsComplete: false,
	})
	if err != nil {
		panic(err)
	}
	log.Println("create ItemModel.")

	log.Println("create finish.")
}

// test get
func GetExampleData() {

	// test UserModel get
	user, err := NewUserModel().Get(uint(1))
	if err != nil {
		panic(err)
	}
	log.Println("test UserModel get:")
	log.Println(user)

	// test TeamModel get
	team, err := NewTeamModel().Get(uint(1))
	if err != nil {
		panic(err)
	}
	log.Println("test TeamModel get:")
	log.Println(team)

	// test ItemModel get
	item, err := NewItemModel().Get(uint(1))
	if err != nil {
		panic(err)
	}
	log.Println("test ItemModel get:")
	log.Println(item)
}

// test updata
func UpdataExampleData() {
	var err error

	// test UserModel updata
	err = NewUserModel().Updata(&UserModel{
		UserUID:      1,
		UserPassword: 2,
	})
	if err != nil {
		panic(err)
	}
	log.Println("updata UserModel.")

	// test TeamModel updata
	err = NewTeamModel().Updata(&TeamModel{
		TeamUID:      1,
		TeamPassword: 2,
	})
	if err != nil {
		panic(err)
	}
	log.Println("updata TeamModel.")

	// test ItemModel updata
	err = NewItemModel().Updata(&ItemModel{
		ItemUID:    1,
		IsComplete: true,
	})
	if err != nil {
		panic(err)
	}
	log.Println("updata ItemModel.")

	log.Println("updata finish.")
}

// test delete
func DeleteExampleData() {
	var err error

	// test UserModel delete
	err = NewUserModel().Delete(uint(1))
	if err != nil {
		panic(err)
	}
	log.Println("delete UserModel")

	// test TeamModel delete
	err = NewTeamModel().Delete(uint(1))
	if err != nil {
		panic(err)
	}
	log.Println("delete TeamModel")

	// test ItemModel delete
	err = NewItemModel().Delete(uint(1))
	if err != nil {
		panic(err)
	}
	log.Println("delete ItemModel")

	log.Println("delete finish")
}
