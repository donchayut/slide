package main

// http://go-database-sql.org/
// https://rfam.readthedocs.io/en/latest/database.html

// 2019/07/16 11:44:29 CL00107 Created in error. ruthe
// 2019/07/16 11:44:29 CL00109 Only one family in clan after QC for Rfam 12.0 swb

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "rfamro@tcp(mysql-rfam-public.ebi.ac.uk:4497)/Rfam")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var (
		acc     string
		comment string
		user    string
	)
	rows, err := db.Query("select clan_acc, comment, user from dead_clan")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&acc, &comment, &user)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(acc, comment, user)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}
