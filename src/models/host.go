package models

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"watchyourlan/helpers"
)

type Host struct {
	Id    uint16
	Name  string
	Ip    string
	Mac   string
	Hw    string
	Date  string
	Known uint16
	Now   uint16
}

type Hosts []Host

func HostsGetAll() (dbHosts []Host) {
	db, _ := sql.Open("sqlite3", AppConfig.DbPath)
	defer db.Close()

	sqlStatement := `SELECT * FROM "now" ORDER BY DATE DESC`
	res, err := db.Query(sqlStatement)
	if err != nil {
		log.Fatal("ERROR: db_select: ", err)
	}

	dbHosts = Hosts{}
	for res.Next() {
		var oneHost Host
		err = res.Scan(&oneHost.Id, &oneHost.Name, &oneHost.Ip, &oneHost.Mac, &oneHost.Hw, &oneHost.Date, &oneHost.Known, &oneHost.Now)
		if err != nil {
			log.Fatal(err)
		}
		oneHost.Name = helpers.Unquote(oneHost.Name)
		oneHost.Hw = helpers.Unquote(oneHost.Hw)
		dbHosts = append(dbHosts, oneHost)
	}

	//fmt.Println("Select all:", dbHosts)
	return dbHosts

}

func (h Host) Update() {
	h.Name = helpers.Quote(h.Name)
	h.Hw = helpers.Quote(h.Hw)
	sqlStatement := `UPDATE "now" set 
		NAME = '%s', IP = '%s', MAC = '%s', HW = '%s', DATE = '%s', 
		KNOWN = '%d', NOW = '%d' 
		WHERE ID = '%d';`
	sqlStatement = fmt.Sprintf(sqlStatement, h.Name, h.Ip, h.Mac, h.Hw, h.Date, h.Known, h.Now, h.Id)
	db_exec(sqlStatement)
}

func (h Host) Add() {
	h.Name = helpers.Quote(h.Name)
	h.Hw = helpers.Quote(h.Hw)
	sqlStatement := `INSERT INTO "now" (NAME, IP, MAC, HW, DATE, KNOWN, NOW) 
		VALUES ('%s','%s','%s','%s','%s','%d','%d');`
	sqlStatement = fmt.Sprintf(sqlStatement, h.Name, h.Ip, h.Mac, h.Hw, h.Date, h.Known, h.Now)
	//fmt.Println("Insert statement:", sqlStatement)
	db_exec(sqlStatement)
}

func (h Hosts) SetLastSeen() {
	sqlStatement := `UPDATE "now" set NOW = '0';`
	db_exec(sqlStatement)
}

func (h Hosts) CreateDBIfNew() {
	log.Println("AppConfig.DbPath", AppConfig.DbPath)
	if _, err := os.Stat(AppConfig.DbPath); err == nil {
		log.Println("INFO: DB exists")
	} else {
		sqlStatement := `CREATE TABLE "now" (
			"ID"	INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT UNIQUE,
			"NAME"	TEXT NOT NULL,
			"IP"	TEXT,
			"MAC"	TEXT,
			"HW"	TEXT,
			"DATE"	TEXT,
			"KNOWN"	INTEGER DEFAULT 0,
			"NOW"	INTEGER DEFAULT 0
		);`
		log.Println("INFO: Table created!")
		db_exec(sqlStatement)
	}
}

func db_exec(sqlStatement string) {
	db, err := sql.Open("sqlite3", AppConfig.DbPath)
	if err != nil {
		log.Fatal("ERROR: db_exec: ", err)
	}
	defer db.Close()

	_, err = db.Exec(sqlStatement)
	if err != nil {
		log.Fatal("ERROR: db_exec: ", err)
	}
}
