package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/vbua/streamer"
	"log"
)

var db *sql.DB

func connectToDb() {
	dbinfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		streamer.DbHost, streamer.DbPort, streamer.DbUser, streamer.DbPass, streamer.DbName)
	var err error
	db, err = sql.Open("postgres", dbinfo)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
}

func insertOrderToDb(uid, data string) {
	// инсертим или апдейтим если есть в базе
	stmt, err := db.Prepare(`INSERT INTO orders(uid, data) VALUES($1, $2) ON CONFLICT (uid) DO UPDATE SET data=$3`)
	if err != nil {
		log.Println("Couldn't prepare sql statement: ", err.Error())
		return
	}

	defer stmt.Close()

	_, err = stmt.Exec(uid, data, data)
	if err != nil {
		log.Println("Couldn't insert or update record: ", err.Error())
		return
	}
	log.Println("Saved order to database")
}

func getByUid(uid string) (string, error) {
	var data string
	err := db.QueryRow("SELECT data FROM orders WHERE uid=$1", uid).Scan(&data)
	if err != nil {
		return "", err
	}
	return data, nil
}

func fillUpCache() {
	rows, err := db.Query("SELECT * FROM orders")
	if err != nil {
		log.Println("Couldn't get all records from database: ", err.Error())
		return
	}
	for rows.Next() {
		var data string
		var uid string
		err = rows.Scan(&uid, &data)
		if err != nil {
			log.Println("Something went wrong when getting all records from database: ", err.Error())
			return
		}
		cache[uid] = data
	}
}
