package models

import (
	"echo_shop/pkg/serializer"
)

const (
	// ProductTableName : products table name
	ProductTableName = "products"
)

// Product 商品 Model
type Product struct {
	BaseModel
	Title       string  // 商品名称
	Description string  `gorm:"type:text"` // 商品详情
	Image       string  // 商品封面图片文件路径
	OnSale      uint    `gorm:"type:tinyint(1);default:1"` // 商品是否正在售卖
	Rating      float64 `gorm:"default:5"`                 // 商品平均评分
	SoldCount   uint    `gorm:"default:0"`                 // 销量
	ReviewCount uint    `gorm:"default:0"`                 // 评价数量
	Price       float64 `gorm:"type:decimal"`              // SKU(库存量单位) 最低价格

	Skus []ProductSku
}

// TableName 表名
func (Product) TableName() string {
	return ProductTableName
}

// Serialize -
func (p *Product) Serialize() serializer.Data {
	return serializer.Data{
		"id":           p.ID,
		"title":        p.Title,
		"description":  p.Description,
		"image":        p.Image,
		"on_sale":      TinyBool(p.OnSale),
		"rating":       p.Rating,
		"sold_count":   p.SoldCount,
		"review_count": p.ReviewCount,
		"price":        p.Price,
	}
}
