package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"rest-api-mysql/database"
	"rest-api-mysql/entity"

	"github.com/gorilla/mux"
)

func enableCors(w *http.ResponseWriter, r *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func GetArtists(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	enableCors(&w, r) // Allow CORS Policy
	var artists []entity.Artist
	result, err := database.Db.Query("SELECT artistId, artistName, albumName, imageURL, releaseDate, price, sampleURL from artists")
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	for result.Next() {
		var artist entity.Artist
		err := result.Scan(&artist.ArtistId, &artist.ArtistName, &artist.AlbumName, &artist.ImageURL, &artist.ReleaseDate, &artist.Price, &artist.SampleURL)
		if err != nil {
			panic(err.Error())
		}
		artists = append(artists, artist)
	}
	json.NewEncoder(w).Encode(artists)
}

func DeleteArtist(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	enableCors(&w, r) // Allow CORS Policy
	params := mux.Vars(r)
	stmt, err := database.Db.Prepare("DELETE FROM artists WHERE artistId = ?")
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec(params["id"])
	if err != nil {
		panic(err.Error())
	}
	json.NewEncoder(w).Encode("OK")
}

func GetArtist(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	result, err := database.Db.Query("SELECT artistId, artistName, albumName, imageURL, releaseDate, price, sampleURL FROM artists WHERE artistId = ?", params["id"])
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	var artist entity.Artist
	for result.Next() {
		err := result.Scan(&artist.ArtistId, &artist.ArtistName, &artist.AlbumName, &artist.ImageURL, &artist.ReleaseDate, &artist.Price, &artist.SampleURL)
		if err != nil {
			panic(err.Error())
		}
	}
	json.NewEncoder(w).Encode(artist)
}

func CreateArtist(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	enableCors(&w, r) // Allow CORS Policy
	stmt, err := database.Db.Prepare("INSERT INTO artists(artistName, albumName, imageURL, releaseDate, price, sampleURL) VALUES(?,?,?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	artistName := keyVal["artistName"]
	albumName := keyVal["albumName"]
	imageURL := keyVal["imageURL"]
	releaseDate := keyVal["releaseDate"]
	price := keyVal["price"]
	sampleURL := keyVal["sampleURL"]

	_, err = stmt.Exec(artistName, albumName, imageURL, releaseDate, price, sampleURL)
	if err != nil {
		panic(err.Error())
	}
	// fmt.Fprintf(w, "New artist was inserted into database")

	json.NewEncoder(w).Encode("OK")
}

func UpdateArtist(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	stmt, err := database.Db.Prepare("UPDATE artists SET artistName = ?, albumName = ?, imageURL = ?, releaseDate = ?, price = ?, sampleURL = ? WHERE artistId = ?")
	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	artistName := keyVal["artistName"]
	albumName := keyVal["albumName"]
	imageURL := keyVal["imageURL"]
	releaseDate := keyVal["releaseDate"]
	price := keyVal["price"]
	sampleURL := keyVal["sampleURL"]
	_, err = stmt.Exec(artistName, albumName, imageURL, releaseDate, price, sampleURL, params["id"])
	if err != nil {
		panic(err.Error())
	}
	json.NewEncoder(w).Encode("OK")
}
