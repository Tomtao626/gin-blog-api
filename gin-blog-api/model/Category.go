package model

import (
	"gin-blog/utils/errmsg"
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	ID   uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"type:varchar(20);not null" json:"name"`
}

// CheckCategory 查询分类是否存在
func CheckCategory(name string) (code int) {
	var cate Category
	db.Select("id").Where("name = ?", name).First(&cate)
	if cate.ID > 0 {
		return errmsg.ERROR_CATENAME_USED // 1001
	}
	return errmsg.SUCCESS
}

// CheckUpCategory 更新查询
func CheckUpCategory(id int, name string) (code int) {
	var cate Category
	db.Select("id, name").Where("name = ?", name).First(&cate)
	if cate.ID == uint(id) {
		return errmsg.SUCCESS
	}
	if cate.ID > 0 {
		return errmsg.ERROR_USERNAME_USED //1001
	}
	return errmsg.ERROR
}

// CreateCategory 新增分类
func CreateCategory(data *Category) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR //500
	}
	return errmsg.SUCCESS //200
}

// GetCategorys 获取分类列表
func GetCategorys(pageSize int, pageNum int) ([]Category, int64) {
	var cate []Category
	var total int64
	err = db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&cate).Count(&total).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0
	}
	return cate, total
}

// EditCategory 编辑分类
func EditCategory(id int, data *Category) int {
	var cate Category
	var maps = make(map[string]interface{})
	maps["name"] = data.Name
	err = db.Model(&cate).Where("id = ? ", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// DeleteCategory 删除分类
func DeleteCategory(id int) int {
	var cate Category
	err = db.Where("id = ? ", id).Delete(&cate).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
