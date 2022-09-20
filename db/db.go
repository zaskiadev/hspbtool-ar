package db

import (
	"database/sql"
	f "fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"hpbtool-ar/models"
	"log"
	"time"
)

var (
	Server   = "10.201.48.11"
	Port     = 5000
	User     = "edpbintaro"
	Password = "sqledpbintaro123"
	Db       = "hspb_tool_ar"
)

func Login(userName string, password string) (int, string, string) {
	var err error

	ConnString := f.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;sslmode=disable", Server, User, Password, Port, Db)

	conn, err := sql.Open("sqlserver", ConnString)
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}
	f.Printf("Connected!\n")
	defer conn.Close()
	var levelUserReturn string

	getProduct_sql := f.Sprintf("select user_name,level_user, user_id from data_user where user_name='%s' and password='%s'", userName, password)
	f.Println(getProduct_sql)
	rows, err := conn.Query(getProduct_sql)
	if err != nil {
		f.Println("Error reading records: ", err.Error())
	}
	defer rows.Close()

	count := 0
	var userID string
	for rows.Next() {
		var user_name string
		var level_user string
		var user_id string
		//var id int
		err := rows.Scan(&user_name, &level_user, &user_id)
		if err != nil {
			f.Println("Error reading rows: " + err.Error())
		}

		//f.Printf("user name: %s, level user: %s", user_name, level_user)
		levelUserReturn = level_user
		userID = user_id
		count++
	}
	return count, levelUserReturn, userID
}

func AddTask(userDestinationTask string, taskID string, companyTask string, picCompanyTask string, salesCompanyTask string, dateDeadLine string, userCreatedTask string, taskNotes string) bool {
	ConnString := f.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;sslmode=disable", Server, User, Password, Port, Db)

	conn, err := sql.Open("sqlserver", ConnString)
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}
	f.Printf("Connected!\n")
	defer conn.Close()
	currentTime := time.Now()
	var success bool
	insertProduct_sql := f.Sprintf("INSERT INTO task (task_id, user_id_create_task, user_id_delegation_task, date_created_task,date_deadline_task,company_id_destination_task,sales_id_destination_task,pic_id_destination_task, task_notes) "+
		"VALUES ('%s' , '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s') ", taskID, userCreatedTask, userDestinationTask, currentTime.Format("2006-01-02 15:04:05"), dateDeadLine, companyTask, salesCompanyTask, picCompanyTask, taskNotes)
	rows, err := conn.Query(insertProduct_sql)
	println(insertProduct_sql)
	if err != nil {
		f.Println("Error occured while inserting a record", err.Error())
		return false
	}

	success = true
	defer rows.Close()

	return success
}

func GetDataTempAddTask() ([]models.TempUserTask, []models.TempCompanyTask, []models.TempPicTask) {

	var err error

	ConnString := f.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;sslmode=disable", Server, User, Password, Port, Db)

	conn, err := sql.Open("sqlserver", ConnString)
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}
	f.Printf("Connected!\n")
	defer conn.Close()

	var tempDataUserTask []models.TempUserTask
	getUserTask_sql := f.Sprintf("select user_id,user_name from data_user where level_user='ar'")
	f.Println(getUserTask_sql)
	rows, err := conn.Query(getUserTask_sql)
	if err != nil {
		f.Println("Error reading records: ", err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var user_id string
		var user_name string
		//var id int
		err := rows.Scan(&user_id, &user_name)
		if err != nil {
			f.Println("Error reading rows: " + err.Error())
		}
		userTask := models.TempUserTask{
			UserID:   user_id,
			UserName: user_name,
		}
		tempDataUserTask = append(tempDataUserTask, userTask)
	}

	var tempDataCompany []models.TempCompanyTask
	getCompanyTask_sql := f.Sprintf("select company_id,name,address,phone, company_id_pms from company")
	f.Println(getCompanyTask_sql)
	rows2, err2 := conn.Query(getCompanyTask_sql)
	if err2 != nil {
		f.Println("Error reading records: ", err2.Error())
	}
	defer rows2.Close()

	for rows2.Next() {
		var companyID string
		var name string
		var address string
		var phone string
		var companyPMSID string
		//var id int
		err2 := rows2.Scan(&companyID, &name, &address, &phone, &companyPMSID)
		if err2 != nil {
			f.Println("Error reading rows: " + err2.Error())
		}
		companyTask := models.TempCompanyTask{
			CompanyID:    companyID,
			CompanyPMSID: companyPMSID,
			CompanyName:  name,
			Address:      address,
			Phone:        phone,
		}
		tempDataCompany = append(tempDataCompany, companyTask)
	}

	var tempDataPIC []models.TempPicTask
	getPICTask_sql := f.Sprintf("select pic_id, name, company_id, identification_id from pic")
	f.Println(getPICTask_sql)
	rows3, err3 := conn.Query(getPICTask_sql)
	if err3 != nil {
		f.Println("Error reading records: ", err.Error())
	}
	defer rows3.Close()

	for rows3.Next() {
		var pic_id string
		var pic_name string
		var company_id string
		var identification_id string
		//var id int
		err3 := rows.Scan(&pic_id, &pic_name, &company_id, &identification_id)
		if err3 != nil {
			f.Println("Error reading rows: " + err.Error())
		}
		PicTask := models.TempPicTask{
			PicID:            pic_id,
			PicName:          pic_name,
			CompanyID:        company_id,
			IdentificationID: identification_id,
		}
		tempDataPIC = append(tempDataPIC, PicTask)
	}

	return tempDataUserTask, tempDataCompany, tempDataPIC

}
