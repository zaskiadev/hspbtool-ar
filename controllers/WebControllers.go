package controllers

import (
	"github.com/julienschmidt/httprouter"
	"hpbtool-ar/db"
	"html/template"
	"log"
	"net/http"
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
			files := []string{
				"./views/base.html",
				"./views/add_task.html",
			}

			htmlTemplate, err := template.ParseFiles(files...)

			if err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}

			var data []models.UserTasks
			db.Find(&data)
			datas := map[string]interface{}{
				"DataUser": data,
			}
			err = htmlTemplate.ExecuteTemplate(w, "base", datas)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}

		}
	*/
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
