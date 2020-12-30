package main

import (
	"database/sql"
	"strconv"

	"github.com/gin-gonic/gin"

	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type student struct {
	Id      int    `json:id`
	Name    string `json:name`
	Address string `json:Address`
	Branch  string `json:Branch`
	Age     int    `json:Age`
}

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "root"
	dbName := "student_db"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}
func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")
	//new template engine

	router.GET("/", func(ctx *gin.Context) {
		//render only file, must full name with extension
		db := dbConn()
		sd, err := db.Query("SELECT * FROM student ORDER BY id DESC")
		if err != nil {
			panic(err.Error())
		}
		emp := student{}
		res := []student{}
		for sd.Next() {
			var id, Age int
			var name, Address, Branch string
			err = sd.Scan(&id, &name, &Address, &Branch, &Age)
			if err != nil {
				panic(err.Error())
			}
			emp.Id = id
			emp.Name = name
			emp.Address = Address
			emp.Branch = Branch
			emp.Age = Age
			res = append(res, emp)
		}
		//var a = "hello words"
		ctx.HTML(http.StatusOK, "index.html", gin.H{"title": "index file title!!", "a": res})
	})

	router.GET("/about", func(ctx *gin.Context) {
		//render only file, must full name with extension
		ctx.HTML(http.StatusOK, "about.html", gin.H{"title": "index file title!!"})
	})

	router.GET("/contact", func(ctx *gin.Context) {
		//render only file, must full name with extension
		ctx.HTML(http.StatusOK, "contact.html", gin.H{"title": "index file title!!"})
	})

	router.GET("/addnew", func(ctx *gin.Context) {
		//render only file, must full name with extension
		ctx.HTML(http.StatusOK, "addnew.html", gin.H{"title": "index file title!!"})
	})

	router.GET("/submit", func(ctx *gin.Context) {
		//render only file, must full name with extension
		var name, Address, Branch string
		var Age int

		name = ctx.Request.FormValue("name")
		Address = ctx.Request.FormValue("Address")
		Branch = ctx.Request.FormValue("Branch")

		sal := ctx.Request.FormValue("Age")
		Age, _ = strconv.Atoi(sal)
		db := dbConn()
		insForm, err := db.Prepare("INSERT INTO student(name, Address, Branch, Age) VALUES(?,?,?,?)")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(name, Address, Branch, Age)

		sd, err := db.Query("SELECT * FROM student ORDER BY id DESC")
		if err != nil {
			panic(err.Error())
		}
		emp := student{}
		res := []student{}
		for sd.Next() {
			var id, Age int
			var name, Address, Branch string
			err = sd.Scan(&id, &name, &Address, &Branch, &Age)
			if err != nil {
				panic(err.Error())
			}
			emp.Id = id
			emp.Name = name
			emp.Address = Address
			emp.Branch = Branch
			emp.Age = Age
			res = append(res, emp)
		}

		ctx.HTML(http.StatusOK, "index.html", gin.H{"title": "index file title!!", "a": res})
	})

	router.Run(":9090")
}
