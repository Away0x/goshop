package factory

import (
	"echo_shop/app/models"
)

// ProductsTableSeeder : products table mock
func ProductsTableSeeder() {
	dropAndCreateTable(&models.Product{})

	// var (
	//   avatars = []string{
	//     "https://cdn.learnku.com/uploads/images/201806/01/5320/7kG1HekGK6.jpg",
	//     "https://cdn.learnku.com/uploads/images/201806/01/5320/1B3n0ATKrn.jpg",
	//     "https://cdn.learnku.com/uploads/images/201806/01/5320/r3BNRe4zXG.jpg",
	//     "https://cdn.learnku.com/uploads/images/201806/01/5320/C0bVuKB2nt.jpg",
	//     "https://cdn.learnku.com/uploads/images/201806/01/5320/82Wf2sg8gM.jpg",
	//     "https://cdn.learnku.com/uploads/images/201806/01/5320/nIvBAQO5Pj.jpg",
	//     "https://cdn.learnku.com/uploads/images/201806/01/5320/XrtIwzrxj7.jpg",
	//     "https://cdn.learnku.com/uploads/images/201806/01/5320/uYEHCJ1oRp.jpg",
	//     "https://cdn.learnku.com/uploads/images/201806/01/5320/2JMRaFwRpo.jpg",
	//     "https://cdn.learnku.com/uploads/images/201806/01/5320/pa7DrV43Mw.jpg",
	//   }
	// )

}
