package main

import (
	"database/sql"
	"log"
	//do a go get github.com/go-sql-driver/mysql before running this
	_ "github.com/go-sql-driver/mysql"
)

var (
	//this is temp. Vars should be passed on from the API module

	//DB schema:
	// +-------+-------------+------+-----+---------+----------------+
	// | Field | Type        | Null | Key | Default | Extra          |
	// +-------+-------------+------+-----+---------+----------------+
	// | id    | int(11)     | NO   | PRI | NULL    | auto_increment |
	// | name  | varchar(45) | NO   |     | NULL    |                |
	// | age   | int(11)     | NO   |     | NULL    |                |
	// | email | varchar(45) | YES  |     | NULL    |                |
	// +-------+-------------+------+-----+---------+----------------+

	id    int
	name  string
	age   int
	email string
	//notice the lower case
)

func main() {
	db, err := sql.Open("mysql",
		"root:password@tcp(127.0.0.1:3306)/db4")
	if err != nil {
		log.Fatal(err)
	}

	//TODO: modularize this code into funcs once its working
	//building and sending query
	id = 1
	rows, err := db.Query("select * from table1 where id = ?", id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(id, name)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	//close the db connection

	defer db.Close()

}
