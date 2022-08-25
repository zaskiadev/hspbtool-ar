package db

import (
	"database/sql"
	"fmt"
	f "fmt"
	"time"
)

//Read Function
func GetProducts(db *sql.DB, userName string, Password string) (int, string, error) {
	var levelUser string
	getProduct_sql := f.Sprintf("select * from Products where user_name='%s' and password='%s'", userName, Password)

	rows, err := db.Query(getProduct_sql)
	if err != nil {
		f.Println("Error reading records: ", err.Error())
	}
	defer rows.Close()

	count := 0
	for rows.Next() {
		var name string
		var price float64
		var id int
		err := rows.Scan(&id, &name, &price)
		if err != nil {
			f.Println("Error reading rows: " + err.Error())
			return -1, "not found", err
		}
		f.Printf("ID: %d, Name: %s, Price: %f\n", id, name, price)
		count++
	}
	return count, levelUser, nil
}

//CreateFunction
func CreateProduct(db *sql.DB, personName string, email string, birthday string, gender string, address string) (int64, error) {
	//var name string
	//f.Print("Please enter your product name: ")
	//f.Scanln(&name)

	//var price float64
	//f.Print("Please enter your product's price: ")
	//f.Scanln(&price)
	currentTime := time.Now()
	insertProduct_sql := f.Sprintf("INSERT INTO person_data (person_name, email, birthday, gender, address,datesubmit) "+
		"VALUES ('%s' , '%s', '%s', '%s', '%s', '%s' ) ", personName, email, birthday, gender, address, currentTime.Format("2006-01-02 15:04:05"))

	fmt.Println(insertProduct_sql)
	rows, err := db.Query(insertProduct_sql)
	if err != nil {
		f.Println("Error occured while inserting a record", err.Error())
		return -1, err
	}

	defer rows.Close()
	//var lastInsertId1 int64
	//for rows.Next() {
	//	rows.Scan(&lastInsertId1)

	//}

	return -1, err
}

//function search
func InfoMsG(db *sql.DB, id int64) {
	infoQuery := f.Sprintf("Select name from Products where id=%d", id)
	rows, err := db.Query(infoQuery)
	if err != nil {
		f.Println("Error occured while giving info: ", err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		var name string
		var id = id
		err := rows.Scan(&name)
		if err != nil {
			f.Println("Error reading end process product id with, ", id, err)
		} else {
			f.Printf(name + " product has been created ")
		}

	}
}

//function update
func UpdateProduct(db *sql.DB) {
	f.Print("Please enter product id which you want to change: ")
	var id int
	f.Scanln(&id)
	f.Print("Please enter new product name ")
	var name string
	f.Scanln(&name)

	f.Print("Please enter new product'price ")
	var price float64
	f.Scanln(&price)

	update_query := f.Sprintf("UPDATE Products set name='%s', price=%f where id=%d", name, price, id)

	_, err := db.Exec(update_query)
	if err != nil {
		f.Println("Failed: " + err.Error())
	}
	f.Println("Product informations updated successfully")
}

//function delete
func DeleteProduct(db *sql.DB) {
	f.Print("Please enter product id which you want to delete: ")
	var id int
	f.Scanln(&id)
	delete_query := f.Sprintf("DELETE FROM Products where id=%d", id)
	_, err := db.Exec(delete_query)
	if err != nil {
		f.Println("Failed: " + err.Error())
	}
	f.Println("Product deleted successfully")
}
