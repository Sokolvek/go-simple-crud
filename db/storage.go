package storage

import (
	"encoding/json"
	"fmt"
	"os"
)

var id int
var Db DB

type DB struct {
	Players []Player
}

func (d *DB) GetPlayers() (string, error) {
	jsonData, _ := json.Marshal(d.Players)

	return string(jsonData), nil
}

func (d *DB) AddPlayer(plr Player) {

	player := Player{id, plr.Name, plr.Health}
	Db.Players = append(Db.Players, player)
	id++
	fmt.Println(Db)
}

func (d *DB) GeyPlayerById(id int) (Player, error) {
	plr := &Player{}

	for _, v := range Db.Players {
		if v.Id == id {
			return v, nil
		}
	}

	return *plr, nil
}

func (d *DB) RemovePlayer(id int) (bool, error) {
	plr, err := d.GeyPlayerById(id)
	if err != nil {
		return false, err
	}

	res := []Player{}

	for _, v := range Db.Players {
		if v.Id == plr.Id {
			continue
		}

		res = append(res, v)
	}

	Db.Players = res
	return true, nil
}

func (d *DB) InitDb() {
	_, err := os.Stat("db.json")

	if err != nil {
		createFile()
	}
}

func (d *DB) UpdateDB() {
	file, _ := os.Create("db.json")
	defer file.Close()

	dbValue, _ := json.Marshal(Db.Players)

	file.WriteString(string(dbValue))
	fmt.Println("updated", string(dbValue))
}

func createFile() {
	os.Create("db.json")
}
