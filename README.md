# Go_http_web_server2
A project demonstrating CRUD operations in Restful Api


## Sample code:
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func home(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Welcome to my website boys!")
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("api/vi", home)
	fmt.Println("listening at port 5002")
	
	log.Fatal(http.ListenAndServe(":5002", router))
}



### installing the Msql locally on your system (if you are using mac)

## Setting Up a MySQL Database Server

You can download it here:

» macOS: https://dev.mysql.com/downloads/file/?id=499568 

» Windows: https://dev.mysql.com/downloads/file/?id=499590


After you’ve downloaded the MySQL Community Server installer, run it and follow the various installation steps. Toward the end of the installation process, you’re asked to provide a password for the root account (the user account that has all the privileges in all the MySQL databases). Be sure to provide a secure password for the root account.
In the following sections, I explain how to:

» Interface with the MySQL server
» Create a database and a table
» Create a user account and grant it permission to access the database and table



## Interfacing with the MySQL server
There are a couple of ways to interface with the MySQL server. You can use the command-line utility mysql, or you can use MySQL Workbench (see Figure 17-1), which is a graphical user interface (GUI) application.
The MySQL installer for Windows automatically installs both the mysql utility and the MySQL Workbench app. 

Mac users need to manually download MySQL Work- bench from https://dev.mysql.com/downloads/file/?id=498743.


If you’re not a fan of typing commands in Terminal or Command Prompt, MySQL Workbench makes interfacing with MySQL less intimidating. But for this project , I show you how to interface with MySQL using the mysql utility.



In Terminal/Command Prompt, type the following command:

 $ mysql -u root -p
Enter password: <password>


Enter the password for the root account.

If you get an error that says the mysql utility is not found, the path to the utility isn’t set correctly, but you can fix that problem. Here’s how:
» If you’re on a Mac, type the following command:

 export PATH=$PATH:/usr/local/mysql/bin

 ###hurray!

Welcome to the MySQL monitor.  Commands end with ; or \g.
Your MySQL connection id is 9
Server version: 8.1.0 MySQL Community Server - GPL

Copyright (c) 2000, 2023, Oracle and/or its affiliates.

Oracle is a registered trademark of Oracle Corporation and/or its
affiliates. Other names may be trademarks of their respective
owners.

Type 'help;' or '\h' for help. Type '\c' to clear the current input statement.

## Creating a database and table

mysql> CREATE DATABASE CoursesDB;
Query OK, 1 row affected (0.03 sec)

mysql>  USE CoursesDB;
Database changed
mysql> CREATE TABLE Course (ID varchar(6) NOT NULL PRIMARY KEY, Details VARCHAR(100));
Query OK, 0 rows affected (0.02 sec)

mysql> INSERT INTO Course (ID, Details) VALUES ("IOT210","Applied Go Programming");
Query OK, 1 row affected (0.01 sec)

mysql> select * from course;
+--------+------------------------+
| ID     | Details                |
+--------+------------------------+
| IOT210 | Applied Go Programming |
+--------+------------------------+
1 row in set (0.00 sec)


## Creating a new account and granting permission

mysql> CREATE USER 'gouser'@'localhost' IDENTIFIED BY 'password';
Query OK, 0 rows affected (0.02 sec)

mysql> GRANT ALL ON *.* TO 'gouser'@'localhost';
Query OK, 0 rows affected (0.01 sec)

mysql> 

### Connecting to the MySQL Database in Go

With the MySQL Server configured with the database and user account, you can finally focus on getting your Go program to talk to the database server.
Go provides a SQL database application programming interface (API) in its stan- dard library database/sql package. However, the specific driver for the database server must be installed separately. This implementation allows developers to use a uniform API while at the same time being able to use different database servers.

To work with the MySQL server in your Go application, you can install the mysql driver by using this command in Terminal/Command Prompt:

$ go get "github.com/go-sql-driver/mysql"

To work with the standard SQL API, you just need to import the database/sql package, as well as the package for the driver of the database server that you’re using. For example, in this chapter, I’m using the MySQL server, so the import looks like this:

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

 ## Now write the code to connect to the MYsql server using Golang, 


 package main
import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
)

