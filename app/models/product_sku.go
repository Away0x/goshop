package models

const (
	// ProductSkuTableName : product_skus table name
	ProductSkuTableName = "product_skus"
)

// ProductSku 商品SKU Model
type ProductSku struct {
	BaseModel
	Title       string // SKU 名称
	Description string // SKU 描述
	Price       int    `gorm:"type:decimal"` // SKU 价格
	Stock       uint   // 库存

	Product   Product
	ProductID uint `sql:"not null; index"` // 所属商品 id
}

// TableName 表名
func (ProductSku) TableName() string {
	return ProductSkuTableName
}
