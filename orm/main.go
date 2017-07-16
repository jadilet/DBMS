package main

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

// New model
type New struct {
	ID           uint `sql:"AUTO_INCREMENT" gorm:"primary_key"`
	ArticleTitle string
	ArticleURL   string
}

// TableName returns tablename
func (u *New) TableName() string {
	return "news"
}

// BeforeSave trigger
func (u *New) BeforeSave() (err error) {
	fmt.Println("Trigger before save")
	return
}

// FindByID select
func FindByID(id uint) New {
	new := New{}
	err := db.Find(&new, id).Error

	if err == gorm.ErrRecordNotFound {
		fmt.Println("Record Not found", id)
	} else {
		PanicOnError(err)
	}

	return new
}

func main() {
	var err error

	db, err = gorm.Open("mysql", "root:2468951@tcp(localhost:3306)/AllInOne_development?charset=utf8")
	PanicOnError(err)
	defer db.Close()

	// непосредственно подключаемся к базе
	db.DB()
	db.DB().Ping()

	//выбираем одиночную запись
	start := time.Now()
	new := FindByID(23)
	fmt.Println(new)
	elapsed := time.Since(start)
	fmt.Println("query time ", elapsed)

	// select all users
	start = time.Now()
	fmt.Println("Select all started")
	news := []New{}
	db.Find(&news)

	for _, new := range news {
		fmt.Println("ID", new.ID)
		fmt.Println(new)
	}

	fmt.Println("query select all done in", time.Since(start))

	// created new Record
	/*
		new = New{
			ArticleTitle: "Google",
			ArticleURL:   "https://google.com",
		}

		db.Create(&new)
		fmt.Println(new.ID)
		fmt.Println(new.ArticleTitle)
		fmt.Println(new.ArticleURL)
	*/
	// update Record
	// userNew.Telephone = "0555571571"
	// db.Save(userNew)
}

// PanicOnError catches errors
func PanicOnError(err error) {
	if err != nil {
		panic(err)
	}
}
