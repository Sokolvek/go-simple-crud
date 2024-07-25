package storage

type Player struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Health int    `json:"hp"`
}
