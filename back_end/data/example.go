package data

import (
	"log"
)

// test create
func StoreExampleData() {
	var err error

	// user
	user := User{
		UserUID:      1,
		UserPassword: 1,
		TeamsBelong: TeamBelongs{
			TeamBelong{
				TeamUID:         1,
				Score:           1,
				PercentComplete: 1,
			},
		},
		Messions: Messions{1},
		TeamsOwn: TeamOwns{1},
	}

	err = NewUser().Create(&user)
	if err != nil {
		panic(err)
	}
	log.Println("create user.")

	// team
	team := Team{
		TeamUID:        1,
		TeamLeader:     1,
		TeamPassword:   1,
		MembersInclude: Members{1},
		ItemsInclude:   Items{1},
	}

	err = NewTeam().Create(&team)
	if err != nil {
		panic(err)
	}
	log.Println("create teams.")

	// item
	item := Item{
		ItemUID:    1,
		Score:      1,
		ShouldBCB:  1,
		BCB:        1,
		IsComplete: false,
	}

	err = NewItem().Create(&item)
	if err != nil {
		panic(err)
	}
	log.Println("create item.")

	log.Println("create finish.")
}

// test get
func GetExampleData() {

	// user
	user, err := NewUser().Get(uint(1))
	if err != nil {
		panic(err)
	}
	log.Println("test User get:")
	log.Println(user)

	// team
	team, err := NewUser().Get(uint(1))
	if err != nil {
		panic(err)
	}
	log.Println("test Team get:")
	log.Println(team)

	// item
	item, err := NewItem().Get(uint(1))
	if err != nil {
		panic(err)
	}
	log.Println("test Item get:")
	log.Println(item)

	log.Println("test get over.")
}

// test updata
func UpdataExampleData() {
	var err error

	// user
	user := User{
		UserUID:      1,
		UserPassword: 2,
		TeamsBelong: TeamBelongs{
			TeamBelong{
				TeamUID:         1,
				Score:           1,
				PercentComplete: 1,
			},
			TeamBelong{
				TeamUID:         2,
				Score:           1,
				PercentComplete: 1,
			},
		},
		Messions: Messions{1, 2},
		TeamsOwn: TeamOwns{1, 2},
	}

	err = NewUser().Updata(&user)
	if err != nil {
		panic(err)
	}
	log.Println("updata User.")

	// team
	team := Team{
		TeamUID:        1,
		TeamLeader:     1,
		TeamPassword:   2,
		MembersInclude: Members{1, 2},
		ItemsInclude:   Items{1, 2},
	}

	err = NewTeam().Updata(&team)
	if err != nil {
		panic(err)
	}
	log.Println("updata Team.")

	// item
	item := Item{
		ItemUID:    1,
		Score:      2,
		ShouldBCB:  1,
		BCB:        1,
		IsComplete: true,
	}

	err = NewItem().Updata(&item)
	if err != nil {
		panic(err)
	}
	log.Println("updata Item.")

	log.Println("updata finish.")
}

// test delete
func DeleteExampleData() {
	var err error

	// user
	err = NewUser().Delete(uint(1))
	if err != nil {
		panic(err)
	}
	log.Println("delete User.")

	// team
	err = NewTeam().Delete(uint(1))
	if err != nil {
		panic(err)
	}
	log.Println("delete Team.")

	// item
	err = NewItem().Delete(uint(1))
	if err != nil {
		panic(err)
	}
	log.Println("delete Item.")

	log.Println("delete finish.")
}
