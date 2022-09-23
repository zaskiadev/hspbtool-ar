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
	//status := "created"
	insertProduct_sql := f.Sprintf("INSERT INTO task (task_id, user_id_create_task, user_id_delegation_task, date_created_task,date_deadline_task,company_id_destination_task,sales_id_destination_task,pic_id_destination_task, task_notes,status_task) "+
		"VALUES ('%s' , '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s','%s') ", taskID, userCreatedTask, userDestinationTask, currentTime.Format("2006-01-02 15:04:05"), dateDeadLine, companyTask, salesCompanyTask, picCompanyTask, taskNotes, "created")
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

func GetDataTask() []models.DataTask {

	var err error

	ConnString := f.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;sslmode=disable", Server, User, Password, Port, Db)

	conn, err := sql.Open("sqlserver", ConnString)
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}
	f.Printf("Connected!\n")
	defer conn.Close()

	var tempDataTask []models.DataTask
	getUserTask_sql := f.Sprintf("select t.task_id, uc.user_name as user_create, ud.user_name as user_delegation, t.date_deadline_task,  " +
		"c.company_id_pms, c.name, c.address, p.name, p.phone, s.name , t.task_notes, t.status_task from task t " +
		"LEFT JOIN data_user uc ON t.user_id_create_task = uc.user_id " +
		"LEFT JOIN data_user ud ON t.user_id_delegation_task= ud.user_id " +
		"LEFT JOIN company c ON t.company_id_destination_task=c.company_id " +
		"LEFT JOIN pic p ON t.pic_id_destination_task=p.pic_id " +
		"LEFT JOIN sales s ON t.sales_id_destination_task=s.sales_id ",
	)
	f.Println(getUserTask_sql)
	rows, err := conn.Query(getUserTask_sql)
	if err != nil {
		f.Println("Error reading records: ", err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var user_created_task string
		var assign_to string
		var deadline string
		var company string
		var company_id_pms string
		var company_alamat string
		var pic string
		var pic_phone string
		var sales string
		var task_notes string
		var status_task string
		var task_id string
		//var id int
		err := rows.Scan(&task_id, &user_created_task, &assign_to, &deadline, &company_id_pms, &company, &company_alamat, &pic, &pic_phone, &sales, &task_notes, &status_task)
		if err != nil {
			f.Println("Error reading rows: " + err.Error())
		}
		dataTask := models.DataTask{
			UserCreatedTask: user_created_task,
			AssignTask:      assign_to,
			Deadline:        deadline,
			Company:         company_id_pms + "-" + company + "" + company_alamat,
			PIC:             pic + "-" + pic_phone,
			Sales:           sales,
			TaskNotes:       task_notes,
			StatusTask:      status_task,
			CodeTask:        task_id,
		}
		tempDataTask = append(tempDataTask, dataTask)
	}

	return tempDataTask

}
func GetDataTaskForCommentTask(taskCode string) models.DataTask {

	var err error

	ConnString := f.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;sslmode=disable", Server, User, Password, Port, Db)

	conn, err := sql.Open("sqlserver", ConnString)
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}
	f.Printf("Connected!\n")
	defer conn.Close()

	var tempDataTask models.DataTask
	getUserTask_sql := f.Sprintf("select t.task_id, uc.user_name as user_create, ud.user_name as user_delegation, t.date_deadline_task,  c.company_id_pms, c.name, c.address, p.name, p.phone, s.name , t.task_notes, t.status_task from task t LEFT JOIN data_user uc ON t.user_id_create_task = uc.user_id LEFT JOIN data_user ud ON t.user_id_delegation_task= ud.user_id LEFT JOIN company c ON t.company_id_destination_task=c.company_id LEFT JOIN pic p ON t.pic_id_destination_task=p.pic_id LEFT JOIN sales s ON t.sales_id_destination_task=s.sales_id where t.task_id='%s'", taskCode)
	f.Println(getUserTask_sql)

	rows, err := conn.Query(getUserTask_sql)
	if err != nil {
		f.Println("Error reading records: ", err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var user_created_task string
		var assign_to string
		var deadline string
		var company string
		var company_id_pms string
		var company_alamat string
		var pic string
		var pic_phone string
		var sales string
		var task_notes string
		var status_task string
		var task_id string
		//var id int
		err := rows.Scan(&task_id, &user_created_task, &assign_to, &deadline, &company_id_pms, &company, &company_alamat, &pic, &pic_phone, &sales, &task_notes, &status_task)
		if err != nil {
			f.Println("Error reading rows: " + err.Error())
		}
		f.Println(task_id + "-" + user_created_task)
		tempDataTask.UserCreatedTask = user_created_task
		tempDataTask.AssignTask = assign_to
		tempDataTask.Deadline = deadline
		tempDataTask.Company = company_id_pms + "-" + company + "" + company_alamat
		tempDataTask.PIC = pic + "-" + pic_phone
		tempDataTask.Sales = sales
		tempDataTask.TaskNotes = task_notes
		tempDataTask.StatusTask = status_task
		tempDataTask.CodeTask = task_id

	}

	return tempDataTask

}

func DoneTask(codeTask string) (stsatus bool) {
	var status bool
	var err error

	ConnString := f.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;sslmode=disable", Server, User, Password, Port, Db)

	conn, err := sql.Open("sqlserver", ConnString)
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}
	f.Printf("Connected!\n")
	defer conn.Close()

	update_sql := f.Sprintf("UPDATE task set status_task='%s' where task_id='%s'", "done", codeTask)

	rows, err := conn.Query(update_sql)

	if err != nil {
		f.Println("Error occured while inserting a record", err.Error())
		return false
	}

	status = true
	defer rows.Close()
	return status
}
