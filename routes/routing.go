package routing

import (
	"encoding/json"
	"fmt"
	"io"
	storage "learn/db"
	"net/http"
	"strconv"
	"strings"
)

var Mux = http.NewServeMux()
var db storage.DB = storage.Db

func InitializeRoutes() {
	// db = storage.Db
	Mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello bitch")
	})

	Mux.HandleFunc("/players", func(w http.ResponseWriter, r *http.Request) {
		players, _ := storage.Db.GetPlayers()

		w.Write([]byte(players))
	})

	Mux.HandleFunc("/add-player", func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Failed to read request body", http.StatusBadRequest)
			return
		}
		var player storage.Player
		defer r.Body.Close()
		json.Unmarshal(body, &player)
		db.AddPlayer(player)
		// fmt.Println(string(body))
		fmt.Println(storage.Db)
	})

	Mux.HandleFunc("/player/{id}", func(w http.ResponseWriter, r *http.Request) {
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

	})

	Mux.HandleFunc("/remove-player/{id}", func(w http.ResponseWriter, r *http.Request) {
		splitedUrl := strings.Split(r.URL.Path, "/")

		id, err := strconv.ParseInt(splitedUrl[len(splitedUrl)-1], 16, 32)
		if err != nil {
			return
		}

		state, _ := db.RemovePlayer(int(id))
		fmt.Println(db.Players)
		w.Write([]byte(strconv.FormatBool(state)))
	})

	Mux.HandleFunc("/upd", func(w http.ResponseWriter, r *http.Request) {
		db.UpdateDB()
	})
}
