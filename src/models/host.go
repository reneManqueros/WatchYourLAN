package models

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

func (h Host) Update() {
	SelectedProvider.Set(h)
}

func (h Host) Add() {
	SelectedProvider.Add(h)
}

func (h Hosts) SetLastSeen() {
	SelectedProvider.SetLastSeen()
}
