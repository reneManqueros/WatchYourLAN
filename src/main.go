package main

import (
	"time"
	"watchyourlan/models"
)

var AllHosts models.Hosts

func scan_and_compare() {
	var foundHosts models.Hosts
	var dbHosts models.Hosts
	for { // Endless
		foundHosts = arp_scan()            // Scan interfaces
		dbHosts = models.HostsGetAll()     // Select everything from DB
		AllHosts.SetLastSeen()             // Mark hosts in DB as offline
		hosts_compare(foundHosts, dbHosts) // Compare hosts online and in DB
		// and add them to DB
		AllHosts = models.HostsGetAll()
		time.Sleep(time.Duration(models.AppConfig.Timeout) * time.Second) // Timeout
	}
}

func main() {
	AllHosts = models.Hosts{}
	models.AppConfig = models.Conf{}
	models.AppConfig.Get()

	AllHosts.CreateDBIfNew() // Check if DB exists. Create if not
	go scan_and_compare()
	webgui() // Start web GUI
}
