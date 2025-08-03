package data

import (
	"contribution/database"
	"strconv"
	"strings"
)

// 从存储转换为可参与运算的对象

// TeamBelong 从存储转换为可参与运算的对象
func TeamsBelongToUse(__tb__ database.TeamsBelong) (tb TeamBelongs) {
	for _, str := range __tb__ {
		strs := strings.Split(str, "|")
		teamUID, _ := strconv.Atoi(strs[0])
		score, _ := strconv.Atoi(strs[1])
		percentComplete, _ := strconv.Atoi(strs[2])
		tb = append(tb, TeamBelong{
			TeamUID:         uint(teamUID),
			Score:           uint(score),
			PercentComplete: uint(percentComplete),
		})
	}
	return
}

// Messions 从存储转换为可参与运算的对象
func MessionsToUse(__m__ database.Messions) (m Messions) {
	for _, str := range __m__ {
		mession, _ := strconv.Atoi(str)
		m = append(m, uint(mession))
	}
	return
}

// TeamsOwn 从存储转换为可参与运算的对象
func TeamsOwnToUse(__to__ database.TeamsOwn) (to TeamOwns) {
	for _, str := range __to__ {
		teamown, _ := strconv.Atoi(str)
		to = append(to, uint(teamown))
	}
	return
}

// User 从存储转换为可参与运算的对象
func UserToUse(__user__ database.UserModel) (user User) {
	user = User{
		UserUID:      __user__.UserUID,
		UserPassword: __user__.UserPassword,
		TeamsBelong:  TeamsBelongToUse(__user__.TeamsBelong),
		Messions:     MessionsToUse(__user__.Messions),
		TeamsOwn:     TeamsOwnToUse(__user__.TeamsOwn),
	}
	return
}

// Members 从存储转换为可参与运算的对象
func MembersToUse(__m__ database.Members) (m Members) {
	for _, str := range __m__ {
		member, _ := strconv.Atoi(str)
		m = append(m, uint(member))
	}
	return
}

// Items 从存储转换为可参与运算的对象
func ItemsToUse(__items__ database.Items) (items Items) {
	for _, str := range __items__ {
		item, _ := strconv.Atoi(str)
		items = append(items, uint(item))
	}
	return
}

// Team 从存储转换为可参与运算的对象
func TeamToUse(__team__ database.TeamModel) (team Team) {
	team = Team{
		TeamUID:        __team__.TeamUID,
		TeamLeader:     __team__.TeamLeader,
		TeamPassword:   __team__.TeamPassword,
		MembersInclude: MembersToUse(__team__.MembersInclude),
		ItemsInclude:   ItemsToUse(__team__.ItemsInclude),
	}
	return
}

// Item 从存储转换为可参与运算的对象
func ItemToUse(__item__ database.ItemModel) (item Item) {
	item = Item{
		ItemUID:    __item__.ItemUID,
		Score:      __item__.Score,
		ShouldBCB:  ShouldBCB(__item__.ShouldBCB),
		BCB:        BCB(__item__.BCB),
		IsComplete: __item__.IsComplete,
	}
	return
}
