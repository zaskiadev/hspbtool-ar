package controllers

import (
	"database/sql"
	f "fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/julienschmidt/httprouter"
	"hpbtool-ar/db"
	"hpbtool-ar/models"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
)

var (
	Server   = "10.201.48.11"
	Port     = 5000
	User     = "edpbintaro"
	Password = "sqledpbintaro123"
	Db       = "hspb_tool_ar"
)

var UserCurrentLogin string = ""

type WebControllers struct{}

func (controller *WebControllers) Login(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	//db, err := gorm.Open(sqlite.Open("databasetask.db"), &gorm.Config{})
	//if err != nil {
	//	panic(err.Error())
	//}
	if r.Method == "POST" {
		var userName = r.FormValue("userName")
		var Password = r.FormValue("password")

		//var data = models.UserTasks{}
		//	var checkLogin = db.First(&data, "user_name=? and password=?", userName, Password)
		var count, _ = db.Login(userName, Password)
		if count > 0 {
			http.Redirect(w, r, "/home", http.StatusFound)
		}
	} else {
		files := "./views/login_user.html"

		htmlTemplate, err := template.ParseFiles(files)

		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		err = htmlTemplate.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func (controller *WebControllers) Home(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	/*db, err := gorm.Open(sqlite.Open("databasetask.db"), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}*/

	files := []string{
		"./views/base.html",
		"./views/home.html",
	}

	htmlTemplate, err := template.ParseFiles(files...)

	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	err = htmlTemplate.ExecuteTemplate(w, "base", "")
	if err != nil {
		println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func (controller *WebControllers) AddTask(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	/*	db, err := gorm.Open(sqlite.Open("databasetask.db"), &gorm.Config{})
		if err != nil {
			panic(err.Error())

		}

		if r.Method == "POST" {
			var selectedValue = r.FormValue("dataUserTaskTo")

			var data = models.UserTasks{}
			db.First(&data, "user_name=? ", strings.Replace(selectedValue, " ", "", -1))

			var getCodeName = data.CodeUserTask
			var codeTask = ""

			var dataTask = models.Task{}
			var dataGetCodeTask = db.Order("code_task desc").First(&dataTask)
			//println(dataGetCodeTask.Get("code_task"))
			if dataGetCodeTask == nil {
				codeTask = "TSK1"
			} else {

				var prefixOld = dataTask.CodeTask
				var getOnlyIntPrefixOld = strings.Replace(prefixOld, "TSK", "", -1)

				var getIncrement, _ = strconv.Atoi(getOnlyIntPrefixOld)
				println(getIncrement)
				getIncrement++
				println(getIncrement)
				codeTask = "TSK" + strconv.Itoa(getIncrement)
			}
			//db.Where("user_name = ?", selectedValue).Select("CodeUserTask").Find(&data)
			//var getSelectedUserTask = data.CodeUserTask
			//fmt.Println("selected user task : %s", getSelectedUserTask)
			println("code taxnya adalah : " + codeTask)
			task := models.Task{
				CodeTask:                codeTask,
				CodeUserCreateTask:      "0",
				CodeUserDestinationTask: getCodeName,
				Task:                    r.FormValue("taskDetail"),
				DateDeadLineTask:        r.FormValue("deadLineTask"),
				StatusTask:              "Create",
				TaskComment:             nil,
			}

			result := db.Create(&task)

			if result.Error != nil {
				log.Println(result.Error)
			}

			http.Redirect(w, r, "/home", http.StatusFound)
		} else {
	*/
	if r.Method == "POST" {
		r.FormValue("dataUserTaskTo")

	} else {
		files := []string{
			"./views/base.html",
			"./views/add_task.html",
		}

		htmlTemplate, err4 := template.ParseFiles(files...)

		if err4 != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			f.Println(err4.Error())
			return
		}

		var err error

		ConnString := f.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;sslmode=disable", Server, User, Password, Port, Db)

		conn, err := sql.Open("sqlserver", ConnString)
		if err != nil {
			log.Fatal("Open connection failed:", err.Error())
		}
		f.Printf("Connected!\n")
		defer conn.Close()

		var tempTask models.TempTask
		tempTask.TaskID = "tes"
		getTaskID_sql := f.Sprintf("SELECT TOP 1 * FROM task ORDER BY task_id DESC ")
		f.Println(getTaskID_sql)
		rows6, err6 := conn.Query(getTaskID_sql)
		defer rows6.Close()
		if err6 != nil {
			f.Println("Error reading records: ", err6.Error())

		} else {

			if rows6.Next() {
				var task_id string
				//var id int
				err6 := rows6.Scan(&task_id)
				if err6 != nil {
					f.Println("Error reading rows: " + err6.Error())
				}
				var prefixOld = task_id

				var getOnlyIntPrefixOld = strings.Replace(prefixOld, "TAR", "", -1)

				var getIncrement, _ = strconv.Atoi(getOnlyIntPrefixOld)
				println(getIncrement)
				getIncrement++
				println(getIncrement)
				tempTask.TaskID = "TAR" + strconv.Itoa(getIncrement)

			} else {
				tempTask.TaskID = "TAR1"
			}

		}
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
			err6 := rows.Scan(&user_id, &user_name)
			if err6 != nil {
				f.Println("Error reading rows: " + err6.Error())
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
		getPICTask_sql := f.Sprintf("select pic_id, name, guest_id_pms, phone, identification_id from pic")
		f.Println(getPICTask_sql)
		rows3, err3 := conn.Query(getPICTask_sql)
		if err3 != nil {
			f.Println("Error reading records: ", err3.Error())
		}
		defer rows3.Close()

		for rows3.Next() {
			var picId string
			var picName string
			var identificationId string
			var guestIDPMS string
			var phone string
			//var id int
			err3 := rows3.Scan(&picId, &picName, &guestIDPMS, &phone, &identificationId)
			if err3 != nil {
				f.Println("Error reading rows: " + err3.Error())
			}

			PicTask := models.TempPicTask{
				PicID:            picId,
				PicName:          picName,
				GuestIDPMS:       guestIDPMS,
				Phone:            phone,
				IdentificationID: identificationId,
			}
			tempDataPIC = append(tempDataPIC, PicTask)
		}

		var tempDataSales []models.TempSales
		getSalesTask_sql := f.Sprintf("select sales_id, name from sales")
		f.Println(getSalesTask_sql)
		rows5, err5 := conn.Query(getSalesTask_sql)
		if err5 != nil {
			f.Println("Error reading records: ", err5.Error())
		}
		defer rows5.Close()

		for rows5.Next() {
			var salesId string
			var salesName string

			//var id int
			err5 := rows5.Scan(&salesId, &salesName)
			if err5 != nil {
				f.Println("Error reading rows: " + err5.Error())
			}

			tempSales := models.TempSales{
				SalesID:   salesId,
				SalesName: salesName,
			}
			tempDataSales = append(tempDataSales, tempSales)
		}
		//var dataUser, dataCompany, dataPIC = db.GetDataTempAddTask()
		//var data []models.UserTasks
		//db.Find(&data)
		f.Println(tempTask.TaskID)
		datas := map[string]interface{}{
			"DataUser":    tempDataUserTask,
			"DataCompany": tempDataCompany,
			"DataPIC":     tempDataPIC,
			"DataSales":   tempDataSales,
			"DataTask":    tempTask,
		}
		err = htmlTemplate.ExecuteTemplate(w, "base", datas)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			f.Println(err.Error())
		}
	}
}

func (controller *WebControllers) AddCommentTask(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	/*	db, err := gorm.Open(sqlite.Open("databasetask.db"), &gorm.Config{})
		if err != nil {
			panic(err.Error())
		}

		files := []string{
			"./views/base.html",
			"./views/add_comment_task.html",
		}

		htmlTemplate, err := template.ParseFiles(files...)

		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		var DataTasks []models.Task
		db.Find(&DataTasks)
		datas := map[string]interface{}{
			"Tasks": DataTasks,
		}
		err = htmlTemplate.ExecuteTemplate(w, "base", datas)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	*/
}

func (controller *WebControllers) ShowAllCommentTask(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	/*	db, err := gorm.Open(sqlite.Open("databasetask.db"), &gorm.Config{})
		if err != nil {
			panic(err.Error())
		}

		files := []string{
			"./views/base.html",
			"./views/login_old.html",
			"./views/home.html",
			"./views/add_task.html",
			"./views/add_comment_task.html",
			"./views/show_all_comment_task ",
		}

		htmlTemplate, err := template.ParseFiles(files...)

		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		var ds []models.Task
		db.Find(&ds)
		datas := map[string]interface{}{
			"Tasks": ds,
		}

		//println(datas)
		err = htmlTemplate.ExecuteTemplate(w, "base", datas)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	*/
}

func (controller *WebControllers) EditTask(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	/*	db, err := gorm.Open(sqlite.Open("databasetask.db"), &gorm.Config{})
		if err != nil {
			panic(err.Error())
		}

		files := []string{
			"./views/base.html",
			"./views/edit_task.html",
		}

		htmlTemplate, err := template.ParseFiles(files...)

		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		/*var ds []models.Task
		db.Find(&ds)
		datas := map[string]interface{}{
			"Tasks": ds,
		}
		var task models.Task
		db.Where("code_task =?", params.ByName("codetask")).Find(&task)

		var user []models.UserTasks
		db.Find(&user)

		datas := map[string]interface{}{
			"Tasks": task,
			"User":  user,
		}

		//println(datas)
		fmt.Print(datas)
		err = htmlTemplate.ExecuteTemplate(w, "base", datas)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	*/
}

func (controller *WebControllers) UpdateTask(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	/*	db, err := gorm.Open(sqlite.Open("databasetask.db"), &gorm.Config{})
		if err != nil {
			panic(err.Error())
		}

		var selectedValue = r.FormValue("dataUserTaskTo")
		var data = models.UserTasks{}
		db.First(&data, "user_name=? ", strings.Replace(selectedValue, " ", "", -1))

		var getCodeName = data.CodeUserTask
		var taskCode = params.ByName("codetask")
		var task models.Task
		db.Where("code_task =?", taskCode).First(&task)

		task.Task = r.FormValue("taskDetail")
		task.DateDeadLineTask = r.FormValue("deadLineTask")
		task.CodeUserDestinationTask = getCodeName
		db.Save(&task)

		http.Redirect(w, r, "/home", http.StatusFound)
		/*var ds []models.Task
		db.Find(&ds)
		datas := map[string]interface{}{
			"Tasks": ds,
		}
	*/
}

func (controller *WebControllers) DoneTask(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	/*db, err := gorm.Open(sqlite.Open("databasetask.db"), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	var taskCode = params.ByName("codetask")
	var task models.Task

	println(taskCode)
	db.Where("code_task =?", taskCode).First(&task)

	task.StatusTask = "Done"

	db.Save(&task)

	http.Redirect(w, r, "/home", http.StatusFound)
	*/
}
