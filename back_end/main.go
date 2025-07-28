package main

import "contribution/database"

func main() {
	database.Open("data.db")
	database.NewUserModel().Create(&database.UserModel{
		UserUID:      1,
		UserPassWord: 1,
		TeamsBelong: database.TeamsBelong{
			"teamUID|score|percentComplate",
			"teamUID|score|percentComplate",
		},
		Messions: database.Messions{"1", "2"},
		TeamsOwn: database.TeamsOwn{"1", "2"},
	})
	database.NewTeamModel().Create(&database.TeamModel{
		TeamUID:        1,
		TeamLeader:     1,
		TeamPassword:   1,
		MembersInclude: database.Members{"123"},
		ItemsInclude:   database.Items{"123"},
	})
	database.NewItemModel().Create(&database.ItemModel{
		ItemUID:    1,
		Score:      1,
		ShouldBCB:  1,
		BCB:        1,
		IsComplete: false,
	})
}
