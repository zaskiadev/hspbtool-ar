package db

import (
	"database/sql"
	f "fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

var (
	Server   = "10.201.48.11"
	Port     = 5000
	User     = "edpbintaro"
	Password = "sqledpbintaro123"
	Db       = "hspb_tool_ar"
)

/*func CheckDbConn(personName string, email string, birthday string, gender string, address string) {
	var err error

	ConnString := f.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;sslmode=disable", Server, User, Password, Port, Db)

	conn, err := sql.Open("sqlserver", ConnString)
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}
	f.Printf("Connected!\n")
	defer conn.Close()
	//option := 0
	//f.Println("0.GET \n1.INSERT \n2.UPDATE \n3.DELETE")
	//f.Scanln(&option)
	//switch option {
	//case 0:
	//	GetProducts(conn)
	//case 1:
	result, _ := CreateProduct(conn, personName, email, birthday, gender, address)

	f.Println(result)
	//case 2:
	//UpdateProduct(conn)
	//case 3:
	//	DeleteProduct(conn)
	//default:
	//	f.Println("Invalid operation request")
	//}
}*/

func Login(userName string, password string) (int, string) {
	var err error

	ConnString := f.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;sslmode=disable", Server, User, Password, Port, Db)

	conn, err := sql.Open("sqlserver", ConnString)
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}
	f.Printf("Connected!\n")
	defer conn.Close()
	var levelUser string

	getProduct_sql := f.Sprintf("select * from data_user where user_name='%s' and password='%s'", userName, password)
	f.Println(getProduct_sql)
	rows, err := conn.Query(getProduct_sql)
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
		}
		//f.Printf("ID: %d, Name: %s, Price: %f\n", id, name, price)
		count++
	}
	return count, levelUser
}
