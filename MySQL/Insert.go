//Inserting data into the database
package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func insert(db *sql.DB) {
	stmt, err := db.Prepare("INSERT INTO user(uusername, passwor) VALUES(?, ?)")
	defer stmt.Close()

	if err != nil {
		log.Println(err)
		return
	}

	stmt.Exec("guotie", "guotie")
	stmt.Exec("testuser", "123123")
}

func main() {
	db, err := sql.Open("mysql", "root:096407@/hello")
	if err != nil {
		log.Fatalf("Open database error: %s\n", err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Fatalf(err)
	}
	insert(db)

	rows, err := db.Query("select id, username form user where id = ?", 1)
	if err != nil {
		log.Println(err)
	}

	defer rows.Close()
	var id int
	var name string
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

}
