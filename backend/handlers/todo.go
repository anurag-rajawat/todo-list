package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/anurag-rajawat/todo-list/backend/types"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"strconv"
)

var db *gorm.DB

func init() {
	log.SetFormatter(&log.TextFormatter{})
	loadEnv()
	connectToDB()
}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file, error:", err)
	}
}

func connectToDB() {
	dbUrl := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
	conn, err := gorm.Open("mysql", dbUrl)
	db = conn
	if err != nil {
		log.Fatal("unable to connect to database, error:", err)
	} else {
		log.Info("connected to the database")
	}
	db.AutoMigrate(&types.Todos{})
}

func CreateItem(w http.ResponseWriter, r *http.Request) {
	description := r.FormValue("description")
	log.Infof(`Adding new todo with "description": "%s"`, description)
	todo := &types.Todos{Description: description, Completed: false}
	db.Create(&todo)
	result := db.Last(&todo).Value
	setContentType(w)
	w.WriteHeader(http.StatusCreated)
	err := json.NewEncoder(w).Encode(result)
	if err != nil {
		log.Fatal(err)
	}
}

func UpdateItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	valid := getTodo(id)
	if !valid {
		setContentType(w)
		_, err := w.Write([]byte(`{"updated": false, "error": "Record Not Found"}`))
		w.WriteHeader(http.StatusNotFound)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		completed, _ := strconv.ParseBool(r.FormValue("completed"))
		log.Info("Updating todo")
		todo := &types.Todos{}
		db.First(&todo, id)
		todo.Completed = completed
		db.Save(&todo)
		setContentType(w)
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte(`{"updated": true}`))
		if err != nil {
			log.Fatal(err)
		}
	}
}

func DeleteItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	valid := getTodo(id)
	if !valid {
		setContentType(w)
		w.WriteHeader(http.StatusNotFound)
		_, err := w.Write([]byte(`{"deleted": false, "error": "Record Not Found"}`))
		if err != nil {
			log.Fatal(err)
		}
	} else {
		log.Info("Deleting todo")
		todo := &types.Todos{}
		db.First(&todo, id)
		db.Delete(&todo)
		setContentType(w)
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte(`{"deleted": true}`))
		if err != nil {
			log.Fatal(err)
		}
	}
}

func getTodo(Id int) bool {
	todo := &types.Todos{}
	result := db.First(&todo, Id)
	if result.Error != nil {
		log.Warn("Todo not found")
		return false
	}
	return true
}

func GetCompletedTodos(w http.ResponseWriter, r *http.Request) {
	log.Info("Get complete todos")
	completedTodos := getTodos(true)
	setContentType(w)
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(completedTodos)
	if err != nil {
		log.Fatal(err)
	}
}

func GetInCompletedTodos(w http.ResponseWriter, r *http.Request) {
	log.Info("Get incomplete todos")
	IncompleteTodos := getTodos(false)
	setContentType(w)
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(IncompleteTodos)
	if err != nil {
		log.Fatal(err)
	}
}

func getTodos(completed bool) interface{} {
	var todos []types.Todos
	TodoItems := db.Where("completed = ?", completed).Find(&todos).Value
	return TodoItems
}

func setContentType(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}