func main() {
// Use mysql as driverName and a valid DSN
 db, err := sql.Open("mysql",
        "gouser:password@tcp(127.0.0.1:3306)/CoursesDB")
    // handle error
    if err != nil {
        panic(err.Error())
    } else {
        fmt.Println("Database object created")
    }
    // defer the close till after the main function has
    // finished executing
    defer db.Close()

## Retrieving a record

Let’s now try to retrieve the record that is already in the Course table from the CoursesDB database. The statements 

in bold do the following:
» Create a type called Course with two fields: ID and Details.

» Define a function named Records to fetch records from the database and
then print them out.  

You use the Query() function to perform a query on the table. You usually use the Query() function to execute queries that return rows. This function returns a Rows struct containing the result of the query, as well as an error (if there is one). The cursor starts before the first row of the result set. You have to use the Next() method to move from one row to another. For each row, you use the Scan() method from the Rows struct to read the records from the table and map it to the fields in the Course struct that you’ve defined.






the full code goes below



package main

import (
    "database/sql"
"fmt"
    _ "github.com/go-sql-driver/mysql"
)

// map this type to the record in the table
type Course struct {
	ID string
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
	} }
// Now lets insert new records into the database
func InsertNewrecoreds(db *sql.DB,ID string,Details string){

	// use parametertised sql statment

	results,err := db.Exec(
		"INSERT INTO Course VALUES (?, ?)", ID, Details)
		if err !=nil{
        panic(err.Error())

		}else
		{
       if count , err :=results.RowsAffected(); err==nil{

	   fmt.Println(count,"rows(s) affected for this transation")


		}
	}
	}



	func main() {
		// Use mysql as driverName and a valid DSN
		db, err := sql.Open("mysql",
			"gouser:password@tcp(127.0.0.1:3306)/CoursesDB")
		// handle error
		if err != nil {
			panic(err.Error())
		} else {
			fmt.Println("Database object created")
			InsertNewrecoreds(db,"ph1012","php programming")
			InsertNewrecoreds(db,"jv104","java programming")
			InsertNewrecoreds(db,"CH103","C# programming")
			GetRecords(db)
	}
		// defer the close till after the main function has
		// finished executing
		defer db.Close()
	}



## The output is shown below

go run main.go
Database object created
1 rows(s) affected for this transation
1 rows(s) affected for this transation
1 rows(s) affected for this transation
Ad1011 - Adroid programming
CH103 - C# programming
IOS101 - IOS programming
IOT210 - Applied Go Programming
jv104 - java programming
ph1012 - php programming
py1003 - python programming
RU1011 - Rust programming

## congratulations! 
If you see the output above ,You’re now able to access the MySQL server. If you don’t see this output, double-check that the database username and password are correct and that you have a database named CoursesDB in your local MySQL server.


## Adding a record
When you’re able to retrieve rows from your table, you’re ready to write the code to insert new records into the table. The following function, InsertRecord(), does just that:

func InsertNewrecoreds(db *sql.DB,ID string,Details string){

	// use parametertised sql statment

	results,err := db.Exec(
		"INSERT INTO Course VALUES (?, ?)", ID, Details)
		if err !=nil{
        panic(err.Error())

		}else
		{
       if count , err :=results.RowsAffected(); err==nil{

	   fmt.Println(count,"rows(s) affected for this transation")


		}
	}
	}

Modifying a record
The next step would be to modify an existing record in the table. To do that, define the EditRecord() function as follows:


## Removing records


result, err := db.Exec(
		"SELECT * FROM Course WHERE ID=?", ID)

	if err != nil {

		fmt.Println(err.Error())
	

	
	} else {
		
			"DELETE FROM Course WHERE ID=?", ID)
	}
	
	if err != nil {
		panic(err.Error())

	} else {
		if count, err := resultjoint.RowsAffected(); err == nil {
			fmt.Println(count, "rows affected")

		}
	}

};

to check if records exist in the database , use the code below:

func Checkifexits(db *sql.DB, ID string) {

	result, err := db.Query(
		"SELECT count(*) FROM Course WHERE ID=ID")

	if err != nil {

		fmt.Println(err.Error())
	} else {
		
		for results.Next() {
			// map this type to the record in the table
			
			var course Course
			err = results.Scan(&course.ID,
				&course.Details)
				
				if &course.ID==nill{
					fmt.Println("This %ID: &ID does not exist")
				}else
			{
				Removerecords(db, ID)
			}


            The full application main.go codes are writen below


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




		
	
	

		