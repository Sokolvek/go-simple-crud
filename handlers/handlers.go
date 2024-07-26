package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	storage "learn/db"
	"learn/middleware"
	"net/http"
	"strconv"
	"strings"
)

var db storage.DB = storage.Db

func GetPlayers(w http.ResponseWriter, r *http.Request) {
	players, _ := storage.Db.GetPlayers()

	w.Write([]byte(players))
}

func AddPlayer(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}
	var player storage.Player
	defer r.Body.Close()
	json.Unmarshal(body, &player)
	db.AddPlayer(player)
	fmt.Println(storage.Db)
}

func AddPlayerHandler() http.Handler {
	return middleware.UpdateDB(http.HandlerFunc(AddPlayer))
}

func GeyPlayerById(w http.ResponseWriter, r *http.Request) {
	splitedUrl := strings.Split(r.URL.Path, "/")

	id, err := strconv.ParseInt(splitedUrl[len(splitedUrl)-1], 16, 32)
	if err != nil {
		return
	}

	plr, err := db.GeyPlayerById(int(id))
	if err != nil {
		return
	}

	jsonPlr, err := json.Marshal(plr)
	if err != nil {
		return
	}

	w.Write([]byte(jsonPlr))

}

func RemovePlayer(w http.ResponseWriter, r *http.Request) {
	splitedUrl := strings.Split(r.URL.Path, "/")

	id, err := strconv.ParseInt(splitedUrl[len(splitedUrl)-1], 16, 32)
	if err != nil {
		return
	}

	state, _ := db.RemovePlayer(int(id))
	fmt.Println(db.Players)
	w.Write([]byte(strconv.FormatBool(state)))
}

func RemovePlayerHandler() http.Handler {
	return middleware.UpdateDB(http.HandlerFunc(RemovePlayer))
}
