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

func NewItem() *Item {
	return &Item{}
}

func (*Item) Get(uid uint) (res Item, err error) {
	item, err := database.NewItemModel().Get(uid)
	res = ItemToUse(item)
	return
}

func (*Item) Create(item *Item) error {
	return database.NewItemModel().Create(item.ToStore())
}

func (*Item) Delete(uid uint) error {
	return database.NewItemModel().Delete(uid)
}

func (*Item) Updata(item *Item) error {
	return database.NewItemModel().Updata(item.ToStore())
}

// 从参与运算的对象转为存储形式

// Item 从参与运算的对象转为存储形式
func (item Item) ToStore() *database.ItemModel {
	return &database.ItemModel{
		ItemUID:    item.ItemUID,
		Score:      item.Score,
		ShouldBCB:  database.ShouldBCB(item.ShouldBCB),
		BCB:        database.BCB(item.BCB),
		IsComplete: item.IsComplete,
	}
}
