package factory

import (
	"echo_shop/bootstrap"
	"echo_shop/database"
)

func dropAndCreateTable(table interface{}) {
	database.DBManager().DropTable(table)
	database.DBManager().CreateTable(table)
}

// Run run mock
func Run() {
	db, _ := bootstrap.SetupDB()
	defer db.Close()

	UsersTableSeeder()
	UserAddressesTableSeeder()
}
