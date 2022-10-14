package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Tag struct {
	Model
	Name       string `json:"name" gorm:"not null;unique"`
	CreatedBy  string `json:"created_by" gorm:"not null"`
	ModifiedBy string `json:"modified_by" gorm:"not null"`
	State      int    `json:"state" gorm:"default:1"`
}

func (tag *Tag) BeforeCreate(scope *gorm.Scope) error {
	err := scope.SetColumn("CreatedOn", time.Now().Unix())
	if err != nil {
		return err
	}
	return nil
}

func (tag *Tag) BeforeUpdate(scope *gorm.Scope) error {
	err := scope.SetColumn("ModifiedOn", time.Now().Unix())
	if err != nil {
		return err
	}
	return nil
}

/*
	1.我们创建了一个Tag struct{}，用于Gorm的使用。并给予了附属属性json，这样子在c.JSON的时候就会自动转换格式，非常的便利
	2.可能会有的初学者看到return，而后面没有跟着变量，会不理解；其实你可以看到在函数末端，我们已经显示声明了返回值，这个变量在函数体内也可以直接使用，因为他在一开始就被声明了
	3.有人会疑惑db是哪里来的；因为在同个models包下，因此db *gorm.DB是可以直接使用的
*/

func ExistTagByName(name string) bool {
	var tag Tag
	db.Select("id").Where("name = ?", name).First(&tag)
	if tag.ID > 0 {
		return true
	}
	return false
}

func ExistTagByID(id int) bool {
	var tag Tag
	db.Select("id").Where("id = ?", id).First(&tag)
	if tag.ID > 0 {
		return true
	}
	return false
}

func DeleteTag(id int) bool {
	db.Where("id = ?", id).Delete(&Tag{})
	return true
}

func AddTag(name string, state int, createdBy string) bool {
	db.Create(&Tag{
		Name:      name,
		State:     state,
		CreatedBy: createdBy,
	})
	return true
}

func EditTag(id int, data interface{}) bool {
	db.Model(&Tag{}).Where("id = ?", id).Updates(data)
	return true
}

/*
以上代码属于gorm的Callbacks，可以将回调方法定义为模型结构的指针，在创建、更新、查询、删除时将被调用，如果任何回调返回错误，gorm将停止未来操作并回滚所有更改。
gorm所支持的回调方法：
	创建：BeforeSave、BeforeCreate、AfterCreate、AfterSave
	更新：BeforeSave、BeforeUpdate、AfterUpdate、AfterSave
	删除：BeforeDelete、AfterDelete
	查询：AfterFind
*/

func GetTags(pageNum, PageSize int, maps interface{}) (tags []Tag) {
	db.Where(maps).Offset(pageNum).Limit(PageSize).Find(&tags)
	return
}

func GetTagTotal(maps interface{}) (count int) {
	db.Model(&Tag{}).Where(maps).Count(&count)
	return
}
