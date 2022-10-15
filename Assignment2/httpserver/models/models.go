package models

import (
	"Assignment2/config"
	"database/sql"
	"fmt"
	"time"

	"log"
)

// type Repo struct {
// 	db *sql.DB
// }

type TheOrder struct {
	Order_ID      uint		`json: "Order_ID" sql:"primary_key"`
	Costumer_Name string	`json: "Costumer_Name"`
	Ordered_At    time.Time	`json: "Ordered_At"`
	Items      	  int		`json: "Items" sql:"foreignkey:Order_ID"`
}

type Item struct {
	LineItem_ID   uint		`json: "LineItem_ID" sql:"primary_key"`
	Item_Code     string	`json: "Item_Code"`
	Description   string	`json: "Description"`
	Quantity      uint		`json: "Quantity"`
	Order_ID      uint		`json: "-"`
}

func GetTheOrder() ([]TheOrder, error)  {
	db := config.ConnectPostgres()

	query := `SELECT * FROM orders_by`
	rows, err := db.Query(query)
	if err != nil {
		log.Fatalf("tidak bisa mengeksekusi query. %v", err)
	}
	defer rows.Close()

	var orders []TheOrder

	for rows.Next() {
		order := TheOrder{}
		err := rows.Scan(
			&order.Order_ID,
			&order.Costumer_Name,
			&order.Ordered_At,
			&order.Items,
		)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	return orders, err
}

func GetOrder(Order_ID int) (TheOrder, error)  {
	db := config.ConnectPostgres()
	defer db.Close()

	var order TheOrder
	query := `SELECT * FROM orders_by WHERE Order_ID=$1`
	row := db.QueryRow(query, order.Order_ID)
	err := row.Scan(
		&order.Order_ID, 
		&order.Costumer_Name,
		&order.Ordered_At,
		&order.Items,
	)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("Tidak ada Data yang dicari")
		return order, nil

	case nil :
		return order, nil
	
	default:
		log.Fatalf("Tidak bisa Mengambil Data. %v", err)
	}

	return order, nil
}

func CreateOrder(order TheOrder) int64  {
	db := config.ConnectPostgres()
	defer db.Close()

	query := `INSERT INTO orders_by (Costumer_Name, Ordered_At, Items) VALUES ($1, $2, $3) RETURNING Order_ID`

	var Order_ID int64
	err := db.QueryRow(query,		
		order.Order_ID,
		order.Costumer_Name,
		order.Ordered_At,
		order.Items,
	).Scan(&Order_ID)
	if err != nil {
		log.Fatalf("Tidak Bisa mengeksekusi query. %v", err)
	}
	fmt.Printf("Insert data single record %v", Order_ID)

	return Order_ID
}

func UpdateOrder(Order_ID int64, order TheOrder) int64  {
	db := config.ConnectPostgres()
	defer db.Close()

	query := `UPDATE orders_by SET Costumer_Name=$2, Ordered_At=%3, Items=$4) WHERE Order_ID=$1`

	res, err := db.Exec(
		query, 
		Order_ID,
		order.Costumer_Name,
		order.Ordered_At,
		order.Items,
	)
	if err != nil {
		log.Fatalf("Tidak bisa mengeksekusi query. %v", err)
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error ketika mengecheck rows/data yang diupdate. %v", err)
	}

	fmt.Printf("Total rows/record yang diupdate %v\n", rowsAffected)

	return rowsAffected
}

func DeleteOrder(Order_ID int64) int64 {
	db := config.ConnectPostgres()
	defer db.Close()

	Query := `DELETE FROM orders_by WHERE Order_ID=$1`
	res, err := db.Exec(Query, Order_ID)
	if err != nil {
		log.Fatalf("tidak bisa mengeksekusi query. %v", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Fatalf("tidak bisa mencari data. %v", err)
	}
	fmt.Printf("Total data yang terhapus %v", rowsAffected)

	return rowsAffected
}