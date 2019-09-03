package product

import (
	"echo_shop/app/context"
	"echo_shop/app/models"
	"echo_shop/pkg/serializer"
)

// Index 商品列表
func Index(c *context.AppContext) error {
	products := make([]*models.Product, 0)
	if err := models.Where("on_sale = ?", models.TrueTinyint).Find(&products).Error; err != nil {
		return err
	}

	return c.RenderHTML("products/index",
		serializer.Wrap("products", products))
	// return c.RenderJSON(serializer.Wrap("products", products))
}
