package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
	c  redis.Conn
)

// New model
type New struct {
	ID           uint `sql:"AUTO_INCREMENT" gorm:"primary_key"`
	ArticleTitle string
	ArticleURL   string
}

// TableName return table names
func (u *New) TableName() string {
	return "news"
}

// BeforeSave trigger on news
func (u *New) BeforeSave() (err error) {
	fmt.Println("Trigger before save")
	return
}

// FindByID find records by id
func FindByID(id uint) New {
	new := New{}
	err := db.Find(&new, id).Error

	if err == gorm.ErrRecordNotFound {
		fmt.Println("Record not found with id ", id)
	} else {
		PanicOnError(err)
	}

	return new
}

// PanicOnError catches error
func PanicOnError(err error) {
	if err != nil {
		panic(err)
	}
}

func getCacheRecord(key string) string {
	// Получает запись
	item, err := redis.String(c.Do("GET", key))

	if err == redis.ErrNil {
		fmt.Println("Record not found in redis (return value is nil)")
		fmt.Println("Cache created!!")

		news := []New{}
		db.Find(&news)

		jsonNews, _ := json.Marshal(news)

		result, _ := redis.String(c.Do("SET", key, jsonNews))

		if result != "OK" {
			panic("Result not ok" + result)
		}

		return string(jsonNews)
	} else if err != nil {
		PanicOnError(err)
	}

	fmt.Println("News from cache")
	return item
}

func main() {
	var err error

	// Mysql connection
	db, err = gorm.Open("mysql", "root:2468951@tcp(localhost:3306)/AllInOne_development?charset=utf8")
	PanicOnError(err)
	defer db.Close()

	// Redis connection
	c, err = redis.Dial("tcp", ":6379")
	PanicOnError(err)

	// Connecting database
	db.DB()
	db.DB().Ping()

	http.HandleFunc("/news", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL.Path)
		defer r.Body.Close()

		switch r.Method {
		case http.MethodGet:
			start := time.Now()
			jsonNews := getCacheRecord("news")
			done := time.Since(start)

			fmt.Println("Query done in mysql ", done)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(jsonNews))
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	http.ListenAndServe(":5715", nil)
}
