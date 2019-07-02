package factory

import (
	"echo_shop/app/models"
	"echo_shop/database"
	"echo_shop/pkg/utils"
	"fmt"
	"time"

	"github.com/Pallinder/go-randomdata"
	"github.com/bluele/factory-go/factory"
)

var (
	// 头像假数据
	avatars = []string{
		"https://iocaffcdn.phphub.org/uploads/images/201710/14/1/s5ehp11z6s.png",
		"https://iocaffcdn.phphub.org/uploads/images/201710/14/1/Lhd1SHqu86.png",
		"https://iocaffcdn.phphub.org/uploads/images/201710/14/1/LOnMrqbHJn.png",
		"https://iocaffcdn.phphub.org/uploads/images/201710/14/1/xAuDMxteQy.png",
		"https://iocaffcdn.phphub.org/uploads/images/201710/14/1/ZqM7iaP4CR.png",
		"https://iocaffcdn.phphub.org/uploads/images/201710/14/1/NDnzMutoxX.png",
	}
)

func userFactory(i int) *factory.Factory {
	now := time.Now()
	u := &models.User{
		Password:        "123456",
		EmailVerifiedAt: &now,
		Activated:       models.TrueTinyint,
	}
	r := utils.RandInt(0, len(avatars)-1)

	return factory.NewFactory(
		u,
	).Attr("Name", func(args factory.Args) (interface{}, error) {
		return fmt.Sprintf("user-%d", i+1), nil
	}).Attr("Avatar", func(args factory.Args) (interface{}, error) {
		return avatars[r], nil
	}).Attr("Email", func(args factory.Args) (interface{}, error) {
		if i == 0 {
			return "1@test.com", nil
		}
		if i == 1 {
			return "2@test.com", nil
		}
		return randomdata.Email(), nil
	})
}

// UsersTableSeeder : users table mock
func UsersTableSeeder() {
	dropAndCreateTable(&models.User{})

	for i := 0; i < 10; i++ {
		user := userFactory(i).MustCreate().(*models.User)
		if err := database.DBManager().Create(&user); err != nil {
			fmt.Printf("mock user error： %v\n", err)
		}
	}
}
