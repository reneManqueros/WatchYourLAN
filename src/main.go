package main

import (
	"time"
	"watchyourlan/models"
)

var AllHosts models.Hosts

func scanAndCompare() {
	var foundHosts models.Hosts
	var dbHosts models.Hosts
	for {
		foundHosts = arp_scan() // Scan interfaces
		dbHosts = models.SelectedProvider.GetAll()
		models.SelectedProvider.SetLastSeen()
		hosts_compare(foundHosts, dbHosts) // Compare hosts online and in DB

		// and add them to DB
		AllHosts = models.SelectedProvider.GetAll()
		time.Sleep(time.Duration(models.AppConfig.Timeout) * time.Second) // Timeout
	}
}

func main() {

	AllHosts = models.Hosts{}
	models.AppConfig = models.Conf{}
	models.AppConfig.Get()

	models.StorageProviders["sqlite"] = models.SQLiteProvider{}
	models.StorageProviders["mongodb"] = models.MongoDBProvider{}
	models.SelectedProvider = models.StorageProviders[models.AppConfig.DbProvider].Initialize(map[string]interface{}{
		"dbPath":        models.AppConfig.DbPath,
		"connectionURI": models.AppConfig.MongoDBConnectionURI,
		"database":      models.AppConfig.MongoDBDatabase,
		"collection":    models.AppConfig.MongoDBCollection,
	}).(models.Storage)

	go scanAndCompare()
	webgui() // Start web GUI
}
