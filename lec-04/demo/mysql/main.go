package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/tfs")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("CREATE TABLE tfs_users(id int NOT NULL AUTO_INCREMENT PRIMARY KEY, name varchar(100));")
	if err != nil {
		fmt.Println("Cannot create table: ", err)
		return
	}
	_, err = db.Exec(`INSERT INTO tfs_users (name) VALUES ("Truong");`)
	if err != nil {
		fmt.Println("Cannot insert into table: ", err)
		return
	}

	rows, err := db.Query("SELECT * FROM tfs_users LIMIT 1")
	if err != nil {
		fmt.Println("Cannot get from users")
		return
	}
	defer rows.Close()
	for ok := rows.Next(); ok; {
		rows.Scan()
	}
}
