package database

type ShouldBeCompletedBy uint
type ShouldBCB ShouldBeCompletedBy
type BCB uint

type ItemModel struct {
	ItemUID    uint      `grom:"unique;primarykey" json:"itemUID"`
	Score      uint      `json:"score"`
	ShouldBCB  ShouldBCB `json:"shouldBeCompletedBy"`
	BCB        BCB       `json:"beCompletedBy"`
	IsComplete bool      `json:"isComplete"`
}

func NewItemModel() *ItemModel {
	return &ItemModel{}
}

func (*ItemModel) Get(uid uint) (res ItemModel, err error) {
	err = database.Model(&TeamModel{}).First(&res, uid).Error
	return
}

func (*ItemModel) Create(item *ItemModel) error {
	return database.Model(&ItemModel{}).Create(item).Error
}

func (*ItemModel) Delete(id uint) error {
	return database.Delete(&ItemModel{}, id).Error
}

func (*ItemModel) Updata(item *ItemModel) error {
	return database.Model(&ItemModel{}).Where("item_uid = ?", item.ItemUID).Updates(item).Error
}

func (*ItemModel) TableName() string {
	return "item"
}
