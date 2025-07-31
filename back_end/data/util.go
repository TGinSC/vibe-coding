package data

import (
	"contribution/database"
	"strconv"
	"strings"
)

// 从存储转换为可参与运算的对象
func ToModel(tb database.TeamsBelong) (tbm []TeamBelong) {
	for _, str := range tb {
		strs := strings.Split(str, "|")
		teamUID, _ := strconv.Atoi(strs[0])
		score, _ := strconv.Atoi(strs[1])
		percentComplete, _ := strconv.Atoi(strs[2])
		tbm = append(tbm, TeamBelong{
			TeamUID:         uint(teamUID),
			Score:           uint(score),
			PercentComplete: uint(percentComplete),
		})
	}
	return
}
