package database

import (
	"strconv"
	"strings"
)

type TeamBelongModel struct {
	TeamUID         uint `json:"teamUID"`
	Score           uint `json:"score"`
	PercentComplete uint `json:"percentComplete"`
}

func (tb TeamsBelong) ToModel() (tbm []TeamBelongModel) {
	for _, str := range tb {
		strs := strings.Split(str, "|")
		teamUID, _ := strconv.Atoi(strs[0])
		score, _ := strconv.Atoi(strs[1])
		percentComplete, _ := strconv.Atoi(strs[2])
		tbm = append(tbm, TeamBelongModel{
			TeamUID:         uint(teamUID),
			Score:           uint(score),
			PercentComplete: uint(percentComplete),
		})
	}
	return
}
