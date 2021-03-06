// @File: type
// @Date: 2020/12/10 16:40
// @Author: 安红豆
// @Description: 类型

package model

import "gorm.io/gorm"

//类型属性
type Type struct {
	Articles []Article `json:"articles" gorm:"one2many:article;foreignKey:TypeId"` //类型下的文章
	gorm.Model
	TypeName  string `json:"typeName" gorm:"type:varchar(100);not null"` //类型名字
	IsDeleted int    `json:"isDeleted" gorm:"int"`                       //是否删除（1删除、2正常）
}

//分页查询的属性
type TypePage struct {
	PageNum  int //第几页
	PageSize int //每页数量
	Type
}

//实现gorm的接口，重命名表名
func (Type) TableName() string {
	return "type"
}
