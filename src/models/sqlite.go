package models

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"watchyourlan/helpers"
)

type SQLiteProvider struct {
	DatabasePath string
}

func (p SQLiteProvider) GetAll() (dbHosts Hosts) {
	db, err := sql.Open("sqlite3", p.DatabasePath)
	if err != nil {
		log.Fatal("ERROR: open: ", err, p.DatabasePath)
	}

	defer db.Close()

	sqlStatement := `SELECT * FROM "now" ORDER BY DATE DESC`
	res, err := db.Query(sqlStatement)
	if err != nil {
		log.Fatal("ERROR: db_select: ", err)
	}

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

	return dbHosts
}

func (p SQLiteProvider) Set(h Host) {
	h.Name = helpers.Quote(h.Name)
	h.Hw = helpers.Quote(h.Hw)
	sqlStatement := `UPDATE "now" set 
		NAME = '%s', IP = '%s', MAC = '%s', HW = '%s', DATE = '%s', 
		KNOWN = '%d', NOW = '%d' 
		WHERE ID = '%d';`
	sqlStatement = fmt.Sprintf(sqlStatement, h.Name, h.Ip, h.Mac, h.Hw, h.Date, h.Known, h.Now, h.Id)
	p.execute(sqlStatement)
}

func (p SQLiteProvider) SetLastSeen() {
	sqlStatement := `UPDATE "now" set NOW = '0';`
	p.execute(sqlStatement)
}

func (p SQLiteProvider) Add(h Host) {
	h.Name = helpers.Quote(h.Name)
	h.Hw = helpers.Quote(h.Hw)
	sqlStatement := `INSERT INTO "now" (NAME, IP, MAC, HW, DATE, KNOWN, NOW) 
		VALUES ('%s','%s','%s','%s','%s','%d','%d');`
	sqlStatement = fmt.Sprintf(sqlStatement, h.Name, h.Ip, h.Mac, h.Hw, h.Date, h.Known, h.Now)
	p.execute(sqlStatement)
}

func (p SQLiteProvider) execute(sqlStatement string) {
	db, err := sql.Open("sqlite3", p.DatabasePath)
	if err != nil {
		log.Fatal("ERROR: db_exec: ", err)
	}
	defer func() {
		err := db.Close()
		if err != nil {
			log.Println("ERROR: sqlite::execute", err)
		}
	}()

	_, err = db.Exec(sqlStatement)
	if err != nil {
		log.Fatal("ERROR: db_exec: ", err)
	}
}

func (p SQLiteProvider) Initialize(connectionString map[string]interface{}) interface{} {
	if val, ok := connectionString["dbPath"]; ok {
		p.DatabasePath = val.(string)
		sqlStatement := `CREATE TABLE  IF NOT EXISTS "now" (
			"ID"	INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT UNIQUE,
			"NAME"	TEXT NOT NULL,
			"IP"	TEXT,
			"MAC"	TEXT,
			"HW"	TEXT,
			"DATE"	TEXT,
			"KNOWN"	INTEGER DEFAULT 0,
			"NOW"	INTEGER DEFAULT 0
		);`
		p.execute(sqlStatement)
	}
	return p
}
