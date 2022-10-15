package models

var StorageProviders = map[string]Storage{}
var SelectedProvider Storage

type Storage interface {
	GetAll() Hosts
	Set(Host)
	Add(host Host)
	SetLastSeen()
	Initialize(map[string]interface{}) interface{}
}
