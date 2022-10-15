package main

import (
	"Assignment2/httpserver"
	"log"
	"net/http"
)



func main() {
	r := httpserver.Router()
	log.Fatal(http.ListenAndServe(":4000", r))
	// db, err := config.ConnectPostgres()
	// if err != nil {
	// 	panic(err)
	// }
	// studentRepo := postgres.NewStudentRepo(db)
	// students, err := studentRepo.GetStudents()
	// if err != nil {
	// 	panic(err)
	// }

	// for _, s := range *students {
	// 	fmt.Println(s)
	// }

	// order = append(order, 
	// 	TheOrder{
	// 		Order_ID: 1,
	// 		Costumer_Name: "salahudin benar",
	// 		Description: "Iphone X",
	// 		Item_ID: 1,
	// 		Item_Code: 101,
	// 		Quantity: 2,
	// 		Ordered_At: 01-02-2015,
	// 	},
	// 	TheOrder{
	// 		Order_ID: 2,
	// 		Costumer_Name: "benarudin salah",
	// 		Description: "Iphone X",
	// 		Item_ID: 2,
	// 		Item_Code: 111,
	// 		Quantity: 1,
	// 		Ordered_At: 06-02-2023,
	// 	},
	// 	TheOrder{
	// 		Order_ID: 3,
	// 		Costumer_Name: "salahudin salah",
	// 		Description: "Iphone 12",
	// 		Item_ID: 3,
	// 		Item_Code: 121,
	// 		Quantity: 5,
	// 		Ordered_At: 07-06-2011,
	// 	},
	// )
}
