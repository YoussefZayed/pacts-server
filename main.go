package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"pacts/server/db"

	"github.com/gorilla/mux"
)

func main() {
	defer db.DB.Close()

	r := mux.NewRouter()

	r.HandleFunc("/tiles", createTileHandler).Methods("POST")
	r.HandleFunc("/tiles/{id}", getTileByIDHandler).Methods("GET")
	r.HandleFunc("/tiles/coordinates", getTileByCoordinatesHandler).Methods("GET")
	r.HandleFunc("/tiles", getAllTilesHandler).Methods("GET")
	r.HandleFunc("/tiles", updateTileHandler).Methods("PUT")
	r.HandleFunc("/tiles/{id}", deleteTileHandler).Methods("DELETE")
	r.HandleFunc("/init-tiles", initTilesHandler).Methods("POST")

	fmt.Println("Server listening on :6080")
	http.ListenAndServe(":6080", r)
}

func createTileHandler(w http.ResponseWriter, r *http.Request) {
	var tile db.Tile
	err := json.NewDecoder(r.Body).Decode(&tile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdTile, err := db.CreateTile(&tile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdTile)
}

func getTileByIDHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid tile ID", http.StatusBadRequest)
		return
	}

	tile, err := db.GetTileByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tile)
}

func getTileByCoordinatesHandler(w http.ResponseWriter, r *http.Request) {
	x, errX := strconv.Atoi(r.URL.Query().Get("x"))
	y, errY := strconv.Atoi(r.URL.Query().Get("y"))
	if errX != nil || errY != nil {
		http.Error(w, "Invalid coordinates", http.StatusBadRequest)
		return
	}

	tile, err := db.GetTileByCoordinates(x, y)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tile)
}

func getAllTilesHandler(w http.ResponseWriter, r *http.Request) {
	tiles, err := db.GetAllTiles()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tiles)
}

func updateTileHandler(w http.ResponseWriter, r *http.Request) {
	var tile db.Tile
	err := json.NewDecoder(r.Body).Decode(&tile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = db.UpdateTile(&tile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func deleteTileHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid tile ID", http.StatusBadRequest)
		return
	}

	err = db.DeleteTile(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func initTilesHandler(w http.ResponseWriter, r *http.Request) {
	x, errX := strconv.Atoi(r.URL.Query().Get("maxX"))
	y, errY := strconv.Atoi(r.URL.Query().Get("maxY"))
	if errX != nil || errY != nil {
		http.Error(w, "Invalid max coordinates", http.StatusBadRequest)
		return
	}

	err := db.InitTiles(x, y)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
