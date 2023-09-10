package main

import (
	"database/sql"
	"fmt"
	

	_ "github.com/go-sql-driver/mysql"
)

// map this type to the record in the table
type Course struct {
	ID      string
	Details string
}

func GetRecords(db *sql.DB) {
	results, err := db.Query("Select * FROM Course")

	if err != nil {
		panic(err.Error())
	}
	for results.Next() {
		// map this type to the record in the table
		var course Course
		err = results.Scan(&course.ID,
			&course.Details)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(course.ID,
			"-", course.Details)
	}
}

func EditRecord(db *sql.DB, ID string, Details string) {
	result, err := db.Exec(
		"UPDATE Course SET Details=? WHERE ID=?",
		Details, ID)
	if err != nil {
		fmt.Println(err.Error())

	} else {

		if count, err := result.RowsAffected(); err == nil {

			fmt.Println(count, "rows affected in this transation")
		}
	}
}

// Now lets insert new records into the database
func InsertNewrecoreds(db *sql.DB, ID string, Details string) {

	// use parametertised sql statment

	results, err := db.Exec(
		"INSERT INTO Course VALUES (?, ?)", ID, Details)
	if err != nil {
		panic(err.Error())

	} else {
		if count, err := results.RowsAffected(); err == nil {

			fmt.Println(count, "rows(s) affected for this transation")

		}
	}
}

//check if recordsexits before carying out removal operations


// Now lets try to remove records from the databse

func Removerecords(db *sql.DB, ID string) {

	//result, err := db.Query("select count(*)FROM Course WHERE ID=?,ID")
	//if err != nil {
		//log.Fatal(err)
	//}

	result, err := db.Exec(
        "DELETE FROM Course WHERE ID=?", ID)
    if err != nil {
        fmt.Println(err.Error())
    } else {
        if count, err := result.RowsAffected();
            err == nil {
            fmt.Println(count, "row(s) affected")
        }
} }



func main() {
	// Use mysql as driverName and a valid DSN
	db, err := sql.Open("mysql",
		"gouser:password@tcp(127.0.0.1:3306)/CoursesDB")
	// handle error
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Database object created")
		//InsertNewrecoreds(db,"ph1012","php programming")
		//EditRecord(db,"ph1012","laravel programming")
		Removerecords(db, "IOS101")
			

		}

		GetRecords(db)
	
	// defer the close till after the main function has
	// finished executing
	//defer db.Close()
	defer db.Close()
}



