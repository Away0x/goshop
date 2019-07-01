package factory

import (
	"echo_shop/app/models"
	"echo_shop/database"
	"echo_shop/pkg/utils"
	"fmt"

	"github.com/Pallinder/go-randomdata"
)

// UserAddressesTableSeeder user_addresses table mock
func UserAddressesTableSeeder() {
	var users []*models.User
	database.DBManager().Find(&users)

	addresses := [][]string{
		{"北京市", "市辖区", "东城区"},
		{"河北省", "石家庄市", "长安区"},
		{"江苏省", "南京市", "浦口区"},
		{"江苏省", "苏州市", "相城区"},
		{"广东省", "深圳市", "福田区"},
	}

	for i, u := range users {
		address := addresses[utils.RandInt(0, len(addresses)-1)]
		for j := 0; j < 3; j++ {
			ud := &models.UserAddress{
				UserID:       u.ID,
				Province:     address[0],
				City:         address[1],
				District:     address[2],
				Address:      fmt.Sprintf("第%d街道第%d号", utils.RandInt(1, 10), utils.RandInt(1, 100)),
				Zip:          uint(10000 + i + j),
				ContactName:  randomdata.FullName(randomdata.Male),
				ContactPhone: randomdata.PhoneNumber(),
			}

			database.DBManager().Create(&ud)
		}
	}
}
