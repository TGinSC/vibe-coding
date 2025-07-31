package data

import "contribution/database"

type ShouldBeCompletedBy uint
type ShouldBCB ShouldBeCompletedBy
type BCB uint

type Item struct {
	ItemUID    uint      `json:"itemUID"`
	Score      uint      `json:"score"`
	ShouldBCB  ShouldBCB `json:"shouldBeCompletedBy"`
	BCB        BCB       `json:"beCompletedBy"`
	IsComplete bool      `json:"isComplete"`
}

func (item Item) ToStore() *database.ItemModel {
	return &database.ItemModel{
		ItemUID:    item.ItemUID,
		Score:      item.Score,
		ShouldBCB:  database.ShouldBCB(item.ShouldBCB),
		BCB:        database.BCB(item.BCB),
		IsComplete: item.IsComplete,
	}
}
